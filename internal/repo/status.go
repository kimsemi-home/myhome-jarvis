package repo

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

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
