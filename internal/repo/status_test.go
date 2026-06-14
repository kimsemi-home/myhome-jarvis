package repo

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestInspectReportsCleanAndDirtyState(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init", "-b", "main")
	runGit(t, root, "config", "user.name", "Test User")
	runGit(t, root, "config", "user.email", "test@example.invalid")
	writeFile(t, root, ".gitignore", "data/private/*\ndata/lake/**\n")
	writeFile(t, root, "tracked.txt", "initial\n")
	runGit(t, root, "add", ".gitignore", "tracked.txt")
	runGit(t, root, "commit", "-m", "initial")

	clean, err := Inspect(root)
	if err != nil {
		t.Fatal(err)
	}
	if clean.Branch != "main" || clean.HeadSHA == "" || !clean.WorktreeClean {
		t.Fatalf("unexpected clean status: %#v", clean)
	}

	writeFile(t, root, "tracked.txt", "changed\n")
	writeFile(t, root, "new.txt", "new\n")
	writeFile(t, root, filepath.Join("data", "private", "token.txt"), "private\n")
	dirty, err := Inspect(root)
	if err != nil {
		t.Fatal(err)
	}
	if dirty.WorktreeClean {
		t.Fatalf("expected dirty worktree: %#v", dirty)
	}
	if len(dirty.TrackedChanges) != 1 || dirty.TrackedChanges[0].Path != "tracked.txt" {
		t.Fatalf("tracked changes = %#v", dirty.TrackedChanges)
	}
	if len(dirty.UntrackedFiles) != 1 || dirty.UntrackedFiles[0] != "new.txt" {
		t.Fatalf("untracked files = %#v", dirty.UntrackedFiles)
	}
	if len(dirty.IgnoredPrivatePaths) != 1 || dirty.IgnoredPrivatePaths[0] != "data/private/" {
		t.Fatalf("ignored private paths = %#v", dirty.IgnoredPrivatePaths)
	}
}

func runGit(t *testing.T, root string, args ...string) {
	t.Helper()
	command := append([]string{"-C", root}, args...)
	cmd := exec.Command("git", command...)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v failed: %v\n%s", args, err, output)
	}
}

func writeFile(t *testing.T, root string, path string, data string) {
	t.Helper()
	fullPath := filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(fullPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
}
