package codexcost

import "testing"

func TestGuardLoopRejectsInvalidPayloads(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	for _, payload := range []string{
		`{"scope":"bad","unit_kind":"codex_tokens","estimated_units":1,"estimated_minutes":1,"evidence_refs":["docs/codex-cost-governor.md"]}`,
		`{"scope":"assistant_loop","unit_kind":"bad","estimated_units":1,"estimated_minutes":1,"evidence_refs":["docs/codex-cost-governor.md"]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","estimated_units":0,"estimated_minutes":1,"evidence_refs":["docs/codex-cost-governor.md"]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","estimated_units":1,"estimated_minutes":1,"evidence_refs":[]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","estimated_units":1,"estimated_minutes":1,"raw_prompt":"private","evidence_refs":["docs/codex-cost-governor.md"]}`,
		`{"scope":"assistant_loop","unit_kind":"codex_tokens","estimated_units":1,"estimated_minutes":1,"evidence_refs":["https://example.invalid"]}`,
	} {
		if _, err := GuardLoop(root, []byte(payload)); err == nil {
			t.Fatalf("expected error for %s", payload)
		}
	}
}
