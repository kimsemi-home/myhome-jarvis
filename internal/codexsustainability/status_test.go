package codexsustainability

import (
	"testing"
	"time"
)

func TestMissingLedgerBlocksSustainabilityStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := statusForRootAt(root, mustTime(t, "2026-06-20T00:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.SustainabilityPosture != "blocked" {
		t.Fatalf("status = %#v", status)
	}
	if status.TrendPosture != "missing" || status.EvidenceFreshness != "missing" {
		t.Fatalf("posture = %#v", status)
	}
}

func TestStatusMeasuresTrendAndSustainabilityDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/codex-sustainability/evidence.jsonl", debtLedgerFixture())

	status, err := statusForRootAt(root, mustTime(t, "2026-06-20T00:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if status.TrendPosture != "slower_than_trend" || status.SustainabilityPosture != "review_required" {
		t.Fatalf("posture = %#v", status)
	}
	if status.CostPerAcceptedChange != 700000 || status.ReviewGateCount < 3 {
		t.Fatalf("review gates = %#v", status)
	}
	if status.OptimizationClaimWithoutEvidenceCount != 1 || status.MissingEvidenceCount != 1 {
		t.Fatalf("evidence counts = %#v", status)
	}
}

func mustTime(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatal(err)
	}
	return parsed
}
