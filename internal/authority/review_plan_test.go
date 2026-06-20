package authority

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReviewPlanRequestsAuthorityReviewWithoutApproval(t *testing.T) {
	policy := testPolicy()
	status := Assess(policy, clearInputs())
	plan := ReviewPlan(policy, status)

	if !plan.PublicSafe || !plan.ReviewRequestable {
		t.Fatalf("review plan = %#v", plan)
	}
	if plan.NextSafeAction != "request_authority_review" {
		t.Fatalf("next action = %q", plan.NextSafeAction)
	}
	if plan.HighRiskBlockedDecisionCount != 6 || plan.ReviewRequiredDecisionCount != 6 {
		t.Fatalf("decision counts = %#v", plan)
	}
	if plan.ExternalWritesAllowedProfileCount != 0 || plan.SelfApprovalBlockedProfileCount != 6 {
		t.Fatalf("approval boundaries = %#v", plan)
	}
}

func TestReviewPlanReflectsReviewCapacity(t *testing.T) {
	policy := testPolicy()
	inputs := clearInputs()
	inputs.Review.CapacityState = "overloaded"
	status := Assess(policy, inputs)
	plan := ReviewPlan(policy, status)

	if plan.ReviewRequestable || plan.NextSafeAction != "resolve_authority_debt" {
		t.Fatalf("review plan = %#v", plan)
	}
}

func TestReviewPlanStatusRedactsPrivateFields(t *testing.T) {
	plan := ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs()))
	body, err := json.Marshal(plan)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_rationale", "raw_evidence", "evidence_refs",
		"reviewer_identity", "linear_url", "token", "credential",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("review plan leaked %q in %s", forbidden, body)
		}
	}
}
