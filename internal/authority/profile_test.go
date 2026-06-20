package authority

import "testing"

func TestAssessSummarizesAssistantAuthorityProfiles(t *testing.T) {
	status := Assess(testPolicy(), clearInputs())

	if status.ProfileCount != 6 || status.SelfApprovalBlockedProfileCount != 6 {
		t.Fatalf("profile counts = %#v", status)
	}
	if status.ReviewRequiredProfileCount != 4 {
		t.Fatalf("review profile count = %#v", status)
	}
	for _, key := range []string{
		"household_finance_copilot",
		"shorts_factory_control_plane",
		"monetization_console",
		"self_improvement_loop",
	} {
		if !contains(status.ReviewRequiredProfiles, key) {
			t.Fatalf("review profiles missing %q in %#v", key, status.ReviewRequiredProfiles)
		}
	}
	if !contains(status.PublicSafetyGatedProfiles, "shorts_factory_control_plane") {
		t.Fatalf("public safety profiles = %#v", status.PublicSafetyGatedProfiles)
	}
}

func TestPolicyRejectsUnsafeAssistantAuthorityProfiles(t *testing.T) {
	policy := testPolicy()
	policy.AssistantAuthorityProfiles[1].RequiresHumanReview = false
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected finance profile without review to fail")
	}

	policy = testPolicy()
	policy.AssistantAuthorityProfiles[0].SelfApprovalAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected self approval profile to fail")
	}
}
