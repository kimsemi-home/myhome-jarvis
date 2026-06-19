package domain

import "path/filepath"

type Summary struct {
	Finance         FinanceSummary         `json:"finance"`
	Commerce        CommerceSummary        `json:"commerce"`
	Storage         StoragePolicy          `json:"storage"`
	Recommendations RecommendationsSummary `json:"recommendations"`
	Household       HouseholdSummary       `json:"household"`
}

func BuildSummary(root string) (Summary, error) {
	finance, err := BuildFinanceSummary(
		filepath.Join(root, "fixtures", "finance_transactions.jsonl"),
	)
	if err != nil {
		return Summary{}, err
	}
	commerce, err := BuildCommerceSummary(
		filepath.Join(root, "fixtures", "commerce_purchases.jsonl"),
	)
	if err != nil {
		return Summary{}, err
	}
	storage, err := ReadStoragePolicy(
		filepath.Join(root, "generated", "storage.generated.json"),
	)
	if err != nil {
		return Summary{}, err
	}
	return Summary{
		Finance:         finance,
		Commerce:        commerce,
		Storage:         storage,
		Recommendations: BuildRecommendationsSummary(finance, commerce),
		Household:       BuildHouseholdSummary(finance, commerce),
	}, nil
}
