package domain

import (
	"fmt"
	"strings"
)

func (state *financeSummaryState) record(
	line int,
	transaction financeTransaction,
) error {
	owner := normalizeOwner(transaction.Owner)
	state.summary.Records++
	recordCurrency(state.currencies, transaction.Amount.Currency)
	recordCategory(state.categories, transaction.Category)
	ownerSummary := financeOwnerSummary(state.owners, owner)
	ownerSummary.Records++
	recordOwnerCurrency(
		state.ownerCurrencies,
		owner,
		transaction.Amount.Currency,
	)
	switch transaction.Direction {
	case "credit":
		state.recordFinanceCredit(ownerSummary, transaction.Amount.MinorUnits)
	case "debit":
		state.recordFinanceDebit(ownerSummary, transaction)
	default:
		return fmt.Errorf(
			"line %d: unknown direction %q",
			line,
			transaction.Direction,
		)
	}
	return nil
}

func (state *financeSummaryState) recordFinanceCredit(
	owner *FinanceOwnerSummary,
	amount int64,
) {
	state.summary.CreditMinorUnits += amount
	owner.CreditMinorUnits += amount
}

func (state *financeSummaryState) recordFinanceDebit(
	owner *FinanceOwnerSummary,
	transaction financeTransaction,
) {
	amount := transaction.Amount.MinorUnits
	state.summary.DebitMinorUnits += amount
	owner.DebitMinorUnits += amount
	if transaction.Category == "subscription" {
		state.summary.SubscriptionMinorUnits += amount
		state.summary.SubscriptionCount++
	}
	if strings.TrimSpace(transaction.CardID) != "" {
		state.summary.CardDebitMinorUnits += amount
		state.summary.CardDebitCount++
	}
}
