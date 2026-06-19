package domain

func BuildCommerceSummary(path string) (CommerceSummary, error) {
	state := newCommerceSummaryState()
	if err := readJSONL(path, state.record); err != nil {
		return CommerceSummary{}, err
	}
	return state.result(), nil
}

type commerceSummaryState struct {
	summary         CommerceSummary
	currencies      map[string]bool
	categories      map[string]bool
	groups          map[string]*RecurringPurchaseCandidate
	owners          map[string]*CommerceOwnerSummary
	ownerCurrencies map[string]map[string]bool
}

func newCommerceSummaryState() *commerceSummaryState {
	return &commerceSummaryState{
		currencies:      map[string]bool{},
		categories:      map[string]bool{},
		groups:          map[string]*RecurringPurchaseCandidate{},
		owners:          map[string]*CommerceOwnerSummary{},
		ownerCurrencies: map[string]map[string]bool{},
	}
}

func (state *commerceSummaryState) result() CommerceSummary {
	state.addRecurringCandidates()
	state.summary.RecurringCandidateCount =
		len(state.summary.RecurringCandidates)
	state.summary.Currency = summaryCurrency(state.currencies)
	state.summary.Categories = sortedKeys(state.categories)
	state.summary.OwnerBreakdown = commerceOwnerBreakdown(
		state.owners,
		state.ownerCurrencies,
	)
	return state.summary
}
