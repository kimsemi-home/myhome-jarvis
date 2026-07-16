package localfinanceevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRehashedUnsafeCreditBatchProofFails(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-loopback.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report CreditReport
	if err := json.Unmarshal(body, &report); err != nil {
		t.Fatal(err)
	}
	batch := &report.ImportTemplate.Onboarding.BatchPreview
	batch.RawFileNamesReported = true
	batch.BatchHash = ""
	unsignedBatch, _ := json.Marshal(batch)
	batch.BatchHash = digest(string(unsignedBatch))
	resealCreditTemplateAndReport(t, &report)
	tampered, _ := json.Marshal(report)
	ref := ProofRef{Component: "ledger", Capability: "credit-collection-rehearsal", ProofSchema: creditProofSchema,
		Path: "proof.json", ArtifactSHA256: digest(string(tampered)), ReportHash: report.ReportHash}
	if err := validateProofBody(tampered, "2026-07", ref); err == nil {
		t.Fatal("accepted rehashed batch proof that claimed raw filenames")
	}
}
