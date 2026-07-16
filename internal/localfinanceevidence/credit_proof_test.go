package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedCreditProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	tampered := bytes.Replace(body, []byte(`"net_card_spend_minor": 18700`), []byte(`"net_card_spend_minor": 18701`), 1)
	if bytes.Equal(body, tampered) {
		t.Fatal("credit proof fixture did not contain expected value")
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{Component: "ledger", Capability: "credit-collection-rehearsal",
		ProofSchema: creditProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
		ReportHash: "577b96ea4bc4aa2bdd90c77c8568a4f2864dd3de219374e65a89eba2d5464079"}
	if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
		t.Fatal("expected tampered credit proof to fail")
	}
}

func TestDisabledOAuthBoundaryFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	tampered := bytes.Replace(body, []byte(`"official_origin_pinned": true`), []byte(`"official_origin_pinned": false`), 1)
	if bytes.Equal(body, tampered) {
		t.Fatal("credit proof fixture did not contain OAuth origin proof")
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{Component: "ledger", Capability: "credit-collection-rehearsal",
		ProofSchema: creditProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
		ReportHash: "577b96ea4bc4aa2bdd90c77c8568a4f2864dd3de219374e65a89eba2d5464079"}
	if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
		t.Fatal("expected disabled OAuth boundary to fail")
	}
}
