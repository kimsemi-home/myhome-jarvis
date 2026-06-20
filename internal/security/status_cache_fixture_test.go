package security

import (
	"os"
	"path/filepath"
	"testing"
)

func cleanCommittedRepo(t *testing.T) string {
	t.Helper()
	root := initGitRepo(t)
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("ok\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "initial")
	return root
}
