package codexcost

import (
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func TestBriefAllowsHealthyLocalFirstLoop(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"assistant_loop","unit_kind":"codex_tokens","amount":100,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")
	writeFile(t, root, "data/private/codex-cost/attribution.jsonl",
		`{"at":"2026-06-19T00:00:01Z","scope":"repo","subject_key":"repo:test","cost_ref":"cost_ref_test","unit_kind":"codex_tokens","amount":100,"basis":"merged_pr","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")
	if _, err := storagearchive.RunForRoot(root); err != nil {
		t.Fatal(err)
	}

	brief, err := BriefForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if brief.Decision != "allow" || brief.NextSafeAction != "continue_local_first_loop" {
		t.Fatalf("brief = %#v", brief)
	}
	if !brief.PublicSafe ||
		brief.StorageArchivePattern != "compress_then_archive" ||
		!brief.NoiseBudgetReady ||
		brief.MaxNoiseRatioPercent != 20 ||
		brief.ArchiveManifestEntryCount == 0 ||
		brief.WarningUnitThreshold != 100000 ||
		brief.ReviewUnitThreshold != 500000 {
		t.Fatalf("evidence summary = %#v", brief)
	}
}
