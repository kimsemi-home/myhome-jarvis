package codexcost

import "testing"

func TestROISummaryUsesAttributionWithoutDoubleCountingTotal(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"assistant_loop","unit_kind":"codex_tokens","amount":100,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")
	writeFile(t, root, "data/private/codex-cost/attribution.jsonl",
		`{"at":"2026-06-19T00:00:01Z","scope":"repo","subject_key":"repo:kimsemi-home/myhome-jarvis","subject_hash":"subject_test","unit_kind":"codex_tokens","amount":40,"basis":"merged_pr","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n"+
			`{"at":"2026-06-19T00:00:02Z","scope":"linear_project","subject_key":"linear:KIM-133","subject_hash":"subject_linear","unit_kind":"codex_tokens","amount":60,"basis":"merged_pr","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")

	summary, err := ROISummaryForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	rows := roiRowsByScope(summary.Rows)
	if summary.TotalUnits != 100 || summary.AttributedUnits != 100 {
		t.Fatalf("summary = %#v", summary)
	}
	if summary.AttributionCoveragePercent != 100 {
		t.Fatalf("coverage = %#v", summary)
	}
	if rows["repo"].Status != "attributed" ||
		rows["repo"].CostUnits != 40 ||
		rows["repo"].AttributionSubjectCount != 1 {
		t.Fatalf("repo row = %#v", rows["repo"])
	}
}
