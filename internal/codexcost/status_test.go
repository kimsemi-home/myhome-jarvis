package codexcost

import "testing"

func TestMissingLedgerReturnsZeroRedactedStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.RecordCount != 0 || status.TotalUnits != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.LedgerPath != "data/private/codex-cost/usage.jsonl" {
		t.Fatalf("ledger path = %s", status.LedgerPath)
	}
}

func TestStatusCountsUsageDebtAndBudgetState(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-18T00:00:00Z","scope":"assistant_loop","unit_kind":"codex_tokens","amount":60000,"status":"recorded","evidence_refs":["generated/verification_graph.generated.json"],"raw_prompt":"private"}`+"\n"+
			`{"at":"2026-06-19T00:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":50000,"status":"review_required","evidence_refs":["docs/assistant-vision.md"],"private_notes":"private"}`+"\n"+
			`{"at":"2026-06-19T01:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":1,"status":"recorded","evidence_refs":[]}`+"\n"+
			`{"at":"2026-06-19T02:00:00Z","scope":"bad","unit_kind":"codex_tokens","amount":1,"status":"recorded","evidence_refs":["docs/assistant-vision.md"]}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.RecordCount != 2 || status.TotalUnits != 110000 {
		t.Fatalf("status = %#v", status)
	}
	if status.BudgetState != "warning" || status.ReviewRequiredCount != 1 {
		t.Fatalf("budget = %#v", status)
	}
	if status.MissingEvidenceCount != 1 || status.InvalidRecordCount != 1 {
		t.Fatalf("debt = %#v", status)
	}
}
