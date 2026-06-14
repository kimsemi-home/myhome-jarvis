package commands

import (
	"fmt"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

type HarnessCase struct {
	Name       string `json:"name"`
	Command    string `json:"command"`
	Payload    string `json:"payload"`
	ShouldPass bool   `json:"should_pass"`
	Contains   string `json:"contains,omitempty"`
}

type HarnessCaseResult struct {
	Name    string `json:"name"`
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}

type HarnessReport struct {
	Name    string              `json:"name"`
	Passed  bool                `json:"passed"`
	Results []HarnessCaseResult `json:"results"`
}

func RunHomeHarness() HarnessReport {
	cases := []HarnessCase{
		{Name: "open_youtube empty payload success", Command: "open-youtube", Payload: `{}`, ShouldPass: true, Contains: "https://www.youtube.com"},
		{Name: "open_youtube_search lofi music success", Command: "open-youtube-search", Payload: `{"query":"lofi music"}`, ShouldPass: true, Contains: "search_query=lofi+music"},
		{Name: "open_ott netflix success", Command: "open-ott", Payload: `{"service":"netflix"}`, ShouldPass: true, Contains: "https://www.netflix.com"},
		{Name: "open_ott unknown fail", Command: "open-ott", Payload: `{"service":"unknown"}`, ShouldPass: false},
		{Name: "volume_set 30 success", Command: "volume-set", Payload: `{"level":30}`, ShouldPass: true, Contains: "30"},
		{Name: "volume_set 101 fail", Command: "volume-set", Payload: `{"level":101}`, ShouldPass: false},
		{Name: "volume_up step 10 success", Command: "volume-up", Payload: `{"step":10}`, ShouldPass: true, Contains: "+ 10"},
		{Name: "volume_down step 10 success", Command: "volume-down", Payload: `{"step":10}`, ShouldPass: true, Contains: "- 10"},
		{Name: "display_sleep success", Command: "display-sleep", Payload: `{}`, ShouldPass: true, Contains: "displaysleepnow"},
		{Name: "open_url https success", Command: "open-url", Payload: `{"url":"https://example.com"}`, ShouldPass: true, Contains: "https://example.com"},
		{Name: "open_url javascript fail", Command: "open-url", Payload: `{"url":"javascript:alert(1)"}`, ShouldPass: false},
		{Name: "movie_mode dry-run success", Command: "movie-mode", Payload: `{}`, ShouldPass: true, Contains: "movie_volume"},
		{Name: "sleep_mode dry-run success", Command: "sleep-mode", Payload: `{}`, ShouldPass: true, Contains: "display_sleep"},
	}

	report := HarnessReport{Name: "home", Passed: true}
	for _, tc := range cases {
		result := HarnessCaseResult{Name: tc.Name}
		plan, err := Build(tc.Command, []byte(tc.Payload))
		if tc.ShouldPass {
			if err != nil {
				result.Message = err.Error()
			} else if tc.Contains != "" && !strings.Contains(planText(plan), tc.Contains) {
				result.Message = "expected output to contain " + tc.Contains
			} else {
				result.Passed = true
				result.Message = "ok"
			}
		} else {
			if err != nil {
				result.Passed = true
				result.Message = "failed safely: " + err.Error()
			} else {
				result.Message = "expected safe failure but command passed"
			}
		}
		if !result.Passed {
			report.Passed = false
		}
		report.Results = append(report.Results, result)
	}
	return report
}

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
	if len(commerce.RecurringCandidates) == 0 {
		report.addCheck("recurring purchase candidate detail", false, "missing recurring candidate")
	} else {
		candidate := commerce.RecurringCandidates[0]
		passed := candidate.MerchantName == "Coupang" &&
			candidate.ItemName == "Bottled water 2L x 6" &&
			candidate.PurchaseCount == 2 &&
			candidate.LatestTotalMinorUnits == 11_800
		report.addCheck("recurring purchase candidate detail", passed, fmt.Sprintf("%s %s count=%d total=%d", candidate.MerchantName, candidate.ItemName, candidate.PurchaseCount, candidate.LatestTotalMinorUnits))
	}

	user := commerceOwner(commerce.OwnerBreakdown, "user")
	report.addCheck("user commerce scope", user != nil && user.Records == 1 && user.PurchaseSpendMinorUnits == 3_200, commerceOwnerMessage(user))
	household := commerceOwner(commerce.OwnerBreakdown, "household")
	report.addCheck("household commerce scope", household != nil && household.Records == 2 && household.PurchaseSpendMinorUnits == 23_600, commerceOwnerMessage(household))
	return report
}

func planText(plan Plan) string {
	var b strings.Builder
	b.WriteString(plan.Name)
	for _, invocation := range plan.Invocations {
		b.WriteString(" ")
		b.WriteString(invocation.Label)
		b.WriteString(" ")
		b.WriteString(invocation.URL)
		for _, arg := range invocation.Argv {
			b.WriteString(" ")
			b.WriteString(arg)
		}
	}
	return b.String()
}

func (report *HarnessReport) addCheck(name string, passed bool, message string) {
	if strings.TrimSpace(message) == "" {
		message = "ok"
	}
	if passed {
		message = "ok"
	}
	if !passed {
		report.Passed = false
	}
	report.Results = append(report.Results, HarnessCaseResult{
		Name:    name,
		Passed:  passed,
		Message: message,
	})
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
