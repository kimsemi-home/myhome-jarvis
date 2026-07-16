package localfinanceevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRehashedUnsafeCreditBalanceProofFails(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-loopback.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report CreditReport
	if err := json.Unmarshal(body, &report); err != nil {
		t.Fatal(err)
	}
	preview := &report.ImportTemplate.Onboarding.Version1Preview
	preview.ExpectedBalance.ClosingMinor++
	preview.PreviewHash = ""
	unsignedPreview, _ := json.Marshal(preview)
	preview.PreviewHash = digest(string(unsignedPreview))
	resealCreditTemplateAndReport(t, &report)
	tampered, _ := json.Marshal(report)
	ref := ProofRef{Component: "ledger", Capability: "credit-collection-rehearsal", ProofSchema: creditProofSchema,
		Path: "proof.json", ArtifactSHA256: digest(string(tampered)), ReportHash: report.ReportHash}
	if err := validateProofBody(tampered, "2026-07", ref); err == nil {
		t.Fatal("accepted rehashed preview with a false closing balance")
	}
}
