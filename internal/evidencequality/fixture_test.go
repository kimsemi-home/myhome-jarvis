package evidencequality

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

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
