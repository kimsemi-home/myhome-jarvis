package incidents

import "testing"

func TestMissingLedgerReturnsZeroRedactedStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.Count != 0 || status.IncidentDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.LedgerPath != "data/private/incidents/incidents.jsonl" || status.PolicyPath != PolicyRelativePath {
		t.Fatalf("paths = %#v", status)
	}
}

func TestStatusCountsValidAndStaleQuarantine(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/incidents/incidents.jsonl",
		`{"id":"inc_1","at":"2026-06-01T00:00:00Z","kind":"quarantine","stage":"owner_assigned","status":"quarantined","owner_role":"governance_steward","quarantine_state":"quarantined","evidence_refs":["generated/incidents.generated.json"],"summary":"private"}`+"\n"+
			`{"id":"inc_2","at":"2026-06-18T00:00:00Z","kind":"evidence_gap","stage":"knowledge_updated","status":"closed","owner_role":"deterministic_verifier","quarantine_state":"released","evidence_refs":["docs/incident-lifecycle.md"],"root_cause_notes":"private"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Count != 2 || status.OpenCount != 1 || status.ClosedCount != 1 {
		t.Fatalf("status = %#v", status)
	}
	if status.StaleQuarantineCount != 1 || status.IncidentDebtCount != 1 {
		t.Fatalf("expected stale quarantine debt, got %#v", status)
	}
	if status.ByKind["quarantine"] != 1 || status.ByStage["knowledge_updated"] != 1 || status.ByOwnerRole["governance_steward"] != 1 {
		t.Fatalf("status maps = %#v %#v %#v", status.ByKind, status.ByStage, status.ByOwnerRole)
	}
}
