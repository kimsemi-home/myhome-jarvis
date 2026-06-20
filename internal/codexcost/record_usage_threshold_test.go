package codexcost

import "testing"

func TestRecordUsagePromotesReviewThreshold(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordUsage(root, []byte(`{
		"scope":"repo",
		"unit_kind":"codex_tokens",
		"amount":600000,
		"status":"recorded",
		"evidence_refs":["data/private/quality/runs.jsonl"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "review_required" ||
		result.BudgetState != "review_required" {
		t.Fatalf("result = %#v", result)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ReviewRequiredCount != 1 || status.RecordCount != 1 {
		t.Fatalf("status = %#v", status)
	}
}
