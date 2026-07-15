package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckSkipsGitWorktreePointer(t *testing.T) {
	root := t.TempDir()
	pointer := "gitdir: /" + strings.Join([]string{"Us", "ers"}, "") + "/example/private/.git/worktrees/example\n"
	if err := os.WriteFile(filepath.Join(root, ".git"), []byte(pointer), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected Git worktree metadata to be skipped, got %+v", report.Findings)
	}
}
