package localfinanceevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRehashedUnsafeCreditTemplateProofFails(t *testing.T) {
	source := filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-loopback.json")
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	var report CreditReport
	if err := json.Unmarshal(body, &report); err != nil {
		t.Fatal(err)
	}
	report.ImportTemplate.Guards.TemplateVersionMutationRejected = false
	report.ImportTemplate.ReportHash = ""
	unsignedTemplate, err := json.Marshal(report.ImportTemplate)
	if err != nil {
		t.Fatal(err)
	}
	report.ImportTemplate.ReportHash = digest(string(unsignedTemplate))
	report.ReportHash = ""
	unsignedReport, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	report.ReportHash = digest(string(unsignedReport))
	tampered, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	ref := ProofRef{
		Component: "ledger", Capability: "credit-collection-rehearsal", ProofSchema: creditProofSchema,
		Path: "proof.json", ArtifactSHA256: digest(string(tampered)), ReportHash: report.ReportHash,
	}
	if err := validateProofBody(tampered, "2026-07", ref); err == nil {
		t.Fatal("accepted rehashed credit proof with a disabled template mutation guard")
	}
}
