package domain

import "sort"

func BuildRecommendationsSummary(
	finance FinanceSummary,
	commerce CommerceSummary,
) RecommendationsSummary {
	items := []RecommendationItem{}
	items = appendCashBufferRecommendation(items, finance)
	items = appendSubscriptionRecommendation(items, finance)
	items = appendCardRecommendation(items, finance)
	items = appendRecurringPurchaseRecommendations(items, commerce)
	sort.Slice(items, func(i, j int) bool {
		if items[i].Score == items[j].Score {
			return items[i].Title < items[j].Title
		}
		return items[i].Score > items[j].Score
	})
	return RecommendationsSummary{Count: len(items), Items: items}
}

func appendCashBufferRecommendation(
	items []RecommendationItem,
	finance FinanceSummary,
) []RecommendationItem {
	if finance.NetMinorUnits <= 0 {
		return items
	}
	return append(items, RecommendationItem{
		Kind:                       "cash_buffer",
		Title:                      "Keep household cash buffer",
		Rationale:                  "Fixture cashflow is positive; reserve surplus before recommendations become executable.",
		Score:                      clampScore(45 + int(finance.NetMinorUnits/1_000_000)),
		Currency:                   finance.Currency,
		EstimatedMonthlyMinorUnits: finance.NetMinorUnits,
		EvidenceCount:              finance.Records,
	})
}
