package commandcenter

import "testing"

func TestNextSafeActionRequestsAuthorityReviewWhenRequestable(t *testing.T) {
	status := Status{
		PublicSafe: true,
		AuthorityReview: AuthorityReviewSummary{
			ReviewRequestable: true,
		},
		BlockedGates: []GateSummary{{Key: "authority"}},
	}

	if got := nextSafeAction(status); got != "request_authority_review" {
		t.Fatalf("next safe action = %q", got)
	}
}
