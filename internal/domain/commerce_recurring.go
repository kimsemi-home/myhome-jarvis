package domain

import (
	"sort"
	"strings"
)

func (state *commerceSummaryState) recordRecurringPurchase(
	purchase commercePurchase,
) {
	key := strings.Join([]string{
		purchase.MerchantName,
		purchase.ItemName,
		purchase.TotalPrice.Currency,
	}, "\x00")
	candidate, ok := state.groups[key]
	if !ok {
		candidate = &RecurringPurchaseCandidate{
			MerchantName:          purchase.MerchantName,
			ItemName:              purchase.ItemName,
			Currency:              purchase.TotalPrice.Currency,
			LatestTotalMinorUnits: purchase.TotalPrice.MinorUnits,
			LatestPurchasedAt:     purchase.PurchasedAt,
		}
		state.groups[key] = candidate
	}
	candidate.PurchaseCount++
	if purchase.PurchasedAt > candidate.LatestPurchasedAt {
		candidate.LatestPurchasedAt = purchase.PurchasedAt
		candidate.LatestTotalMinorUnits = purchase.TotalPrice.MinorUnits
	}
}

func (state *commerceSummaryState) addRecurringCandidates() {
	for _, candidate := range state.groups {
		if candidate.PurchaseCount >= 2 {
			state.summary.RecurringCandidates = append(
				state.summary.RecurringCandidates,
				*candidate,
			)
		}
	}
	sort.Slice(state.summary.RecurringCandidates, func(i, j int) bool {
		left := state.summary.RecurringCandidates[i]
		right := state.summary.RecurringCandidates[j]
		if left.MerchantName == right.MerchantName {
			return left.ItemName < right.ItemName
		}
		return left.MerchantName < right.MerchantName
	})
}
