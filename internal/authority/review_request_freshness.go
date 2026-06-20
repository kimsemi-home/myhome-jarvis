package authority

import "time"

const reviewRequestStaleAfterHours = 24

func applyReviewRequestFreshness(plan *ReviewPlanStatus, now time.Time) {
	plan.ReviewRequestStaleAfterHours = reviewRequestStaleAfterHours
	plan.ReviewRequestEscalationAction = "none"
	if !plan.ReviewRequestRecorded || plan.ReviewRequestLastRecordedAt == "" {
		return
	}
	recordedAt, err := time.Parse(time.RFC3339, plan.ReviewRequestLastRecordedAt)
	if err != nil {
		markStaleReviewRequest(plan, "refresh_authority_review_request")
		return
	}
	age := now.Sub(recordedAt)
	if age < 0 {
		age = 0
	}
	plan.ReviewRequestAgeHours = int(age / time.Hour)
	if plan.ReviewRequestAgeHours >= reviewRequestStaleAfterHours {
		markStaleReviewRequest(plan, "escalate_human_authority_review")
	}
}

func markStaleReviewRequest(plan *ReviewPlanStatus, action string) {
	plan.ReviewRequestStale = true
	plan.ReviewRequestEscalationAction = action
	if plan.NextSafeAction == "await_human_authority_review" {
		plan.NextSafeAction = action
	}
}
