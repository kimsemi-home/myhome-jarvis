package domain

func (state *commerceSummaryState) record(
	_ int,
	purchase commercePurchase,
) error {
	owner := normalizeOwner(purchase.Owner)
	state.summary.Records++
	recordCurrency(state.currencies, purchase.TotalPrice.Currency)
	state.summary.TotalSpendMinorUnits += purchase.TotalPrice.MinorUnits
	recordCategory(state.categories, purchase.Category)
	ownerSummary := commerceOwnerSummary(state.owners, owner)
	ownerSummary.Records++
	ownerSummary.PurchaseSpendMinorUnits += purchase.TotalPrice.MinorUnits
	recordOwnerCurrency(
		state.ownerCurrencies,
		owner,
		purchase.TotalPrice.Currency,
	)
	if purchase.RecurringCandidate {
		state.recordRecurringPurchase(purchase)
	}
	return nil
}
