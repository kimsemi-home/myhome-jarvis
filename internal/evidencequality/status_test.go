package evidencequality

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

func TestStatusJSONDoesNotLeakRawSnapshotDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/evidence-quality/snapshots.jsonl",
		`{"id":"eq_1","at":"2026-06-18T00:00:00Z","evidence_ref":"data/private/quality/runs.jsonl","purpose":"release_gate","quality_level":"medium","schema_version":"quality:v1","ontology_version":"concepts:v1","mapping_confidence":"medium","assessed_by":"go","reassessment_reasons":["schema_version_change"],"raw_notes":"private quality notes","raw_evidence":"private evidence body"}`+"\n")

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
		"raw_notes",
		"raw_evidence",
		"evidence_ref",
		"private quality",
		"private evidence",
		"data/private/quality/runs.jsonl",
		"token",
		"secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("evidence quality status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsRawPublicSnapshots(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawSnapshotPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw public snapshot policy to fail")
	}
}

func testPolicy() Policy {
	return Policy{
		Context:                  "AgentCluster",
		Version:                  "v1",
		GeneratedArtifact:        "generated/evidence_quality.generated.json",
		PrivateSnapshotLedger:    "data/private/evidence-quality/snapshots.jsonl",
		AppendOnly:               true,
		PublicStatusRedacted:     true,
		RawSnapshotPublicAllowed: false,
		StaleAfterHours:          720,
		QualityLevels:            []string{"high", "medium", "low", "blocked"},
		MappingConfidenceLevels:  []string{"high", "medium", "low", "unknown"},
		AllowedPurposes:          []string{"root_cause", "confidence_assessment", "incident_review", "release_gate", "conformance", "revalidation", "quarantine_release", "semantic_debug"},
		ReassessmentReasons:      []string{"age", "schema_version_change", "ontology_version_change", "counter_evidence", "agent_reliability_drop", "security_incident", "quarantine", "translation_loss"},
		RequiredFields:           []string{"at", "evidence_ref", "purpose", "quality_level", "schema_version", "ontology_version", "mapping_confidence", "assessed_by", "reassessment_reasons"},
		AllowedEvidencePrefixes:  []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:      []string{"policy_path", "ledger_path", "exists", "snapshot_count", "invalid_snapshot_count", "reassessment_debt_count", "missing_evidence_count", "stale_snapshot_count", "low_quality_count", "blocked_quality_count", "mapping_drift_count", "stale_after_hours", "by_quality_level", "by_mapping_confidence", "by_purpose", "by_reassessment_reason", "last_observed_at", "checked_at"},
		Commands:                 []string{"mhj evidence-quality status"},
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
