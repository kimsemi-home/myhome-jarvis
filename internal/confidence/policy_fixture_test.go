package confidence

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
		GeneratedArtifact:        "generated/confidence.generated.json",
		AssessorKey:              "confidence_assessor",
		ConfidenceIsCap:          true,
		SelfReportAllowed:        false,
		PublicStatusRedacted:     true,
		RawEvidencePublicAllowed: false,
		Levels:                   []string{"blocked", "low", "medium", "high"},
		Inputs:                   []string{"evidence_graph", "learning_ledger", "quality_gate", "public_safety"},
		CapRules:                 testCapRules(),
		PublicSummaryFields:      testPublicSummaryFields(),
		Commands:                 []string{"mhj confidence status"},
	}
}

func testCapRules() []CapRule {
	return []CapRule{
		{Key: "public_safety_findings", When: "public_safety_not_ok", Cap: "blocked"},
		{Key: "quality_failing", When: "latest_quality_failed", Cap: "blocked"},
		{Key: "missing_evidence_links", When: "evidence_edge_count_zero", Cap: "low"},
		{Key: "dangling_refs", When: "dangling_evidence_ref_count_positive", Cap: "low"},
		{Key: "open_learning_debt", When: "open_learning_count_positive", Cap: "medium"},
		{Key: "quality_unrecorded", When: "latest_quality_missing", Cap: "medium"},
		{Key: "evidence_backed", When: "evidence_links_and_verification_clear", Cap: "high"},
	}
}

func testPublicSummaryFields() []string {
	return []string{"policy_path", "assessor_key", "level_cap", "blocked", "self_report_allowed", "evidence_link_count", "dangling_evidence_ref_count", "open_learning_count", "quality_recorded", "quality_ok", "public_safety_ok", "active_rule", "checked_at"}
}

func writePolicy(t *testing.T, root string, selfReport bool) {
	t.Helper()
	policy := testPolicy()
	policy.SelfReportAllowed = selfReport
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}
