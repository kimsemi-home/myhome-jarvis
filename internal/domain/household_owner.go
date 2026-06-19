package domain

func ownerScopeSummary(
	scope string,
	label string,
	finance *FinanceOwnerSummary,
	commerce *CommerceOwnerSummary,
) HouseholdScopeSummary {
	var summary HouseholdScopeSummary
	summary.Scope = scope
	summary.Label = label
	summary.Currency = "KRW"
	if finance != nil {
		applyFinanceOwnerScope(&summary, finance)
	}
	if commerce != nil {
		applyCommerceOwnerScope(&summary, commerce)
	}
	return summary
}

func applyFinanceOwnerScope(
	summary *HouseholdScopeSummary,
	finance *FinanceOwnerSummary,
) {
	summary.Currency = firstCurrency(finance.Currency, summary.Currency)
	summary.FinanceRecords = finance.Records
	summary.FinanceCreditMinorUnits = finance.CreditMinorUnits
	summary.FinanceDebitMinorUnits = finance.DebitMinorUnits
	summary.FinanceNetMinorUnits = finance.NetMinorUnits
}

func applyCommerceOwnerScope(
	summary *HouseholdScopeSummary,
	commerce *CommerceOwnerSummary,
) {
	summary.Currency = firstCurrency(summary.Currency, commerce.Currency)
	summary.PurchaseRecords = commerce.Records
	summary.PurchaseSpendMinorUnits = commerce.PurchaseSpendMinorUnits
}
