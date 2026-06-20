package codexsustainability

import "testing"

func TestRecordProposalRejectsInvalidPayloads(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, payload := range []string{
		`{"proposal_id":"opt","cost_per_accepted_change":1,"median_cycle_minutes":1,"cache_savings_units":0,"defect_rework_rate":0,"monetization_ref":"KIM-132","evidence_refs":[]}`,
		`{"proposal_id":"opt","cost_per_accepted_change":1,"median_cycle_minutes":1,"cache_savings_units":0,"defect_rework_rate":1.1,"monetization_ref":"KIM-132","evidence_refs":["docs/codex-sustainability.md"]}`,
		`{"proposal_id":"opt","cost_per_accepted_change":1,"median_cycle_minutes":1,"cache_savings_units":0,"defect_rework_rate":0,"monetization_ref":"https://example.invalid","evidence_refs":["docs/codex-sustainability.md"]}`,
		`{"proposal_id":"opt","cost_per_accepted_change":1,"median_cycle_minutes":1,"cache_savings_units":0,"defect_rework_rate":0,"monetization_ref":"KIM-132","raw_prompt":"private","evidence_refs":["docs/codex-sustainability.md"]}`,
	} {
		if _, err := RecordProposal(root, []byte(payload)); err == nil {
			t.Fatalf("expected error for %s", payload)
		}
	}
}

func TestRecordProposalReportsReviewGateState(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordProposal(root, []byte(`{
		"proposal_id":"expensive-loop",
		"cost_per_accepted_change":700000,
		"median_cycle_minutes":30,
		"cache_savings_units":0,
		"defect_rework_rate":0,
		"monetization_ref":"KIM-132",
		"evidence_refs":["docs/codex-sustainability.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.ReviewState != "review_required" || result.ReviewGateCount == 0 {
		t.Fatalf("result = %#v", result)
	}
}
