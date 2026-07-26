package financeconnector

import (
	"path/filepath"
	"testing"
)

func TestTossMyDataFixtureMeetsParityTarget(t *testing.T) {
	root := filepath.Join("..", "..")
	report, err := VerifyFixture(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.Pass {
		t.Fatalf("parity report = %#v", report)
	}
	if report.ParityPercent < ParityThreshold {
		t.Fatalf("parity = %.2f, threshold = %.2f", report.ParityPercent, ParityThreshold)
	}
	if report.RecordsWithRequiredFields != report.Records {
		t.Fatalf("required fields = %d/%d", report.RecordsWithRequiredFields, report.Records)
	}
}

func TestParityRejectsMissingRequiredField(t *testing.T) {
	report := AssessParity([]SourceTransaction{{
		TransactionID: "txn",
		Source:        ProviderTossMyData,
		Owner:         "user",
		OccurredAt:    "2026-06-01T00:00:00+09:00",
		Currency:      "KRW",
		Direction:     "debit",
		RawRef:        "toss:txn",
	}})
	if report.Pass {
		t.Fatalf("missing amount should fail: %#v", report)
	}
}
