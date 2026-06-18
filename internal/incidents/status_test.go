package incidents

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

func TestStatusJSONDoesNotLeakRawIncidentDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/incidents/incidents.jsonl",
		`{"id":"inc_1","at":"2026-06-18T00:00:00Z","kind":"quality_regression","stage":"fix_planned","status":"mitigating","owner_role":"go","quarantine_state":"none","evidence_refs":["data/private/quality/runs.jsonl"],"summary":"private incident summary","root_cause_notes":"private root cause"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range []string{
		"summary",
		"root_cause",
		"private incident",
		"private root cause",
		"evidence_refs",
		"data/private/quality/runs.jsonl",
		"token",
		"secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("incident status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsRawPublicIncidentDetails(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawIncidentPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw incident public policy to fail")
	}
}

func testPolicy() Policy {
	return Policy{
		Context:                   "AgentCluster",
		Version:                   "v1",
		GeneratedArtifact:         "generated/incidents.generated.json",
		PrivateIncidentLedger:     "data/private/incidents/incidents.jsonl",
		AppendOnly:                true,
		PublicStatusRedacted:      true,
		RawIncidentPublicAllowed:  false,
		QuarantineStaleAfterHours: 168,
		AllowedKinds:              []string{"quality_regression", "public_safety", "evidence_gap", "authority_violation", "quarantine", "translation_loss", "control_plane", "feedback_loop_gap"},
		Lifecycle:                 []string{"observed", "evidence_recorded", "classified", "owner_assigned", "fix_planned", "verified", "knowledge_updated"},
		AllowedStatuses:           []string{"open", "mitigating", "verified", "closed", "quarantined"},
		OwnerRoles:                []string{"producer", "independent_reviewer", "adversarial_reviewer", "deterministic_verifier", "governance_steward", "go"},
		QuarantineStates:          []string{"none", "quarantined", "release_requested", "released"},
		RequiredFields:            []string{"at", "kind", "stage", "status", "owner_role", "evidence_refs"},
		AllowedEvidencePrefixes:   []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:       []string{"policy_path", "ledger_path", "exists", "count", "open_count", "closed_count", "invalid_incident_count", "incident_debt_count", "missing_owner_count", "missing_evidence_ref_count", "stale_quarantine_count", "quarantine_stale_after_hours", "by_kind", "by_stage", "by_status", "by_owner_role", "by_quarantine_state", "last_observed_at", "checked_at"},
		Commands:                  []string{"mhj incidents status"},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}
