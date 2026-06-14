package repo

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Status struct {
	Branch              string   `json:"branch"`
	HeadSHA             string   `json:"head_sha"`
	ShortSHA            string   `json:"short_sha"`
	Upstream            string   `json:"upstream,omitempty"`
	Origin              string   `json:"origin,omitempty"`
	WorktreeClean       bool     `json:"worktree_clean"`
	TrackedChanges      []Change `json:"tracked_changes"`
	UntrackedFiles      []string `json:"untracked_files"`
	IgnoredPrivatePaths []string `json:"ignored_private_paths"`
	CheckedAt           string   `json:"checked_at"`
}

type Change struct {
	Code string `json:"code"`
	Path string `json:"path"`
}

func Inspect(root string) (Status, error) {
	if strings.TrimSpace(root) == "" {
		return Status{}, errors.New("repo root is required")
	}
	if _, err := git(root, "rev-parse", "--show-toplevel"); err != nil {
		return Status{}, fmt.Errorf("inspect git worktree: %w", err)
	}

	branch, _ := git(root, "rev-parse", "--abbrev-ref", "HEAD")
	head, err := git(root, "rev-parse", "HEAD")
	if err != nil {
		return Status{}, fmt.Errorf("inspect git head: %w", err)
	}
	upstream, _ := git(root, "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	origin, _ := git(root, "config", "--get", "remote.origin.url")
	porcelain, err := gitRaw(root, "status", "--porcelain=v1", "-z")
	if err != nil {
		return Status{}, fmt.Errorf("inspect git status: %w", err)
	}
	ignored, _ := git(root, "status", "--ignored", "--short", "--", "data/private", "data/lake")
	tracked, untracked := parsePorcelain(porcelain)
	ignoredPrivate := parseIgnoredPrivate(ignored)
	return Status{
		Branch:              branch,
		HeadSHA:             head,
		ShortSHA:            shortSHA(head),
		Upstream:            upstream,
		Origin:              origin,
		WorktreeClean:       len(tracked) == 0 && len(untracked) == 0,
		TrackedChanges:      trackedChanges(tracked),
		UntrackedFiles:      stringList(untracked),
		IgnoredPrivatePaths: stringList(ignoredPrivate),
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func git(root string, args ...string) (string, error) {
	output, err := gitRaw(root, args...)
	return strings.TrimSpace(output), err
}

func gitRaw(root string, args ...string) (string, error) {
	command := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", command...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), strings.TrimSpace(output.String()))
	}
	return output.String(), nil
}

func parsePorcelain(data string) ([]Change, []string) {
	var tracked []Change
	var untracked []string
	fields := strings.Split(data, "\x00")
	for index := 0; index < len(fields); index++ {
		field := fields[index]
		if len(field) < 4 {
			continue
		}
		code := field[:2]
		path := filepath.ToSlash(strings.TrimLeft(field[2:], " \t"))
		if code == "??" {
			untracked = append(untracked, path)
			continue
		}
		tracked = append(tracked, Change{Code: code, Path: path})
		if code[0] == 'R' || code[0] == 'C' {
			index++
		}
	}
	return tracked, untracked
}

func parseIgnoredPrivate(data string) []string {
	var paths []string
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "!! ") {
			continue
		}
		path := filepath.ToSlash(strings.TrimSpace(strings.TrimPrefix(line, "!! ")))
		if strings.HasPrefix(path, "data/private/") || strings.HasPrefix(path, "data/lake/") {
			paths = append(paths, path)
		}
	}
	return paths
}

func shortSHA(head string) string {
	if len(head) <= 12 {
		return head
	}
	return head[:12]
}

func trackedChanges(changes []Change) []Change {
	if changes == nil {
		return []Change{}
	}
	return changes
}

func stringList(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}
