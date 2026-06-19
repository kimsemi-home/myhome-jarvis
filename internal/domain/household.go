package domain

func BuildHouseholdSummary(
	finance FinanceSummary,
	commerce CommerceSummary,
) HouseholdSummary {
	financeOwners := mapFinanceOwners(finance.OwnerBreakdown)
	commerceOwners := mapCommerceOwners(commerce.OwnerBreakdown)
	return HouseholdSummary{Scopes: []HouseholdScopeSummary{
		ownerScopeSummary(
			"user",
			"User",
			financeOwners["user"],
			commerceOwners["user"],
		),
		ownerScopeSummary(
			"spouse",
			"Spouse",
			financeOwners["spouse"],
			commerceOwners["spouse"],
		),
		householdScopeSummary(finance, commerce),
	}}
}

func householdScopeSummary(
	finance FinanceSummary,
	commerce CommerceSummary,
) HouseholdScopeSummary {
	return HouseholdScopeSummary{
		Scope:                   "household",
		Label:                   "Household",
		Currency:                firstCurrency(finance.Currency, commerce.Currency, "KRW"),
		FinanceRecords:          finance.Records,
		FinanceCreditMinorUnits: finance.CreditMinorUnits,
		FinanceDebitMinorUnits:  finance.DebitMinorUnits,
		FinanceNetMinorUnits:    finance.NetMinorUnits,
		PurchaseRecords:         commerce.Records,
		PurchaseSpendMinorUnits: commerce.TotalSpendMinorUnits,
	}
}
