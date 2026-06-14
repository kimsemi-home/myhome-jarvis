package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Summary struct {
	Finance         FinanceSummary         `json:"finance"`
	Commerce        CommerceSummary        `json:"commerce"`
	Storage         StoragePolicy          `json:"storage"`
	Recommendations RecommendationsSummary `json:"recommendations"`
	Household       HouseholdSummary       `json:"household"`
}

type FinanceSummary struct {
	Records                int                   `json:"records"`
	Currency               string                `json:"currency"`
	CreditMinorUnits       int64                 `json:"credit_minor_units"`
	DebitMinorUnits        int64                 `json:"debit_minor_units"`
	NetMinorUnits          int64                 `json:"net_minor_units"`
	SubscriptionMinorUnits int64                 `json:"subscription_minor_units"`
	SubscriptionCount      int                   `json:"subscription_count"`
	CardDebitMinorUnits    int64                 `json:"card_debit_minor_units"`
	CardDebitCount         int                   `json:"card_debit_count"`
	Categories             []string              `json:"categories"`
	OwnerBreakdown         []FinanceOwnerSummary `json:"owner_breakdown"`
}

type FinanceOwnerSummary struct {
	Owner            string `json:"owner"`
	Records          int    `json:"records"`
	Currency         string `json:"currency"`
	CreditMinorUnits int64  `json:"credit_minor_units"`
	DebitMinorUnits  int64  `json:"debit_minor_units"`
	NetMinorUnits    int64  `json:"net_minor_units"`
}

type CommerceSummary struct {
	Records                 int                          `json:"records"`
	Currency                string                       `json:"currency"`
	TotalSpendMinorUnits    int64                        `json:"total_spend_minor_units"`
	RecurringCandidateCount int                          `json:"recurring_candidate_count"`
	RecurringCandidates     []RecurringPurchaseCandidate `json:"recurring_candidates"`
	Categories              []string                     `json:"categories"`
	OwnerBreakdown          []CommerceOwnerSummary       `json:"owner_breakdown"`
}

type CommerceOwnerSummary struct {
	Owner                   string `json:"owner"`
	Records                 int    `json:"records"`
	Currency                string `json:"currency"`
	PurchaseSpendMinorUnits int64  `json:"purchase_spend_minor_units"`
}

type RecurringPurchaseCandidate struct {
	MerchantName          string `json:"merchant_name"`
	ItemName              string `json:"item_name"`
	Currency              string `json:"currency"`
	PurchaseCount         int    `json:"purchase_count"`
	LatestTotalMinorUnits int64  `json:"latest_total_minor_units"`
	LatestPurchasedAt     string `json:"latest_purchased_at"`
}

type StoragePolicy struct {
	FixtureFormat  string   `json:"fixture_format"`
	LakeLayers     []string `json:"lake_layers"`
	Datasets       []string `json:"datasets"`
	LongTermFormat string   `json:"long_term_format"`
	Compression    string   `json:"compression"`
	PrivateRoot    string   `json:"private_root"`
}

type RecommendationsSummary struct {
	Count int                  `json:"count"`
	Items []RecommendationItem `json:"items"`
}

type RecommendationItem struct {
	Kind                       string `json:"kind"`
	Title                      string `json:"title"`
	Rationale                  string `json:"rationale"`
	Score                      int    `json:"score"`
	Currency                   string `json:"currency"`
	EstimatedMonthlyMinorUnits int64  `json:"estimated_monthly_minor_units"`
	EvidenceCount              int    `json:"evidence_count"`
}

type HouseholdSummary struct {
	Scopes []HouseholdScopeSummary `json:"scopes"`
}

type HouseholdScopeSummary struct {
	Scope                   string `json:"scope"`
	Label                   string `json:"label"`
	Currency                string `json:"currency"`
	FinanceRecords          int    `json:"finance_records"`
	FinanceCreditMinorUnits int64  `json:"finance_credit_minor_units"`
	FinanceDebitMinorUnits  int64  `json:"finance_debit_minor_units"`
	FinanceNetMinorUnits    int64  `json:"finance_net_minor_units"`
	PurchaseRecords         int    `json:"purchase_records"`
	PurchaseSpendMinorUnits int64  `json:"purchase_spend_minor_units"`
}

type moneyAmount struct {
	MinorUnits int64  `json:"minor_units"`
	Currency   string `json:"currency"`
}

type financeTransaction struct {
	Owner     string      `json:"owner"`
	Amount    moneyAmount `json:"amount"`
	Direction string      `json:"direction"`
	Category  string      `json:"category"`
	CardID    string      `json:"card_id"`
}

type commercePurchase struct {
	Owner              string      `json:"owner"`
	PurchasedAt        string      `json:"purchased_at"`
	MerchantName       string      `json:"merchant_name"`
	ItemName           string      `json:"item_name"`
	TotalPrice         moneyAmount `json:"total_price"`
	Category           string      `json:"category"`
	RecurringCandidate bool        `json:"recurring_candidate"`
}

func BuildSummary(root string) (Summary, error) {
	finance, err := BuildFinanceSummary(filepath.Join(root, "fixtures", "finance_transactions.jsonl"))
	if err != nil {
		return Summary{}, err
	}
	commerce, err := BuildCommerceSummary(filepath.Join(root, "fixtures", "commerce_purchases.jsonl"))
	if err != nil {
		return Summary{}, err
	}
	storage, err := ReadStoragePolicy(filepath.Join(root, "generated", "storage.generated.json"))
	if err != nil {
		return Summary{}, err
	}
	recommendations := BuildRecommendationsSummary(finance, commerce)
	household := BuildHouseholdSummary(finance, commerce)
	return Summary{
		Finance:         finance,
		Commerce:        commerce,
		Storage:         storage,
		Recommendations: recommendations,
		Household:       household,
	}, nil
}

func BuildFinanceSummary(path string) (FinanceSummary, error) {
	var summary FinanceSummary
	currencies := map[string]bool{}
	categories := map[string]bool{}
	owners := map[string]*FinanceOwnerSummary{}
	ownerCurrencies := map[string]map[string]bool{}
	err := readJSONL(path, func(line int, transaction financeTransaction) error {
		owner := normalizeOwner(transaction.Owner)
		summary.Records++
		if transaction.Amount.Currency != "" {
			currencies[transaction.Amount.Currency] = true
		}
		if strings.TrimSpace(transaction.Category) != "" {
			categories[transaction.Category] = true
		}
		ownerSummary := financeOwnerSummary(owners, owner)
		ownerSummary.Records++
		recordOwnerCurrency(ownerCurrencies, owner, transaction.Amount.Currency)
		switch transaction.Direction {
		case "credit":
			summary.CreditMinorUnits += transaction.Amount.MinorUnits
			ownerSummary.CreditMinorUnits += transaction.Amount.MinorUnits
		case "debit":
			summary.DebitMinorUnits += transaction.Amount.MinorUnits
			ownerSummary.DebitMinorUnits += transaction.Amount.MinorUnits
			if transaction.Category == "subscription" {
				summary.SubscriptionMinorUnits += transaction.Amount.MinorUnits
				summary.SubscriptionCount++
			}
			if strings.TrimSpace(transaction.CardID) != "" {
				summary.CardDebitMinorUnits += transaction.Amount.MinorUnits
				summary.CardDebitCount++
			}
		default:
			return fmt.Errorf("line %d: unknown direction %q", line, transaction.Direction)
		}
		return nil
	})
	if err != nil {
		return FinanceSummary{}, err
	}
	summary.NetMinorUnits = summary.CreditMinorUnits - summary.DebitMinorUnits
	summary.Currency = summaryCurrency(currencies)
	summary.Categories = sortedKeys(categories)
	summary.OwnerBreakdown = financeOwnerBreakdown(owners, ownerCurrencies)
	return summary, nil
}

func BuildRecommendationsSummary(finance FinanceSummary, commerce CommerceSummary) RecommendationsSummary {
	items := []RecommendationItem{}
	if finance.NetMinorUnits > 0 {
		items = append(items, RecommendationItem{
			Kind:                       "cash_buffer",
			Title:                      "Keep household cash buffer",
			Rationale:                  "Fixture cashflow is positive; reserve surplus before recommendations become executable.",
			Score:                      clampScore(45 + int(finance.NetMinorUnits/1_000_000)),
			Currency:                   finance.Currency,
			EstimatedMonthlyMinorUnits: finance.NetMinorUnits,
			EvidenceCount:              finance.Records,
		})
	}
	if finance.SubscriptionMinorUnits > 0 {
		items = append(items, RecommendationItem{
			Kind:                       "subscription_review",
			Title:                      "Review household subscriptions",
			Rationale:                  "Subscription-like debit fixtures exist; keep this as a review-only recommendation.",
			Score:                      clampScore(55 + int(finance.SubscriptionMinorUnits/10_000)),
			Currency:                   finance.Currency,
			EstimatedMonthlyMinorUnits: finance.SubscriptionMinorUnits,
			EvidenceCount:              finance.SubscriptionCount,
		})
	}
	if finance.CardDebitMinorUnits > 0 {
		items = append(items, RecommendationItem{
			Kind:                       "card_usage_review",
			Title:                      "Review card-linked household spend",
			Rationale:                  "Card-linked debit fixtures exist; keep this as a review-only recommendation, not a card action.",
			Score:                      clampScore(52 + int(finance.CardDebitMinorUnits/10_000)),
			Currency:                   finance.Currency,
			EstimatedMonthlyMinorUnits: finance.CardDebitMinorUnits,
			EvidenceCount:              finance.CardDebitCount,
		})
	}
	for _, candidate := range commerce.RecurringCandidates {
		items = append(items, RecommendationItem{
			Kind:                       "recurring_purchase_review",
			Title:                      "Compare recurring purchase: " + candidate.ItemName,
			Rationale:                  candidate.MerchantName + " appears repeatedly in local purchase fixtures.",
			Score:                      clampScore(50 + candidate.PurchaseCount*10 + int(candidate.LatestTotalMinorUnits/1_000)),
			Currency:                   candidate.Currency,
			EstimatedMonthlyMinorUnits: candidate.LatestTotalMinorUnits,
			EvidenceCount:              candidate.PurchaseCount,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Score == items[j].Score {
			return items[i].Title < items[j].Title
		}
		return items[i].Score > items[j].Score
	})
	return RecommendationsSummary{Count: len(items), Items: items}
}

func BuildCommerceSummary(path string) (CommerceSummary, error) {
	var summary CommerceSummary
	currencies := map[string]bool{}
	categories := map[string]bool{}
	groups := map[string]*RecurringPurchaseCandidate{}
	owners := map[string]*CommerceOwnerSummary{}
	ownerCurrencies := map[string]map[string]bool{}
	err := readJSONL(path, func(line int, purchase commercePurchase) error {
		owner := normalizeOwner(purchase.Owner)
		summary.Records++
		if purchase.TotalPrice.Currency != "" {
			currencies[purchase.TotalPrice.Currency] = true
		}
		summary.TotalSpendMinorUnits += purchase.TotalPrice.MinorUnits
		if strings.TrimSpace(purchase.Category) != "" {
			categories[purchase.Category] = true
		}
		ownerSummary := commerceOwnerSummary(owners, owner)
		ownerSummary.Records++
		ownerSummary.PurchaseSpendMinorUnits += purchase.TotalPrice.MinorUnits
		recordOwnerCurrency(ownerCurrencies, owner, purchase.TotalPrice.Currency)
		if !purchase.RecurringCandidate {
			return nil
		}
		key := strings.Join([]string{
			purchase.MerchantName,
			purchase.ItemName,
			purchase.TotalPrice.Currency,
		}, "\x00")
		candidate, ok := groups[key]
		if !ok {
			candidate = &RecurringPurchaseCandidate{
				MerchantName:          purchase.MerchantName,
				ItemName:              purchase.ItemName,
				Currency:              purchase.TotalPrice.Currency,
				LatestTotalMinorUnits: purchase.TotalPrice.MinorUnits,
				LatestPurchasedAt:     purchase.PurchasedAt,
			}
			groups[key] = candidate
		}
		candidate.PurchaseCount++
		if purchase.PurchasedAt > candidate.LatestPurchasedAt {
			candidate.LatestPurchasedAt = purchase.PurchasedAt
			candidate.LatestTotalMinorUnits = purchase.TotalPrice.MinorUnits
		}
		return nil
	})
	if err != nil {
		return CommerceSummary{}, err
	}
	for _, candidate := range groups {
		if candidate.PurchaseCount < 2 {
			continue
		}
		summary.RecurringCandidates = append(summary.RecurringCandidates, *candidate)
	}
	sort.Slice(summary.RecurringCandidates, func(i, j int) bool {
		left := summary.RecurringCandidates[i]
		right := summary.RecurringCandidates[j]
		if left.MerchantName == right.MerchantName {
			return left.ItemName < right.ItemName
		}
		return left.MerchantName < right.MerchantName
	})
	summary.RecurringCandidateCount = len(summary.RecurringCandidates)
	summary.Currency = summaryCurrency(currencies)
	summary.Categories = sortedKeys(categories)
	summary.OwnerBreakdown = commerceOwnerBreakdown(owners, ownerCurrencies)
	return summary, nil
}

func BuildHouseholdSummary(finance FinanceSummary, commerce CommerceSummary) HouseholdSummary {
	financeOwners := mapFinanceOwners(finance.OwnerBreakdown)
	commerceOwners := mapCommerceOwners(commerce.OwnerBreakdown)
	scopes := []HouseholdScopeSummary{
		ownerScopeSummary("user", "User", financeOwners["user"], commerceOwners["user"]),
		ownerScopeSummary("spouse", "Spouse", financeOwners["spouse"], commerceOwners["spouse"]),
		{
			Scope:                   "household",
			Label:                   "Household",
			Currency:                firstCurrency(finance.Currency, commerce.Currency, "KRW"),
			FinanceRecords:          finance.Records,
			FinanceCreditMinorUnits: finance.CreditMinorUnits,
			FinanceDebitMinorUnits:  finance.DebitMinorUnits,
			FinanceNetMinorUnits:    finance.NetMinorUnits,
			PurchaseRecords:         commerce.Records,
			PurchaseSpendMinorUnits: commerce.TotalSpendMinorUnits,
		},
	}
	return HouseholdSummary{Scopes: scopes}
}

func normalizeOwner(owner string) string {
	switch strings.TrimSpace(strings.ToLower(owner)) {
	case "user":
		return "user"
	case "spouse":
		return "spouse"
	case "household":
		return "household"
	default:
		return "unknown"
	}
}

func financeOwnerSummary(owners map[string]*FinanceOwnerSummary, owner string) *FinanceOwnerSummary {
	summary, ok := owners[owner]
	if !ok {
		summary = &FinanceOwnerSummary{Owner: owner}
		owners[owner] = summary
	}
	return summary
}

func commerceOwnerSummary(owners map[string]*CommerceOwnerSummary, owner string) *CommerceOwnerSummary {
	summary, ok := owners[owner]
	if !ok {
		summary = &CommerceOwnerSummary{Owner: owner}
		owners[owner] = summary
	}
	return summary
}

func recordOwnerCurrency(currencies map[string]map[string]bool, owner string, currency string) {
	if strings.TrimSpace(currency) == "" {
		return
	}
	if _, ok := currencies[owner]; !ok {
		currencies[owner] = map[string]bool{}
	}
	currencies[owner][currency] = true
}

func financeOwnerBreakdown(
	owners map[string]*FinanceOwnerSummary,
	currencies map[string]map[string]bool,
) []FinanceOwnerSummary {
	breakdown := make([]FinanceOwnerSummary, 0, len(owners))
	for _, owner := range ownerBreakdownOrder(owners) {
		summary := *owners[owner]
		summary.NetMinorUnits = summary.CreditMinorUnits - summary.DebitMinorUnits
		summary.Currency = summaryCurrency(currencies[owner])
		breakdown = append(breakdown, summary)
	}
	return breakdown
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

func ownerBreakdownOrder[T any](owners map[string]*T) []string {
	keys := make([]string, 0, len(owners))
	for key := range owners {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return ownerRank(keys[i]) < ownerRank(keys[j])
	})
	return keys
}

func ownerRank(owner string) int {
	switch owner {
	case "user":
		return 0
	case "spouse":
		return 1
	case "household":
		return 2
	default:
		return 3
	}
}

func mapFinanceOwners(items []FinanceOwnerSummary) map[string]*FinanceOwnerSummary {
	mapped := map[string]*FinanceOwnerSummary{}
	for index := range items {
		mapped[items[index].Owner] = &items[index]
	}
	return mapped
}

func mapCommerceOwners(items []CommerceOwnerSummary) map[string]*CommerceOwnerSummary {
	mapped := map[string]*CommerceOwnerSummary{}
	for index := range items {
		mapped[items[index].Owner] = &items[index]
	}
	return mapped
}

func ownerScopeSummary(
	scope string,
	label string,
	finance *FinanceOwnerSummary,
	commerce *CommerceOwnerSummary,
) HouseholdScopeSummary {
	var summary HouseholdScopeSummary
	summary.Scope = scope
	summary.Label = label
	summary.Currency = "KRW"
	if finance != nil {
		summary.Currency = firstCurrency(finance.Currency, summary.Currency)
		summary.FinanceRecords = finance.Records
		summary.FinanceCreditMinorUnits = finance.CreditMinorUnits
		summary.FinanceDebitMinorUnits = finance.DebitMinorUnits
		summary.FinanceNetMinorUnits = finance.NetMinorUnits
	}
	if commerce != nil {
		summary.Currency = firstCurrency(summary.Currency, commerce.Currency)
		summary.PurchaseRecords = commerce.Records
		summary.PurchaseSpendMinorUnits = commerce.PurchaseSpendMinorUnits
	}
	return summary
}

func firstCurrency(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func ReadStoragePolicy(path string) (StoragePolicy, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return StoragePolicy{}, err
	}
	var policy StoragePolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		return StoragePolicy{}, err
	}
	return policy, nil
}

func readJSONL[T any](path string, handle func(line int, value T) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		var value T
		if err := json.Unmarshal([]byte(text), &value); err != nil {
			return fmt.Errorf("line %d: %w", line, err)
		}
		if err := handle(line, value); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func summaryCurrency(currencies map[string]bool) string {
	switch len(currencies) {
	case 0:
		return ""
	case 1:
		for currency := range currencies {
			return currency
		}
	}
	return "mixed"
}

func sortedKeys(values map[string]bool) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func clampScore(value int) int {
	switch {
	case value < 0:
		return 0
	case value > 100:
		return 100
	default:
		return value
	}
}
