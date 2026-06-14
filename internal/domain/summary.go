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
	Finance  FinanceSummary  `json:"finance"`
	Commerce CommerceSummary `json:"commerce"`
	Storage  StoragePolicy   `json:"storage"`
}

type FinanceSummary struct {
	Records          int      `json:"records"`
	Currency         string   `json:"currency"`
	CreditMinorUnits int64    `json:"credit_minor_units"`
	DebitMinorUnits  int64    `json:"debit_minor_units"`
	NetMinorUnits    int64    `json:"net_minor_units"`
	Categories       []string `json:"categories"`
}

type CommerceSummary struct {
	Records                 int                          `json:"records"`
	RecurringCandidateCount int                          `json:"recurring_candidate_count"`
	RecurringCandidates     []RecurringPurchaseCandidate `json:"recurring_candidates"`
	Categories              []string                     `json:"categories"`
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

type moneyAmount struct {
	MinorUnits int64  `json:"minor_units"`
	Currency   string `json:"currency"`
}

type financeTransaction struct {
	Amount    moneyAmount `json:"amount"`
	Direction string      `json:"direction"`
	Category  string      `json:"category"`
}

type commercePurchase struct {
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
	return Summary{Finance: finance, Commerce: commerce, Storage: storage}, nil
}

func BuildFinanceSummary(path string) (FinanceSummary, error) {
	var summary FinanceSummary
	currencies := map[string]bool{}
	categories := map[string]bool{}
	err := readJSONL(path, func(line int, transaction financeTransaction) error {
		summary.Records++
		if transaction.Amount.Currency != "" {
			currencies[transaction.Amount.Currency] = true
		}
		if strings.TrimSpace(transaction.Category) != "" {
			categories[transaction.Category] = true
		}
		switch transaction.Direction {
		case "credit":
			summary.CreditMinorUnits += transaction.Amount.MinorUnits
		case "debit":
			summary.DebitMinorUnits += transaction.Amount.MinorUnits
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
	return summary, nil
}

func BuildCommerceSummary(path string) (CommerceSummary, error) {
	var summary CommerceSummary
	categories := map[string]bool{}
	groups := map[string]*RecurringPurchaseCandidate{}
	err := readJSONL(path, func(line int, purchase commercePurchase) error {
		summary.Records++
		if strings.TrimSpace(purchase.Category) != "" {
			categories[purchase.Category] = true
		}
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
	summary.Categories = sortedKeys(categories)
	return summary, nil
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
