package domain

import "testing"

func assertRecommendationSummary(t *testing.T, summary Summary) {
	t.Helper()
	if summary.Recommendations.Count != 4 {
		t.Fatalf("recommendation count = %d", summary.Recommendations.Count)
	}
	if summary.Recommendations.Items[0].Kind != "recurring_purchase_review" {
		t.Fatalf("top recommendation = %#v", summary.Recommendations.Items[0])
	}
	if summary.Recommendations.Items[0].Score < summary.Recommendations.Items[1].Score {
		t.Fatalf("recommendations are not ranked: %#v", summary.Recommendations.Items)
	}
	foundCardReview := false
	for _, item := range summary.Recommendations.Items {
		if item.Kind == "card_usage_review" {
			foundCardReview = true
			if item.EstimatedMonthlyMinorUnits != 153_200 || item.EvidenceCount != 2 {
				t.Fatalf("card recommendation = %#v", item)
			}
		}
	}
	if !foundCardReview {
		t.Fatalf("missing card recommendation: %#v", summary.Recommendations.Items)
	}
}
