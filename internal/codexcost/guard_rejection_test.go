package codexcost

import "testing"

func TestGuardLoopRequiresReviewForProjectedReviewThreshold(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	result, err := GuardLoop(root, []byte(`{
		"scope":"assistant_loop",
		"unit_kind":"codex_tokens",
		"estimated_units":600000,
		"estimated_minutes":5,
		"evidence_refs":["docs/codex-cost-governor.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.Decision != "review_required" ||
		result.ProjectedBudgetState != "review_required" {
		t.Fatalf("result = %#v", result)
	}
}

func TestGuardLoopRequiresReviewForSustainabilityPosture(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	result, err := GuardLoop(root, []byte(`{
		"scope":"assistant_loop",
		"unit_kind":"codex_tokens",
		"estimated_units":1000,
		"estimated_minutes":5,
		"evidence_refs":["docs/codex-cost-governor.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.Decision != "review_required" ||
		result.SustainabilityPosture != "blocked" {
		t.Fatalf("result = %#v", result)
	}
}
