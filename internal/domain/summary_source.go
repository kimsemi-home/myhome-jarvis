package domain

import (
	"path/filepath"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func BuildSummaryWithFinance(
	root string,
	transactions []financeconnector.SourceTransaction,
	parity financeconnector.ParityReport,
) (Summary, error) {
	finance, err := BuildFinanceSummaryFromSource(transactions)
	if err != nil {
		return Summary{}, err
	}
	finance.ConnectorProvider = parity.Provider
	finance.ConnectorFixtureOnly = strings.HasPrefix(parity.FixturePath, "fixtures/")
	finance.ConnectorParityPercent = parity.ParityPercent
	finance.ConnectorParityPass = parity.Pass
	commerce, err := BuildCommerceSummary(filepath.Join(root, "fixtures", "commerce_purchases.jsonl"))
	if err != nil {
		return Summary{}, err
	}
	storage, err := ReadStoragePolicy(filepath.Join(root, "generated", "storage.generated.json"))
	if err != nil {
		return Summary{}, err
	}
	return Summary{
		Finance: finance, Commerce: commerce, Storage: storage,
		Recommendations: BuildRecommendationsSummary(finance, commerce),
		Household:       BuildHouseholdSummary(finance, commerce),
	}, nil
}
