package localfinanceevidence

import (
	"path/filepath"
	"testing"
)

func TestFixtureManifestValidates(t *testing.T) {
	path := filepath.Join("..", "..", "fixtures", "local_finance", "manifest.json")
	manifest, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	if manifest.Month != "2026-07" || len(manifest.Receipts) != 4 || len(manifest.ExecutionProofs) != 1 {
		t.Fatalf("unexpected manifest: %#v", manifest)
	}
}

func TestTamperedReceiptFails(t *testing.T) {
	path := filepath.Join("..", "..", "fixtures", "local_finance", "manifest.json")
	manifest, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	manifest.Receipts[0].ArtifactHash = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	if err := Validate(manifest); err == nil {
		t.Fatal("expected tampered artifact hash to fail")
	}
}
