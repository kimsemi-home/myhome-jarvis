package mergeevidence

import "testing"

func TestPolicyRequiresEligibleMergeGates(t *testing.T) {
	policy := testPolicy()
	policy.Gates = policy.Gates[:2]
	if err := ValidatePolicy(policy); err == nil {
		t.Fatal("expected missing merge gates to fail")
	}
}

func TestPolicyRejectsUnreviewedMerge(t *testing.T) {
	policy := testPolicy()
	policy.MergeWithoutReviewAllowed = true
	if err := ValidatePolicy(policy); err == nil {
		t.Fatal("expected unreviewed merge policy to fail")
	}
}
