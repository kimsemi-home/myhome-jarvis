package incidents

import "testing"

func TestStatusCountsMalformedMissingOwnerAndMissingEvidenceDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/incidents/incidents.jsonl",
		`{`+"\n"+
			`{"id":"inc_2","at":"2026-06-18T00:00:00Z","kind":"evidence_gap","stage":"classified","status":"open","owner_role":"","evidence_refs":["generated/incidents.generated.json"]}`+"\n"+
			`{"id":"inc_3","at":"2026-06-18T00:00:00Z","kind":"evidence_gap","stage":"classified","status":"open","owner_role":"go","evidence_refs":[]}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidIncidentCount != 1 || status.MissingOwnerCount != 1 || status.MissingEvidenceRefCount != 1 || status.IncidentDebtCount != 3 {
		t.Fatalf("status = %#v", status)
	}
}
