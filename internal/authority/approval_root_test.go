package authority

import (
	"os"
	"path/filepath"
	"testing"
)

func writeApprovalRootPolicies(t *testing.T, root string) {
	t.Helper()
	writePolicy(t, root, testPolicy())
	copyRepoFile(t, root, "generated/context_pack.generated.json")
	copyRepoFile(t, root, "generated/external_evidence.generated.json")
}

func copyRepoFile(t *testing.T, root string, rel string) {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(repoRoot(t), filepath.FromSlash(rel)))
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, body, 0o600); err != nil {
		t.Fatal(err)
	}
}
