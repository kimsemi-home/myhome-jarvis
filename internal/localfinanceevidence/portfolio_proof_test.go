package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedPortfolioProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "portfolio-holdings-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	tampered := bytes.Replace(body, []byte(`"total_value_minor": 200000`), []byte(`"total_value_minor": 200001`), 1)
	if bytes.Equal(body, tampered) {
		t.Fatal("Portfolio proof fixture did not contain expected value")
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{Component: "portfolio", Capability: "readonly-collection-rehearsal",
		ProofSchema: portfolioProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
		ReportHash: "afbb35fcf06e22e7d6ab108562c50d297db65efebaa98eca3dee3d57f2715c51"}
	if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
		t.Fatal("expected tampered Portfolio proof to fail")
	}
}
