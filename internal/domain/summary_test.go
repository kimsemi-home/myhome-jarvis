package domain

import "testing"

func TestBuildSummaryFromRepoFixtures(t *testing.T) {
	root := repoRoot(t)
	summary, err := BuildSummary(root)
	if err != nil {
		t.Fatal(err)
	}

	assertFinanceSummary(t, summary)
	assertCommerceSummary(t, summary)
	assertStoragePolicy(t, summary)
	assertRecommendationSummary(t, summary)
	assertHouseholdSummary(t, summary)
}
