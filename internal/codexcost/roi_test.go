package codexcost

import "testing"

func TestROISummaryIncludesKnownScopesAndArchiveEvidence(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"assistant_loop","unit_kind":"codex_tokens","amount":100,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")

	summary, err := ROISummaryForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if summary.ScopeCount != 4 || summary.TrackedScopeCount != 1 {
		t.Fatalf("scope counts = %#v", summary)
	}
	if !summary.StorageArchiveReady || !summary.NoiseBudgetReady {
		t.Fatalf("archive evidence = %#v", summary)
	}
	rows := roiRowsByScope(summary.Rows)
	if rows["assistant_loop"].Status != "tracked" {
		t.Fatalf("assistant row = %#v", rows["assistant_loop"])
	}
	if rows["linear_project"].Recommendation != "no_usage_yet" {
		t.Fatalf("linear row = %#v", rows["linear_project"])
	}
}

func TestROISummaryRequiresReviewWhenSustainabilityIsBlocked(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":10,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")

	summary, err := ROISummaryForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	row := roiRowsByScope(summary.Rows)["repo"]
	if row.Recommendation != "review_before_scaling" {
		t.Fatalf("recommendation = %#v", row)
	}
}
