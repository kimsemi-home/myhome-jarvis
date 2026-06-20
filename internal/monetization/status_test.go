package monetization

import "testing"

func TestMissingLedgerReturnsZeroRedactedStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.ExperimentCount != 0 || status.DecisionCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.LedgerPath != "data/private/monetization/experiments.jsonl" {
		t.Fatalf("ledger path = %s", status.LedgerPath)
	}
}

func TestStatusCountsExperimentDebtAndStates(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/monetization/experiments.jsonl", ledgerFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ExperimentCount != 2 || status.DecisionCount != 3 {
		t.Fatalf("status = %#v", status)
	}
	if status.ReviewRequiredCount != 1 || status.ExpectedValueUnknownCount != 1 {
		t.Fatalf("review/value counts = %#v", status)
	}
	if status.MissingEvidenceCount != 1 || status.MissingCostEstimateCount != 1 {
		t.Fatalf("missing counts = %#v", status)
	}
	if status.InvalidRecordCount != 1 || status.MonetizationDebtCount != 5 {
		t.Fatalf("debt counts = %#v", status)
	}
}
