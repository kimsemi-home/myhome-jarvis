package confidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func TestAssessReturnsHighWhenEvidenceAndVerificationAreClear(t *testing.T) {
	status := Assess(testPolicy(), Inputs{
		Evidence: evidence.Status{EdgeCount: 2},
		Quality:  qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{
			OK: true,
		},
	})
	if status.LevelCap != "high" || status.Blocked {
		t.Fatalf("status = %#v", status)
	}
	if status.ActiveRule != "evidence_backed" {
		t.Fatalf("active rule = %q", status.ActiveRule)
	}
}

func TestAssessCapsLowWithoutEvidenceLinks(t *testing.T) {
	status := Assess(testPolicy(), Inputs{
		Evidence:     evidence.Status{EdgeCount: 0},
		Quality:      qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{OK: true},
	})
	if status.LevelCap != "low" || status.ActiveRule != "missing_evidence_links" {
		t.Fatalf("status = %#v", status)
	}
}

func TestAssessBlocksOnPublicSafetyFindings(t *testing.T) {
	status := Assess(testPolicy(), Inputs{
		Evidence:     evidence.Status{EdgeCount: 2},
		Quality:      qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{OK: false},
	})
	if status.LevelCap != "blocked" || !status.Blocked || status.ActiveRule != "public_safety_findings" {
		t.Fatalf("status = %#v", status)
	}
}

func TestAssessCapsMediumForOpenLearningDebt(t *testing.T) {
	status := Assess(testPolicy(), Inputs{
		Evidence:     evidence.Status{EdgeCount: 2, OpenLearningCount: 1},
		Quality:      qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{OK: true},
	})
	if status.LevelCap != "medium" || status.ActiveRule != "open_learning_debt" {
		t.Fatalf("status = %#v", status)
	}
}

func TestStatusJSONDoesNotLeakRawEvidence(t *testing.T) {
	status := Assess(testPolicy(), Inputs{
		Evidence:     evidence.Status{EdgeCount: 1, DanglingEvidenceRefCount: 1},
		Quality:      qualitylog.Status{Exists: true, Last: &qualitylog.Run{OK: true}},
		PublicSafety: security.Status{OK: true},
	})
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, forbidden := range []string{
		"summary",
		"next_action",
		"evidence_refs",
		"raw_prompt",
		"raw_transcript",
		"token",
		"secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(strings.ToLower(body), forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsSelfReporting(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, true)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected self-reporting policy to fail")
	}
}

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
		CapRules: []CapRule{
			{Key: "public_safety_findings", When: "public_safety_not_ok", Cap: "blocked"},
			{Key: "quality_failing", When: "latest_quality_failed", Cap: "blocked"},
			{Key: "missing_evidence_links", When: "evidence_edge_count_zero", Cap: "low"},
			{Key: "dangling_refs", When: "dangling_evidence_ref_count_positive", Cap: "low"},
			{Key: "open_learning_debt", When: "open_learning_count_positive", Cap: "medium"},
			{Key: "quality_unrecorded", When: "latest_quality_missing", Cap: "medium"},
			{Key: "evidence_backed", When: "evidence_links_and_verification_clear", Cap: "high"},
		},
		PublicSummaryFields: []string{"policy_path", "assessor_key", "level_cap", "blocked", "self_report_allowed", "evidence_link_count", "dangling_evidence_ref_count", "open_learning_count", "quality_recorded", "quality_ok", "public_safety_ok", "active_rule", "checked_at"},
		Commands:            []string{"mhj confidence status"},
	}
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
