package domain

func financeOwnerSummary(
	owners map[string]*FinanceOwnerSummary,
	owner string,
) *FinanceOwnerSummary {
	summary, ok := owners[owner]
	if !ok {
		summary = &FinanceOwnerSummary{Owner: owner}
		owners[owner] = summary
	}
	return summary
}

func financeOwnerBreakdown(
	owners map[string]*FinanceOwnerSummary,
	currencies map[string]map[string]bool,
) []FinanceOwnerSummary {
	breakdown := make([]FinanceOwnerSummary, 0, len(owners))
	for _, owner := range ownerBreakdownOrder(owners) {
		summary := *owners[owner]
		summary.NetMinorUnits =
			summary.CreditMinorUnits - summary.DebitMinorUnits
		summary.Currency = summaryCurrency(currencies[owner])
		breakdown = append(breakdown, summary)
	}
	return breakdown
}

func mapFinanceOwners(
	items []FinanceOwnerSummary,
) map[string]*FinanceOwnerSummary {
	mapped := map[string]*FinanceOwnerSummary{}
	for index := range items {
		mapped[items[index].Owner] = &items[index]
	}
	return mapped
}
