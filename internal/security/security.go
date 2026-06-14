package security

import (
	"bytes"
	"errors"
	"io/fs"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Finding struct {
	Path    string `json:"path"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Report struct {
	Root     string    `json:"root"`
	OK       bool      `json:"ok"`
	Findings []Finding `json:"findings"`
}

type HistoryFinding struct {
	Commit  string `json:"commit,omitempty"`
	Path    string `json:"path"`
	Line    int    `json:"line,omitempty"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HistoryReport struct {
	Root     string           `json:"root"`
	OK       bool             `json:"ok"`
	Findings []HistoryFinding `json:"findings"`
}

type historyPattern struct {
	Code    string
	Pattern string
	Message string
}

type gitCommandError struct {
	err     error
	message string
}

func (err gitCommandError) Error() string {
	return err.message
}

func (err gitCommandError) Unwrap() error {
	return err.err
}

var secretHistoryPattern = regexp.MustCompile(`(?i)(BEGIN (RSA|OPENSSH|EC|DSA) PRIVATE KEY|Authorization:[[:space:]]*Bearer[[:space:]]+[A-Za-z0-9._~+/=-]{20,}|(api[_-]?key|secret|password|token)[[:space:]]*[:=][[:space:]]*[A-Za-z0-9._~+/=-]{20,})`)

func Check(root string) (Report, error) {
	report := Report{Root: root, OK: true}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}
		if entry.IsDir() {
			if shouldSkipDir(entry.Name()) {
				return filepath.SkipDir
			}
			if strings.EqualFold(entry.Name(), "secrets") {
				report.add(rel, "forbidden_secret_dir", "secrets directories must not be tracked")
				return filepath.SkipDir
			}
			return nil
		}
		checkPath(rel, &report)
		return nil
	})
	sort.Slice(report.Findings, func(i, j int) bool {
		if report.Findings[i].Path == report.Findings[j].Path {
			return report.Findings[i].Code < report.Findings[j].Code
		}
		return report.Findings[i].Path < report.Findings[j].Path
	})
	report.OK = len(report.Findings) == 0
	return report, err
}

func CheckHistory(root string) (HistoryReport, error) {
	report := HistoryReport{Root: ".", OK: true}
	if _, err := exec.LookPath("git"); err != nil {
		return report, err
	}
	commits, err := gitLines(root, "rev-list", "--all")
	if err != nil {
		return report, err
	}
	if len(commits) == 0 {
		return report, nil
	}
	if err := checkHistoryPaths(root, &report); err != nil {
		return report, err
	}
	if err := checkHistoryMetadata(root, &report); err != nil {
		return report, err
	}
	privateIdentityPattern := privateIdentityHistoryPattern()
	patterns := []historyPattern{
		{
			Code:    "history_private_identity",
			Pattern: privateIdentityPattern,
			Message: "git history must not contain private user, path, or old-repository identity markers",
		},
		{
			Code:    "history_secret_literal",
			Pattern: `(BEGIN (RSA|OPENSSH|EC|DSA) PRIVATE KEY|Authorization:[[:space:]]*Bearer[[:space:]]+[A-Za-z0-9._~+/=-]{20,}|(api[_-]?key|secret|password|token)[[:space:]]*[:=][[:space:]]*[A-Za-z0-9._~+/=-]{20,})`,
			Message: "git history must not contain secret-looking literal values",
		},
	}
	for _, pattern := range patterns {
		if err := checkHistoryContent(root, commits, pattern, &report); err != nil {
			return report, err
		}
	}
	sort.Slice(report.Findings, func(i, j int) bool {
		left := report.Findings[i]
		right := report.Findings[j]
		if left.Commit != right.Commit {
			return left.Commit < right.Commit
		}
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		if left.Line != right.Line {
			return left.Line < right.Line
		}
		return left.Code < right.Code
	})
	report.OK = len(report.Findings) == 0
	return report, nil
}

func shouldSkipDir(name string) bool {
	switch name {
	case ".git", "target", "build", "dist", "bin", ".dart_tool":
		return true
	default:
		return false
	}
}

func checkPath(rel string, report *Report) {
	base := strings.ToLower(filepath.Base(rel))
	ext := strings.ToLower(filepath.Ext(rel))

	if rel == ".env" || (strings.HasPrefix(rel, ".env.") && rel != ".env.example") {
		report.add(rel, "forbidden_env_file", "environment files must stay local")
	}
	if strings.HasPrefix(rel, "data/private/") || strings.HasPrefix(rel, "data/lake/") {
		return
	}

	switch ext {
	case ".py", ".pyi", ".ipynb":
		report.add(rel, "forbidden_python_file", "Python files are not allowed")
	case ".ts", ".tsx", ".js", ".jsx", ".mjs", ".cjs":
		report.add(rel, "forbidden_node_typescript_file", "Node.js and TypeScript files are not allowed")
	}

	switch base {
	case "pyproject.toml", "pytest.ini", "requirements.txt", "requirements-dev.txt", "poetry.lock", "uv.lock", "pipfile", "pipfile.lock", "setup.py", "tox.ini":
		report.add(rel, "forbidden_python_dependency", "Python dependency or tooling files are not allowed")
	case "package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock", "tsconfig.json":
		report.add(rel, "forbidden_node_dependency", "Node.js and TypeScript dependency files are not allowed")
	}

	for _, marker := range []string{"token", "secret", "credential", "cookie"} {
		if strings.Contains(base, marker) {
			report.add(rel, "forbidden_sensitive_path", "sensitive-looking files must stay under data/private")
			break
		}
	}
}

func (report *Report) add(path string, code string, message string) {
	report.Findings = append(report.Findings, Finding{Path: path, Code: code, Message: message})
}

func checkHistoryPaths(root string, report *HistoryReport) error {
	lines, err := gitLines(root, "log", "--all", "--name-only", "--pretty=format:__MHJ_COMMIT__%H")
	if err != nil {
		return err
	}
	commit := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "__MHJ_COMMIT__") {
			commit = strings.TrimPrefix(line, "__MHJ_COMMIT__")
			continue
		}
		rel := filepath.ToSlash(strings.TrimSpace(line))
		if rel == "" {
			continue
		}
		checkHistoryPath(commit, rel, report)
	}
	return nil
}

func checkHistoryPath(commit string, rel string, report *HistoryReport) {
	base := strings.ToLower(filepath.Base(rel))
	ext := strings.ToLower(filepath.Ext(rel))

	if rel == ".env" || (strings.HasPrefix(rel, ".env.") && rel != ".env.example") {
		report.addHistory(commit, rel, 0, "history_forbidden_env_file", "environment files must never appear in git history")
	}
	if (strings.HasPrefix(rel, "data/private/") || strings.HasPrefix(rel, "data/lake/")) && !isAllowedPrivatePlaceholder(rel) {
		report.addHistory(commit, rel, 0, "history_private_data_path", "private data and lake files must never appear in git history")
	}
	if hasPathSegment(rel, "secrets") {
		report.addHistory(commit, rel, 0, "history_forbidden_secret_dir", "secrets directories must never appear in git history")
	}

	switch ext {
	case ".py", ".pyi", ".ipynb":
		report.addHistory(commit, rel, 0, "history_forbidden_python_file", "Python files must never appear in git history")
	case ".ts", ".tsx", ".js", ".jsx", ".mjs", ".cjs":
		report.addHistory(commit, rel, 0, "history_forbidden_node_typescript_file", "Node.js and TypeScript files must never appear in git history")
	}

	switch base {
	case "pyproject.toml", "pytest.ini", "requirements.txt", "requirements-dev.txt", "poetry.lock", "uv.lock", "pipfile", "pipfile.lock", "setup.py", "tox.ini":
		report.addHistory(commit, rel, 0, "history_forbidden_python_dependency", "Python dependency or tooling files must never appear in git history")
	case "package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock", "tsconfig.json":
		report.addHistory(commit, rel, 0, "history_forbidden_node_dependency", "Node.js and TypeScript dependency files must never appear in git history")
	}

	for _, marker := range []string{"token", "secret", "credential", "cookie"} {
		if strings.Contains(base, marker) {
			report.addHistory(commit, rel, 0, "history_forbidden_sensitive_path", "sensitive-looking files must never appear in git history")
			break
		}
	}
}

func checkHistoryMetadata(root string, report *HistoryReport) error {
	lines, err := gitLines(root, "log", "--all", "--format=%H%x1f%an%x1f%ae%x1f%cn%x1f%ce%x1f%s")
	if err != nil {
		return err
	}
	privateIdentity := regexp.MustCompile(`(?i)` + privateIdentityHistoryPattern())
	for _, line := range lines {
		parts := strings.Split(line, "\x1f")
		if len(parts) < 6 {
			continue
		}
		commit := parts[0]
		for _, field := range parts[1:] {
			if privateIdentity.MatchString(field) {
				report.addHistory(commit, "(commit metadata)", 0, "history_private_identity_metadata", "git commit metadata must not contain private identity markers")
				break
			}
			if secretHistoryPattern.MatchString(field) {
				report.addHistory(commit, "(commit metadata)", 0, "history_secret_metadata", "git commit metadata must not contain secret-looking literal values")
				break
			}
		}
	}
	return nil
}

func privateIdentityHistoryPattern() string {
	oldOwner := strings.Join([]string{"kim", "jooyoon"}, "")
	oldTeam := strings.Join([]string{"kim-joo", "-yoon"}, "")
	localUser := strings.Join([]string{"al", "ice"}, "")
	return strings.Join([]string{
		oldOwner,
		oldTeam,
		"github\\.com/" + oldOwner,
		"/" + "Users",
		"(^|[^[:alnum:]_])" + localUser + "([^[:alnum:]_]|$)",
	}, "|")
}

func checkHistoryContent(root string, commits []string, pattern historyPattern, report *HistoryReport) error {
	const batchSize = 64
	for start := 0; start < len(commits); start += batchSize {
		end := start + batchSize
		if end > len(commits) {
			end = len(commits)
		}
		args := []string{"grep", "-n", "-I", "-E", "-e", pattern.Pattern}
		args = append(args, commits[start:end]...)
		args = append(args, "--")
		lines, err := gitLinesAllowNoMatches(root, args...)
		if err != nil {
			return err
		}
		for _, line := range lines {
			commit, path, lineNumber, ok := parseGitGrepLine(line)
			if !ok {
				continue
			}
			report.addHistory(commit, path, lineNumber, pattern.Code, pattern.Message)
		}
	}
	return nil
}

func parseGitGrepLine(line string) (string, string, int, bool) {
	if len(line) < 42 || line[40] != ':' {
		return "", "", 0, false
	}
	commit := line[:40]
	rest := line[41:]
	pathEnd := strings.IndexByte(rest, ':')
	if pathEnd < 0 {
		return "", "", 0, false
	}
	path := filepath.ToSlash(rest[:pathEnd])
	rest = rest[pathEnd+1:]
	lineEnd := strings.IndexByte(rest, ':')
	if lineEnd < 0 {
		return "", "", 0, false
	}
	lineNumber, err := strconv.Atoi(rest[:lineEnd])
	if err != nil {
		return "", "", 0, false
	}
	return commit, path, lineNumber, true
}

func gitLines(root string, args ...string) ([]string, error) {
	output, err := gitOutput(root, args...)
	if err != nil {
		return nil, err
	}
	return splitOutputLines(output), nil
}

func gitLinesAllowNoMatches(root string, args ...string) ([]string, error) {
	output, err := gitOutput(root, args...)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, err
	}
	return splitOutputLines(output), nil
}

func gitOutput(root string, args ...string) (string, error) {
	allArgs := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", allArgs...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		message := strings.TrimSpace(output.String())
		if message == "" {
			message = "git command failed"
		}
		return "", gitCommandError{err: err, message: message}
	}
	return output.String(), nil
}

func splitOutputLines(output string) []string {
	trimmed := strings.TrimRight(output, "\n")
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, "\n")
}

func hasPathSegment(rel string, segment string) bool {
	for _, part := range strings.Split(filepath.ToSlash(rel), "/") {
		if strings.EqualFold(part, segment) {
			return true
		}
	}
	return false
}

func isAllowedPrivatePlaceholder(rel string) bool {
	switch filepath.ToSlash(rel) {
	case "data/private/.keep", "data/private/.gitkeep", "data/lake/.keep", "data/lake/.gitkeep":
		return true
	default:
		return false
	}
}

func (report *HistoryReport) addHistory(commit string, path string, line int, code string, message string) {
	report.Findings = append(report.Findings, HistoryFinding{
		Commit:  commit,
		Path:    filepath.ToSlash(path),
		Line:    line,
		Code:    code,
		Message: message,
	})
}
