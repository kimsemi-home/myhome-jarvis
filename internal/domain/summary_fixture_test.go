package domain

import "testing"

func assertFinanceSummary(t *testing.T, summary Summary) {
	t.Helper()
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
	if summary.Finance.CardDebitMinorUnits != 153_200 || summary.Finance.CardDebitCount != 2 {
		t.Fatalf("card debit summary = %#v", summary.Finance)
	}
	if len(summary.Finance.OwnerBreakdown) != 2 {
		t.Fatalf("finance owner breakdown = %#v", summary.Finance.OwnerBreakdown)
	}
}

func assertCommerceSummary(t *testing.T, summary Summary) {
	t.Helper()
	if summary.Commerce.Records != 3 {
		t.Fatalf("commerce records = %d", summary.Commerce.Records)
	}
	if summary.Commerce.TotalSpendMinorUnits != 26_800 {
		t.Fatalf("commerce spend = %d", summary.Commerce.TotalSpendMinorUnits)
	}
	if summary.Commerce.RecurringCandidateCount != 1 {
		t.Fatalf("recurring candidates = %d", summary.Commerce.RecurringCandidateCount)
	}
	if summary.Commerce.RecurringCandidates[0].MerchantName != "Coupang" {
		t.Fatalf("recurring merchant = %q", summary.Commerce.RecurringCandidates[0].MerchantName)
	}
}

func assertStoragePolicy(t *testing.T, summary Summary) {
	t.Helper()
	if summary.Storage.LongTermFormat != "parquet" || summary.Storage.Compression != "zstd" {
		t.Fatalf("storage policy = %#v", summary.Storage)
	}
	if summary.Storage.LogArchive.Compression != "gzip" ||
		!summary.Storage.EvidenceNoiseBudget.BreachBlocksArchive {
		t.Fatalf("storage archive policy = %#v", summary.Storage)
	}
}
