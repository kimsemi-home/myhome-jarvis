package authority

import (
	"testing"
	"time"
)

func TestReviewRequestFreshnessKeepsRecentReviewPending(t *testing.T) {
	plan := recordedReviewPlanAt(t, "2026-06-20T05:00:00Z")
	applyReviewRequestFreshness(&plan, parseTime(t, "2026-06-21T04:00:00Z"))
	if plan.ReviewRequestStale ||
		plan.ReviewRequestAgeHours != 23 ||
		plan.NextSafeAction != "await_human_authority_review" ||
		plan.ReviewRequestEscalationAction != "none" {
		t.Fatalf("freshness = %#v", plan)
	}
}

func TestReviewRequestFreshnessEscalatesStaleReview(t *testing.T) {
	plan := recordedReviewPlanAt(t, "2026-06-20T05:00:00Z")
	applyReviewRequestFreshness(&plan, parseTime(t, "2026-06-21T06:00:00Z"))
	if !plan.ReviewRequestStale ||
		plan.ReviewRequestAgeHours != 25 ||
		plan.NextSafeAction != "escalate_human_authority_review" ||
		plan.ReviewRequestEscalationAction != "escalate_human_authority_review" {
		t.Fatalf("freshness = %#v", plan)
	}
}

func recordedReviewPlanAt(t *testing.T, at string) ReviewPlanStatus {
	t.Helper()
	return ReviewPlanStatus{
		ReviewRequestRecorded:       true,
		ReviewRequestLastRecordedAt: at,
		NextSafeAction:              "await_human_authority_review",
	}
}

func parseTime(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatal(err)
	}
	return parsed
}
