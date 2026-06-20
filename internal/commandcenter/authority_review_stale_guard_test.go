package commandcenter

import "testing"

func TestStatusIncludesAuthorityReviewStaleGuard(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	review := status.AuthorityReview
	if review.ReviewRequestStaleAfterHours != 24 ||
		review.ReviewRequestEscalationAction == "" ||
		review.ReviewRequestAgeHours < 0 {
		t.Fatalf("authority review stale guard = %#v", review)
	}
	if review.ReviewRequestStale &&
		review.NextSafeAction != review.ReviewRequestEscalationAction {
		t.Fatalf("authority review escalation = %#v", review)
	}
}
