package domain

func commerceOwnerSummary(
	owners map[string]*CommerceOwnerSummary,
	owner string,
) *CommerceOwnerSummary {
	summary, ok := owners[owner]
	if !ok {
		summary = &CommerceOwnerSummary{Owner: owner}
		owners[owner] = summary
	}
	return summary
}

func commerceOwnerBreakdown(
	owners map[string]*CommerceOwnerSummary,
	currencies map[string]map[string]bool,
) []CommerceOwnerSummary {
	breakdown := make([]CommerceOwnerSummary, 0, len(owners))
	for _, owner := range ownerBreakdownOrder(owners) {
		summary := *owners[owner]
		summary.Currency = summaryCurrency(currencies[owner])
		breakdown = append(breakdown, summary)
	}
	return breakdown
}

func mapCommerceOwners(
	items []CommerceOwnerSummary,
) map[string]*CommerceOwnerSummary {
	mapped := map[string]*CommerceOwnerSummary{}
	for index := range items {
		mapped[items[index].Owner] = &items[index]
	}
	return mapped
}
