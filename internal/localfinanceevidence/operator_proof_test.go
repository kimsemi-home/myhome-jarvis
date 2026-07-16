package localfinanceevidence

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTamperedOperatorProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "finance-operator-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	tampered := bytes.Replace(body, []byte(`"tracked_surplus_minor": -12400`), []byte(`"tracked_surplus_minor": -12399`), 1)
	if bytes.Equal(body, tampered) {
		t.Fatal("Finance Operator proof fixture did not contain expected value")
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "proof.json"), tampered, 0o600); err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{Component: "finance-operator", Capability: "monthly-orchestration-rehearsal",
		ProofSchema: operatorProofSchema, Path: "proof.json", ArtifactSHA256: digest(string(tampered)),
		ReportHash: "c5ae49b386b5b7e78a90bef22a7b56447cbf996a9da9647b60596fd9c770056c"}
	if err := validateProofFiles(root, "2026-07", []ProofRef{ref}); err == nil {
		t.Fatal("expected tampered Finance Operator proof to fail")
	}
}
