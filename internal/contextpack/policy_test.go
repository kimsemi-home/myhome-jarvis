package contextpack

import "testing"

func TestPolicyRejectsUnsafeContracts(t *testing.T) {
	policy := testPolicy()
	policy.AuthorityContract.SelfApprovalAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected self approval to be rejected")
	}
	policy = testPolicy()
	policy.SecurityContract.LocalPathsPublicAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected local path exposure to be rejected")
	}
}
