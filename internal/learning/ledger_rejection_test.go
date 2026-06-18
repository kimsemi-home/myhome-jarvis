package learning

import "testing"

func TestRecordRejectsMissingEvidence(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"gap","owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected missing evidence refs to fail")
	}
}

func TestRecordRejectsSensitiveText(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"Bearer abc123 appeared in output.","evidence_refs":["data/private/quality/runs.jsonl"],"owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected sensitive text to fail")
	}
}

func TestRecordRejectsAbsoluteEvidenceRef(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"gap","evidence_refs":["/tmp/evidence.jsonl"],"owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected absolute evidence ref to fail")
	}
}
