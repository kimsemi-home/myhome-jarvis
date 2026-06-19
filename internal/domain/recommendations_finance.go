package domain

func appendSubscriptionRecommendation(
	items []RecommendationItem,
	finance FinanceSummary,
) []RecommendationItem {
	if finance.SubscriptionMinorUnits <= 0 {
		return items
	}
	return append(items, RecommendationItem{
		Kind:                       "subscription_review",
		Title:                      "Review household subscriptions",
		Rationale:                  "Subscription-like debit fixtures exist; keep this as a review-only recommendation.",
		Score:                      clampScore(55 + int(finance.SubscriptionMinorUnits/10_000)),
		Currency:                   finance.Currency,
		EstimatedMonthlyMinorUnits: finance.SubscriptionMinorUnits,
		EvidenceCount:              finance.SubscriptionCount,
	})
}

func appendCardRecommendation(
	items []RecommendationItem,
	finance FinanceSummary,
) []RecommendationItem {
	if finance.CardDebitMinorUnits <= 0 {
		return items
	}
	return append(items, RecommendationItem{
		Kind:                       "card_usage_review",
		Title:                      "Review card-linked household spend",
		Rationale:                  "Card-linked debit fixtures exist; keep this as a review-only recommendation, not a card action.",
		Score:                      clampScore(52 + int(finance.CardDebitMinorUnits/10_000)),
		Currency:                   finance.Currency,
		EstimatedMonthlyMinorUnits: finance.CardDebitMinorUnits,
		EvidenceCount:              finance.CardDebitCount,
	})
}
