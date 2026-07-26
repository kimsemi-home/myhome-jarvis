package domain

import (
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func TestBuildFinanceSummaryFromLiveCardSource(t *testing.T) {
	summary, err := BuildFinanceSummaryFromSource([]financeconnector.SourceTransaction{
		{Owner: "spouse", Amount: 42000, Currency: "KRW", Direction: "debit", Category: "uncategorized", CardID: "card-1"},
		{Owner: "spouse", Amount: 10000, Currency: "KRW", Direction: "credit", Category: "uncategorized", CardID: "card-1"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Records != 2 || summary.DebitMinorUnits != 42000 || summary.CreditMinorUnits != 10000 || summary.CardDebitCount != 1 {
		t.Fatalf("summary = %#v", summary)
	}
}
