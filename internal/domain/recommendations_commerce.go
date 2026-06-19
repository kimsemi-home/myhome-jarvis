package domain

func appendRecurringPurchaseRecommendations(
	items []RecommendationItem,
	commerce CommerceSummary,
) []RecommendationItem {
	for _, candidate := range commerce.RecurringCandidates {
		items = append(items, RecommendationItem{
			Kind:                       "recurring_purchase_review",
			Title:                      "Compare recurring purchase: " + candidate.ItemName,
			Rationale:                  candidate.MerchantName + " appears repeatedly in local purchase fixtures.",
			Score:                      recurringPurchaseScore(candidate),
			Currency:                   candidate.Currency,
			EstimatedMonthlyMinorUnits: candidate.LatestTotalMinorUnits,
			EvidenceCount:              candidate.PurchaseCount,
		})
	}
	return items
}

func recurringPurchaseScore(candidate RecurringPurchaseCandidate) int {
	return clampScore(
		50 + candidate.PurchaseCount*10 +
			int(candidate.LatestTotalMinorUnits/1_000),
	)
}
