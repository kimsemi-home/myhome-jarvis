package codexsustainability

import "testing"

func TestPolicyRequiresPrivateVersionedLedger(t *testing.T) {
	policy := testPolicy()
	if err := validatePolicy(policy); err != nil {
		t.Fatal(err)
	}
	policy.PrivateEvidenceLedger = "generated/codex-sustainability.jsonl"
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected generated ledger to be rejected")
	}
	policy = testPolicy()
	policy.TrendBaselinesVersioned = false
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected unversioned trends to be rejected")
	}
}
