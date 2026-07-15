package repogovernance

import "testing"

func TestBidirectionalChangeGate(t *testing.T) {
	manifest := Manifest{Groups: []Group{{
		ID: "shorts", DocumentSources: []string{"docs-src"}, GeneratedDocuments: []string{"docs"},
		Code: []string{"internal"}, Tests: []string{"tests"}, ChangePolicy: "bidirectional",
	}}}
	if err := CheckChanges(manifest, []string{"docs-src/a.json", "docs/a.md", "internal/a.go"}); err != nil {
		t.Fatal(err)
	}
	if err := CheckChanges(manifest, []string{"internal/a.go"}); err == nil {
		t.Fatal("expected code-only change to fail")
	}
	if err := CheckChanges(manifest, []string{"docs-src/a.json", "docs/a.md"}); err == nil {
		t.Fatal("expected docs-only change to fail")
	}
}
