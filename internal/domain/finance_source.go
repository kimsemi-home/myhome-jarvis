package domain

import "github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"

func BuildFinanceSummaryFromSource(
	transactions []financeconnector.SourceTransaction,
) (FinanceSummary, error) {
	state := newFinanceSummaryState()
	for index, transaction := range transactions {
		category := transaction.Category
		if category == "" {
			category = "uncategorized"
		}
		if err := state.record(index+1, financeTransaction{
			Owner: transaction.Owner, Amount: moneyAmount{
				MinorUnits: transaction.Amount, Currency: transaction.Currency,
			}, Direction: transaction.Direction, Category: category,
			CardID: transaction.CardID,
		}); err != nil {
			return FinanceSummary{}, err
		}
	}
	return state.result(), nil
}
