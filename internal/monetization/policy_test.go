package monetization

import "testing"

func TestPolicyRequiresPrivateRedactedLedger(t *testing.T) {
	policy := testPolicy()
	if err := validatePolicy(policy); err != nil {
		t.Fatal(err)
	}
	policy.RawRevenuePublicAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected raw revenue exposure to be rejected")
	}
}
