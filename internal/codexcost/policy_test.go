package codexcost

import "testing"

func TestPolicyRequiresPrivateRedactedLedger(t *testing.T) {
	policy := testPolicy()
	policy.PrivateUsageLedger = "generated/usage.jsonl"
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected generated ledger to be rejected")
	}
}
