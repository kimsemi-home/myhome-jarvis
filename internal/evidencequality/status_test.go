package evidencequality

import "testing"

func TestMissingLedgerReturnsZeroRedactedStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.SnapshotCount != 0 || status.ReassessmentDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.LedgerPath != "data/private/evidence-quality/snapshots.jsonl" || status.PolicyPath != PolicyRelativePath {
		t.Fatalf("paths = %#v", status)
	}
}

func TestStatusCountsStaleLowBlockedAndMappingDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/evidence-quality/snapshots.jsonl",
		`{"id":"eq_1","at":"2026-04-01T00:00:00Z","evidence_ref":"generated/evidence.generated.json","purpose":"confidence_assessment","quality_level":"high","schema_version":"evidence:v1","ontology_version":"concepts:v1","mapping_confidence":"high","assessed_by":"deterministic_verifier","reassessment_reasons":["age"],"raw_notes":"private"}`+"\n"+
			`{"id":"eq_2","at":"2026-06-18T00:00:00Z","evidence_ref":"docs/evidence-graph.md","purpose":"root_cause","quality_level":"low","schema_version":"evidence:v1","ontology_version":"concepts:v1","mapping_confidence":"low","assessed_by":"governance_steward","reassessment_reasons":["ontology_version_change"],"raw_evidence":"private"}`+"\n"+
			`{"id":"eq_3","at":"2026-06-18T00:00:00Z","evidence_ref":"docs/incident-lifecycle.md","purpose":"incident_review","quality_level":"blocked","schema_version":"evidence:v1","ontology_version":"concepts:v1","mapping_confidence":"unknown","assessed_by":"governance_steward","reassessment_reasons":["security_incident"]}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SnapshotCount != 3 {
		t.Fatalf("status = %#v", status)
	}
	if status.StaleSnapshotCount != 1 || status.LowQualityCount != 1 || status.BlockedQualityCount != 1 || status.MappingDriftCount != 2 {
		t.Fatalf("debt counters = %#v", status)
	}
	if status.ReassessmentDebtCount != 5 {
		t.Fatalf("expected 5 reassessment debt, got %#v", status)
	}
	if status.ByQualityLevel["high"] != 1 || status.ByMappingConfidence["unknown"] != 1 || status.ByPurpose["incident_review"] != 1 || status.ByReassessmentReason["security_incident"] != 1 {
		t.Fatalf("status maps = %#v %#v %#v %#v", status.ByQualityLevel, status.ByMappingConfidence, status.ByPurpose, status.ByReassessmentReason)
	}
}
