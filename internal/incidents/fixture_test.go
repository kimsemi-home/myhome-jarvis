package incidents

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

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
