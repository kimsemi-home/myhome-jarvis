package security

import (
	"io/fs"
	"path/filepath"
	"sort"
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
