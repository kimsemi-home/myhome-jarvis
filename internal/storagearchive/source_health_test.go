package storagearchive

import (
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func TestStatusReportsSourceHealth(t *testing.T) {
	root := t.TempDir()
	quality := privateQualitySource()
	audit := domain.PrivateLogSource{
		Key: "audit", Path: "data/private/audit/command-intents.jsonl", Format: "jsonl",
	}
	policy := testStoragePolicy(quality)
	policy.PrivateLogSources = []domain.PrivateLogSource{quality, audit}
	writeStoragePolicy(t, root, policy)
	writePrivateFile(t, root, quality.Path,
		`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	writePrivateFile(t, root, audit.Path,
		`{"source":"audit","kind":"intent","evidence_ref":"b"}`+"\n"+
			`{"source":"audit","kind":"intent","evidence_ref":"b"}`+"\n")
	if _, err := RunForRoot(root); err != nil {
		t.Fatal(err)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	healthy := sourceHealthByKey(status.SourceHealth, "quality")
	if !healthy.ArchiveEvidencePresent ||
		!healthy.HashCacheKeyPresent ||
		!healthy.BudgetOK ||
		healthy.HealthDebt ||
		healthy.LatestArchivedAt == "" {
		t.Fatalf("healthy source = %#v", healthy)
	}
	breached := sourceHealthByKey(status.SourceHealth, "audit")
	if breached.ArchiveEvidencePresent ||
		breached.BudgetOK ||
		!breached.HealthDebt ||
		breached.LatestBudgetVerdict != "breach" ||
		breached.NoiseRatioPercent == 0 {
		t.Fatalf("breached source = %#v", breached)
	}
	if status.SourceHealthDebtCount != 1 {
		t.Fatalf("source health debt count = %#v", status)
	}
}

func sourceHealthByKey(sources []SourceHealth, key string) SourceHealth {
	for _, source := range sources {
		if source.SourceKey == key {
			return source
		}
	}
	return SourceHealth{}
}
