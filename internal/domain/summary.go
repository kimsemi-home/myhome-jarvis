package domain

import "github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"

type Summary struct {
	Finance         FinanceSummary         `json:"finance"`
	Commerce        CommerceSummary        `json:"commerce"`
	Storage         StoragePolicy          `json:"storage"`
	Recommendations RecommendationsSummary `json:"recommendations"`
	Household       HouseholdSummary       `json:"household"`
}

func BuildSummary(root string) (Summary, error) {
	transactions, err := financeconnector.LoadFixture(root)
	if err != nil {
		return Summary{}, err
	}
	parity := financeconnector.AssessParity(transactions)
	parity.FixturePath = financeconnector.FixtureRelativePath
	return BuildSummaryWithFinance(root, transactions, parity)
}
