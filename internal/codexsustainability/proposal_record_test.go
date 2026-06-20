package codexsustainability

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRecordProposalAppendsEvidenceAndReturnsRedactedResult(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordProposal(root, []byte(`{
		"proposal_id":"codex-cache-hit-ratio",
		"cost_per_accepted_change":100000,
		"median_cycle_minutes":4,
		"cache_savings_units":25000,
		"defect_rework_rate":0.1,
		"monetization_ref":"KIM-132",
		"evidence_refs":["data/private/codex-cost/usage.jsonl"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.ProposalID != "codex-cache-hit-ratio" ||
		result.EvidenceRefCount != 1 {
		t.Fatalf("result = %#v", result)
	}
	if bytes.Contains(mustJSON(t, result), []byte("evidence_refs")) {
		t.Fatalf("result leaked evidence refs: %#v", result)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.FeatureProposalCount != 1 ||
		status.OptimizationClaimWithoutEvidenceCount != 0 {
		t.Fatalf("status = %#v", status)
	}
}

func mustJSON(t *testing.T, value any) []byte {
	t.Helper()
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return body
}
