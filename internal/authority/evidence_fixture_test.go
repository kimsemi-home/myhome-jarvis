package authority

import (
	"os"
	"path/filepath"
	"testing"
)

func writeReviewRequestableEvidence(t *testing.T, root string) {
	t.Helper()
	writeTestFile(t, root, "docs/evidence-fixture.md", "# Evidence Fixture\n")
	writeTestFile(t, root, "data/private/quality/runs.jsonl", `{"at":"2026-06-18T00:01:00Z","ok":true}`+"\n")
	writeTestFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_1","at":"2026-06-18T00:00:00Z","status":"closed","evidence_refs":["docs/evidence-fixture.md"]}`+"\n")
}

func writeTestFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
