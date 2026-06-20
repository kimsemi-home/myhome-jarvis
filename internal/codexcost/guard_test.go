package codexcost

import (
	"bytes"
	"testing"
)

func TestGuardLoopAllowsSmallLoopWhenEvidenceIsHealthy(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)

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
	if result.Decision != "allow" || len(result.Reasons) != 0 {
		t.Fatalf("result = %#v", result)
	}
	if result.EvidenceRefCount != 1 ||
		result.SustainabilityPosture != "sustainable" {
		t.Fatalf("redacted result = %#v", result)
	}
}

func TestGuardLoopDoesNotExposeEvidenceRefs(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	result, err := GuardLoop(root, []byte(`{
		"scope":"assistant_loop",
		"unit_kind":"codex_tokens",
		"estimated_units":1000,
		"estimated_minutes":5,
		"evidence_refs":["data/private/codex-cost/usage.jsonl"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	body := mustJSON(t, result)
	if bytes.Contains(body, []byte("evidence_refs")) ||
		bytes.Contains(body, []byte("data/private/codex-cost/usage.jsonl")) {
		t.Fatalf("guard leaked private refs: %s", body)
	}
}
