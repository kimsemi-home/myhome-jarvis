package cicache

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFixture(t *testing.T, root string, workflow string) {
	t.Helper()
	writeFile(t, root, graphPath, graphFixture())
	writeFile(t, root, workflowPath, workflow)
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func replaceAll(body string, old string, next string) string {
	return strings.ReplaceAll(body, old, next)
}
