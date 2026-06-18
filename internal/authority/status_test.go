package authority

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

func TestAssessKeepsHighRiskDecisionsBlockedInPublicRepoMode(t *testing.T) {
	status := Assess(testPolicy(), clearInputs())

	if status.Outcome != "limited" || status.ActiveRule != "public_repo_high_risk_block" {
		t.Fatalf("status = %#v", status)
	}
	if status.ReasoningTierGrantsApproval || status.SelfAuthorityAllowed {
		t.Fatalf("authority flags = %#v", status)
	}
	if !contains(status.AllowedDecisions, "read_status") || !contains(status.AllowedDecisions, "low_risk_fixture_change") {
		t.Fatalf("allowed decisions = %#v", status.AllowedDecisions)
	}
	for _, blocked := range []string{"major_ontology_change", "security_boundary_change", "production_change", "evidence_pruning", "quarantine_release", "high_risk_automation"} {
		if !contains(status.BlockedDecisions, blocked) {
			t.Fatalf("blocked decisions missing %q in %#v", blocked, status.BlockedDecisions)
		}
	}
}

func TestAssessBlocksOnPublicSafetyOrLowConfidence(t *testing.T) {
	inputs := clearInputs()
	inputs.PublicSafety.OK = false
	status := Assess(testPolicy(), inputs)
	if status.Outcome != "blocked" || status.ActiveRule != "public_safety_not_ok" {
		t.Fatalf("status = %#v", status)
	}

	inputs = clearInputs()
	inputs.Confidence.LevelCap = "low"
	status = Assess(testPolicy(), inputs)
	if status.Outcome != "blocked" || status.ActiveRule != "confidence_cap_low" {
		t.Fatalf("status = %#v", status)
	}
}

func TestAssessRequiresReviewWhenAuthorityDebtExists(t *testing.T) {
	inputs := clearInputs()
	inputs.EvidenceQuality.ReassessmentDebtCount = 2
	inputs.Incidents.IncidentDebtCount = 1

	status := Assess(testPolicy(), inputs)
	if status.Outcome != "review_required" || status.ActiveRule != "evidence_quality_debt" {
		t.Fatalf("status = %#v", status)
	}
	if status.AuthorityDebtCount != 3 || status.EvidenceQualityDebtCount != 2 || status.IncidentDebtCount != 1 {
		t.Fatalf("debt counts = %#v", status)
	}
	if contains(status.AllowedDecisions, "low_risk_fixture_change") {
		t.Fatalf("fixture change should require review while debt exists: %#v", status.AllowedDecisions)
	}
}

func TestStatusForRootReturnsRedactedAuthoritySummary(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, expected := range []string{
		`"policy_path":"generated/authority.generated.json"`,
		`"reasoning_tier_grants_approval":false`,
		`"self_authority_allowed":false`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		"raw_rationale",
		"raw_evidence",
		"evidence_ref",
		"evidence_refs",
		"raw_prompt",
		"raw_transcript",
		"token",
		"secret",
		"credential",
		"linear.app",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("authority status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsReasoningTierApproval(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.ReasoningTierGrantsApproval = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected reasoning tier approval policy to fail")
	}
}

func clearInputs() Inputs {
	return Inputs{
		Confidence: confidence.Status{
			LevelCap:       "high",
			PublicSafetyOK: true,
		},
		EvidenceQuality: evidencequality.Status{},
		Incidents:       incidents.Status{},
		ControlPlane:    controlplane.Status{},
		Translation:     translation.Status{},
		PublicSafety:    security.Status{OK: true, CurrentOK: true, HistoryOK: true},
	}
}

func testPolicy() Policy {
	return Policy{
		Context:                     "AgentCluster",
		Version:                     "v1",
		GeneratedArtifact:           "generated/authority.generated.json",
		PublicStatusRedacted:        true,
		SelfAuthorityAllowed:        false,
		ReasoningTierGrantsApproval: false,
		PublicRepoHighRiskBlocked:   true,
		RequiredInputs:              []string{"confidence_assessor", "evidence_quality", "incident_lifecycle", "control_plane", "translation", "public_safety"},
		ReasoningTiers: []ReasoningTier{
			{Key: "r0_compiler", May: []string{"deterministic_transform"}, MustNot: []string{"approve"}},
			{Key: "r1_low", May: []string{"small_candidate"}, MustNot: []string{"approve"}},
			{Key: "r2_medium", May: []string{"multi_file_candidate"}, MustNot: []string{"approve"}},
			{Key: "r3_high", May: []string{"root_cause_candidate"}, MustNot: []string{"approve"}},
			{Key: "r4_governance", May: []string{"policy_review"}, MustNot: []string{"solo_approve"}},
		},
		RolePermissions: []RolePermission{
			{Role: "producer", May: []string{"propose"}, MustNot: []string{"self_approve"}},
			{Role: "independent_reviewer", May: []string{"review_mapping"}, MustNot: []string{"self_approve"}},
			{Role: "adversarial_reviewer", May: []string{"challenge_evidence"}, MustNot: []string{"self_approve"}},
			{Role: "deterministic_verifier", May: []string{"run_checks"}, MustNot: []string{"approve_semantics"}},
			{Role: "governance_steward", May: []string{"gate_authority"}, MustNot: []string{"solo_major_ontology_change"}},
		},
		DomainAttributes:     []string{"agent_reliability", "reasoning_tier", "ontology_maturity", "evidence_quality", "security_impact", "data_sensitivity", "change_risk", "verification_scope", "lease_status", "quarantine_state", "human_review_capacity"},
		Decisions:            testDecisions(),
		Outcomes:             []string{"limited", "review_required", "blocked"},
		AuthorityDebtClasses: []string{"public_safety", "confidence_cap", "evidence_quality", "incident", "control_plane", "translation", "human_review"},
		PublicSummaryFields:  []string{"policy_path", "outcome", "active_rule", "input_count", "decision_count", "allowed_decision_count", "blocked_decision_count", "authority_debt_count", "public_repo_mode", "reasoning_tier_grants_approval", "self_authority_allowed", "public_safety_ok", "confidence_cap", "evidence_quality_debt_count", "incident_debt_count", "control_plane_debt_count", "translation_debt_count", "allowed_decisions", "blocked_decisions", "by_risk", "checked_at"},
		Commands:             []string{"mhj authority status"},
	}
}

func testDecisions() []Decision {
	return []Decision{
		{Key: "read_status", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "evidence_collection", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "deterministic_verification", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "low_risk_fixture_change", Risk: "medium", PublicRepoAllowed: true},
		{Key: "incident_response", Risk: "medium", PublicRepoAllowed: true, RequiresHumanReview: true, AllowedWhenBlocked: true},
		{Key: "major_ontology_change", Risk: "high", RequiresHumanReview: true},
		{Key: "security_boundary_change", Risk: "high", RequiresHumanReview: true},
		{Key: "production_change", Risk: "high", RequiresHumanReview: true},
		{Key: "evidence_pruning", Risk: "high", RequiresHumanReview: true},
		{Key: "quarantine_release", Risk: "high", RequiresHumanReview: true},
		{Key: "high_risk_automation", Risk: "high", RequiresHumanReview: true},
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not locate repo root")
		}
		dir = next
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
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
