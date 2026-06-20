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

func TestNextSafeActionAwaitsRecordedAuthorityReview(t *testing.T) {
	status := Status{
		PublicSafe: true,
		AuthorityReview: AuthorityReviewSummary{
			ReviewRequestable:     true,
			ReviewRequestRecorded: true,
			NextSafeAction:        "await_human_authority_review",
		},
		BlockedGates: []GateSummary{{Key: "authority"}},
	}

	if got := nextSafeAction(status); got != "await_human_authority_review" {
		t.Fatalf("next safe action = %q", got)
	}
}

func TestNextSafeActionRepairsStorageArchiveNoiseBudget(t *testing.T) {
	status := Status{
		PublicSafe:   true,
		BlockedGates: []GateSummary{{Key: "storage_archive"}},
	}

	if got := nextSafeAction(status); got != "repair_storage_archive_noise_budget" {
		t.Fatalf("next safe action = %q", got)
	}
}
