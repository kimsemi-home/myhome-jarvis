package authority

import (
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/review"
)

func TestAssessRequiresReviewWhenAuthorityDebtExists(t *testing.T) {
	inputs := clearInputs()
	inputs.EvidenceQuality.ReassessmentDebtCount = 2
	inputs.Incidents.IncidentDebtCount = 1

	status := Assess(testPolicy(), inputs)
	if status.Outcome != "review_required" || status.ActiveRule != "evidence_quality_debt" {
		t.Fatalf("status = %#v", status)
	}
	if status.AuthorityDebtCount != 3 || status.EvidenceQualityDebtCount != 2 || status.IncidentDebtCount != 1 {
		t.Fatalf("debt counts = %#v", status)
	}
	if contains(status.AllowedDecisions, "low_risk_fixture_change") {
		t.Fatalf("fixture change should require review while debt exists: %#v", status.AllowedDecisions)
	}
}

func TestAssessRequiresReviewWhenHumanReviewOverloaded(t *testing.T) {
	inputs := clearInputs()
	inputs.Review = review.Status{
		CapacityState:   "overloaded",
		ReviewDebtCount: 1,
	}

	status := Assess(testPolicy(), inputs)
	if status.Outcome != "review_required" || status.ActiveRule != "human_review_overloaded" {
		t.Fatalf("status = %#v", status)
	}
	if status.AuthorityDebtCount != 1 || status.HumanReviewDebtCount != 1 {
		t.Fatalf("debt counts = %#v", status)
	}
	if contains(status.AllowedDecisions, "low_risk_fixture_change") {
		t.Fatalf("fixture change should require review while reviewers are overloaded: %#v", status.AllowedDecisions)
	}
}
