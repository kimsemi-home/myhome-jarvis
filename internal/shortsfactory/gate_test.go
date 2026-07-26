package shortsfactory

import (
	"path/filepath"
	"testing"
)

func TestPublicFixtureApprovesEvidenceGate(t *testing.T) {
	root := filepath.Join("..", "..")
	result, err := Verify(root)
	if err != nil {
		t.Fatal(err)
	}
	if result.Decision != "approved" || result.ReceiptRef == "" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGateRejectsMissingEvidenceAndDuplicateContent(t *testing.T) {
	root := filepath.Join("..", "..")
	contract, err := LoadContract(filepath.Join(root, "contracts", "shorts-factory.json"))
	if err != nil {
		t.Fatal(err)
	}
	request, err := LoadRequest(filepath.Join(root, "fixtures", "shorts_gate_pass.json"))
	if err != nil {
		t.Fatal(err)
	}
	request.ClaimEvidence.IndependentSources = 1
	request.Originality.CrossChannelDuplicate = true
	result, err := Evaluate(request, contract)
	if err != nil {
		t.Fatal(err)
	}
	if result.Decision != "rejected" || result.ReceiptRef != "" || len(result.ReleasedOpenLoopSteps) != 0 {
		t.Fatalf("unsafe request was released: %+v", result)
	}
}
