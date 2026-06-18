package evidence

import (
	"os"
	"path/filepath"
	"testing"
)

func writeEvidencePolicy(t *testing.T, root string, redacted bool) {
	t.Helper()
	rawPublic := "false"
	if !redacted {
		rawPublic = "true"
	}
	writeFile(t, root, PolicyRelativePath, evidencePolicyFixture(rawPublic))
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
