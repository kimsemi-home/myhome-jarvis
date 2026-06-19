package commands

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func RunFinanceHarness(root string) HarnessReport {
	report := HarnessReport{Name: "finance", Passed: true}
	summary, err := domain.BuildSummary(root)
	if err != nil {
		report.addCheck("finance summary builds", false, err.Error())
		return report
	}
	finance := summary.Finance
	report.addCheck("finance fixture records", finance.Records == 3, fmt.Sprintf("records=%d", finance.Records))
	report.addCheck("finance currency", finance.Currency == "KRW", "currency="+finance.Currency)
	report.addCheck("finance credit total", finance.CreditMinorUnits == 4_500_000, fmt.Sprintf("credit=%d", finance.CreditMinorUnits))
	report.addCheck("finance debit total", finance.DebitMinorUnits == 153_200, fmt.Sprintf("debit=%d", finance.DebitMinorUnits))
	report.addCheck("finance net total", finance.NetMinorUnits == 4_346_800, fmt.Sprintf("net=%d", finance.NetMinorUnits))
	report.addCheck("subscription review candidates", finance.SubscriptionCount == 1 && finance.SubscriptionMinorUnits == 65_900, fmt.Sprintf("subscriptions=%d total=%d", finance.SubscriptionCount, finance.SubscriptionMinorUnits))
	report.addCheck("card-linked debit review candidates", finance.CardDebitCount == 2 && finance.CardDebitMinorUnits == 153_200, fmt.Sprintf("card_debits=%d total=%d", finance.CardDebitCount, finance.CardDebitMinorUnits))

	user := financeOwner(finance.OwnerBreakdown, "user")
	report.addCheck("user finance scope", user != nil && user.Records == 1 && user.NetMinorUnits == -87_300, financeOwnerMessage(user))
	household := financeOwner(finance.OwnerBreakdown, "household")
	report.addCheck("household finance scope", household != nil && household.Records == 2 && household.NetMinorUnits == 4_434_100, financeOwnerMessage(household))
	return report
}

func financeOwner(items []domain.FinanceOwnerSummary, owner string) *domain.FinanceOwnerSummary {
	for index := range items {
		if items[index].Owner == owner {
			return &items[index]
		}
	}
	return nil
}

func financeOwnerMessage(owner *domain.FinanceOwnerSummary) string {
	if owner == nil {
		return "missing owner"
	}
	return fmt.Sprintf("%s records=%d net=%d", owner.Owner, owner.Records, owner.NetMinorUnits)
}
