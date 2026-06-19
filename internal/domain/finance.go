package domain

func BuildFinanceSummary(path string) (FinanceSummary, error) {
	state := newFinanceSummaryState()
	if err := readJSONL(path, state.record); err != nil {
		return FinanceSummary{}, err
	}
	return state.result(), nil
}

type financeSummaryState struct {
	summary         FinanceSummary
	currencies      map[string]bool
	categories      map[string]bool
	owners          map[string]*FinanceOwnerSummary
	ownerCurrencies map[string]map[string]bool
}

func newFinanceSummaryState() *financeSummaryState {
	return &financeSummaryState{
		currencies:      map[string]bool{},
		categories:      map[string]bool{},
		owners:          map[string]*FinanceOwnerSummary{},
		ownerCurrencies: map[string]map[string]bool{},
	}
}

func (state *financeSummaryState) result() FinanceSummary {
	state.summary.NetMinorUnits =
		state.summary.CreditMinorUnits - state.summary.DebitMinorUnits
	state.summary.Currency = summaryCurrency(state.currencies)
	state.summary.Categories = sortedKeys(state.categories)
	state.summary.OwnerBreakdown = financeOwnerBreakdown(
		state.owners,
		state.ownerCurrencies,
	)
	return state.summary
}
