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
	for _, replacement := range []struct{ old, new string }{
		{`"total_value_minor": 200000`, `"total_value_minor": 200001`},
		{`"official_origin_pinned": true`, `"official_origin_pinned": false`},
	} {
		tampered := bytes.Replace(body, []byte(replacement.old), []byte(replacement.new), 1)
		if bytes.Equal(body, tampered) {
			t.Fatal("Portfolio proof fixture did not contain expected value")
		}
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
			t.Fatal(err)
		}
		ref := ProofRef{Component: "portfolio", Capability: "readonly-collection-rehearsal",
			ProofSchema: portfolioProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
			ReportHash: "4ad3677cc2c488d7fef9e526dc9134d0250983ef2e077bf03567e9332f53ce03"}
		if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
			t.Fatal("expected tampered Portfolio proof to fail")
		}
	}
}
