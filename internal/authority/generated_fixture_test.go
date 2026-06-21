package authority

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func copyGeneratedFixtures(t *testing.T, root string) {
	t.Helper()
	sourceRoot := filepath.Join(repoRoot(t), "generated")
	targetRoot := filepath.Join(root, "generated")
	if err := filepath.WalkDir(sourceRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(sourceRoot, path)
		if err != nil {
			return err
		}
		target := filepath.Join(targetRoot, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o700)
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, body, 0o600)
	}); err != nil {
		t.Fatal(err)
	}
}
