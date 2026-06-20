package repofactory

import "testing"

func TestPolicyRequiresReviewAndPublicSafetyGates(t *testing.T) {
	policy := testPolicy()
	policy.AuthorityReviewRequired = false
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected missing authority review to fail")
	}

	policy = testPolicy()
	policy.CreationGates = policy.CreationGates[:1]
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected missing creation gates to fail")
	}
}

func TestPolicyRejectsUnreviewedRepoCreation(t *testing.T) {
	policy := testPolicy()
	policy.RepoCreationAllowedWithoutReview = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected unreviewed repo creation to fail")
	}
}

func TestPolicyRequiresTemplateRoles(t *testing.T) {
	policy := testPolicy()
	policy.TemplateFiles = policy.TemplateFiles[:2]
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected missing template roles to fail")
	}
}
