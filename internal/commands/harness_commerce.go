package commands

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func RunCommerceHarness(root string) HarnessReport {
	report := HarnessReport{Name: "commerce", Passed: true}
	summary, err := domain.BuildSummary(root)
	if err != nil {
		report.addCheck("commerce summary builds", false, err.Error())
		return report
	}
	commerce := summary.Commerce
	report.addCheck("commerce fixture records", commerce.Records == 3, fmt.Sprintf("records=%d", commerce.Records))
	report.addCheck("commerce currency", commerce.Currency == "KRW", "currency="+commerce.Currency)
	report.addCheck("commerce spend total", commerce.TotalSpendMinorUnits == 26_800, fmt.Sprintf("spend=%d", commerce.TotalSpendMinorUnits))
	report.addCheck("recurring purchase candidates", commerce.RecurringCandidateCount == 1, fmt.Sprintf("recurring=%d", commerce.RecurringCandidateCount))
	checkRecurringPurchase(&report, commerce)

	user := commerceOwner(commerce.OwnerBreakdown, "user")
	report.addCheck("user commerce scope", user != nil && user.Records == 1 && user.PurchaseSpendMinorUnits == 3_200, commerceOwnerMessage(user))
	household := commerceOwner(commerce.OwnerBreakdown, "household")
	report.addCheck("household commerce scope", household != nil && household.Records == 2 && household.PurchaseSpendMinorUnits == 23_600, commerceOwnerMessage(household))
	return report
}

func checkRecurringPurchase(report *HarnessReport, commerce domain.CommerceSummary) {
	if len(commerce.RecurringCandidates) == 0 {
		report.addCheck("recurring purchase candidate detail", false, "missing recurring candidate")
		return
	}
	candidate := commerce.RecurringCandidates[0]
	passed := candidate.MerchantName == "Coupang" &&
		candidate.ItemName == "Bottled water 2L x 6" &&
		candidate.PurchaseCount == 2 &&
		candidate.LatestTotalMinorUnits == 11_800
	message := fmt.Sprintf("%s %s count=%d total=%d", candidate.MerchantName, candidate.ItemName, candidate.PurchaseCount, candidate.LatestTotalMinorUnits)
	report.addCheck("recurring purchase candidate detail", passed, message)
}

func commerceOwner(items []domain.CommerceOwnerSummary, owner string) *domain.CommerceOwnerSummary {
	for index := range items {
		if items[index].Owner == owner {
			return &items[index]
		}
	}
	return nil
}

func commerceOwnerMessage(owner *domain.CommerceOwnerSummary) string {
	if owner == nil {
		return "missing owner"
	}
	return fmt.Sprintf("%s records=%d spend=%d", owner.Owner, owner.Records, owner.PurchaseSpendMinorUnits)
}
