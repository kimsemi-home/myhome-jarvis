package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedRevenueProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "revenue-youtube-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	tampered := bytes.Replace(body, []byte(`"net_minor": 6300`), []byte(`"net_minor": 6301`), 1)
	if bytes.Equal(body, tampered) {
		t.Fatal("Revenue proof fixture did not contain expected value")
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{Component: "revenue", Capability: "youtube-revenue-collection-rehearsal",
		ProofSchema: revenueProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
		ReportHash: "a913f1b9498e45873d2a93ad29b42ab1f3a0b497b2495a55577c692a4c2a5ed3"}
	if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
		t.Fatal("expected tampered Revenue proof to fail")
	}
}
