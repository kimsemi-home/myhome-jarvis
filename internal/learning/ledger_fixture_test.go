package learning

import (
	"os"
	"path/filepath"
	"testing"
)

func copyPolicyFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	source := filepath.Join(repoRoot(t), filepath.FromSlash(PolicyRelativePath))
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, body, 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
