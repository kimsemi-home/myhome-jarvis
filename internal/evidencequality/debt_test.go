package evidencequality

import "testing"

func TestStatusCountsMalformedAndMissingEvidenceRefDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/evidence-quality/snapshots.jsonl",
		`{`+"\n"+
			`{"id":"eq_2","at":"2026-06-18T00:00:00Z","evidence_ref":"","purpose":"root_cause","quality_level":"high","schema_version":"evidence:v1","ontology_version":"concepts:v1","mapping_confidence":"high","assessed_by":"go"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidSnapshotCount != 1 || status.MissingEvidenceCount != 1 || status.ReassessmentDebtCount != 2 {
		t.Fatalf("status = %#v", status)
	}
}
