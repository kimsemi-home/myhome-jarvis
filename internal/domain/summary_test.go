package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildSummaryFromRepoFixtures(t *testing.T) {
	root := repoRoot(t)
	summary, err := BuildSummary(root)
	if err != nil {
		t.Fatal(err)
	}

	if summary.Finance.Records != 3 {
		t.Fatalf("finance records = %d", summary.Finance.Records)
	}
	if summary.Finance.Currency != "KRW" {
		t.Fatalf("finance currency = %q", summary.Finance.Currency)
	}
	if summary.Finance.NetMinorUnits != 4_346_800 {
		t.Fatalf("finance net = %d", summary.Finance.NetMinorUnits)
	}
	if summary.Finance.SubscriptionMinorUnits != 65_900 {
		t.Fatalf("subscription total = %d", summary.Finance.SubscriptionMinorUnits)
	}
	if summary.Commerce.Records != 3 {
		t.Fatalf("commerce records = %d", summary.Commerce.Records)
	}
	if summary.Commerce.RecurringCandidateCount != 1 {
		t.Fatalf("recurring candidates = %d", summary.Commerce.RecurringCandidateCount)
	}
	if summary.Commerce.RecurringCandidates[0].MerchantName != "Coupang" {
		t.Fatalf("recurring merchant = %q", summary.Commerce.RecurringCandidates[0].MerchantName)
	}
	if summary.Storage.LongTermFormat != "parquet" || summary.Storage.Compression != "zstd" {
		t.Fatalf("storage policy = %#v", summary.Storage)
	}
	if summary.Recommendations.Count != 3 {
		t.Fatalf("recommendation count = %d", summary.Recommendations.Count)
	}
	if summary.Recommendations.Items[0].Kind != "recurring_purchase_review" {
		t.Fatalf("top recommendation = %#v", summary.Recommendations.Items[0])
	}
	if summary.Recommendations.Items[0].Score < summary.Recommendations.Items[1].Score {
		t.Fatalf("recommendations are not ranked: %#v", summary.Recommendations.Items)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
