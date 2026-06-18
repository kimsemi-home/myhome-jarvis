package authority

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

const PolicyRelativePath = "generated/authority.generated.json"

type Policy struct {
	Context                     string           `json:"context"`
	Version                     string           `json:"version"`
	GeneratedArtifact           string           `json:"generated_artifact"`
	PublicStatusRedacted        bool             `json:"public_status_redacted"`
	SelfAuthorityAllowed        bool             `json:"self_authority_allowed"`
	ReasoningTierGrantsApproval bool             `json:"reasoning_tier_grants_approval"`
	PublicRepoHighRiskBlocked   bool             `json:"public_repo_high_risk_blocked"`
	RequiredInputs              []string         `json:"required_inputs"`
	ReasoningTiers              []ReasoningTier  `json:"reasoning_tiers"`
	RolePermissions             []RolePermission `json:"role_permissions"`
	DomainAttributes            []string         `json:"domain_attributes"`
	Decisions                   []Decision       `json:"decisions"`
	Outcomes                    []string         `json:"outcomes"`
	AuthorityDebtClasses        []string         `json:"authority_debt_classes"`
	PublicSummaryFields         []string         `json:"public_summary_fields"`
	ForbiddenPublicFields       []string         `json:"forbidden_public_fields"`
	Commands                    []string         `json:"commands"`
}

type ReasoningTier struct {
	Key     string   `json:"key"`
	Label   string   `json:"label"`
	May     []string `json:"may"`
	MustNot []string `json:"must_not"`
}

type RolePermission struct {
	Role    string   `json:"role"`
	May     []string `json:"may"`
	MustNot []string `json:"must_not"`
}

type Decision struct {
	Key                 string `json:"key"`
	Risk                string `json:"risk"`
	PublicRepoAllowed   bool   `json:"public_repo_allowed"`
	RequiresHumanReview bool   `json:"requires_human_review"`
	AllowedWhenBlocked  bool   `json:"allowed_when_blocked"`
}

type Inputs struct {
	Confidence      confidence.Status      `json:"-"`
	EvidenceQuality evidencequality.Status `json:"-"`
	Incidents       incidents.Status       `json:"-"`
	ControlPlane    controlplane.Status    `json:"-"`
	Translation     translation.Status     `json:"-"`
	PublicSafety    security.Status        `json:"-"`
}

type Status struct {
	PolicyPath                  string         `json:"policy_path"`
	Outcome                     string         `json:"outcome"`
	ActiveRule                  string         `json:"active_rule"`
	InputCount                  int            `json:"input_count"`
	DecisionCount               int            `json:"decision_count"`
	AllowedDecisionCount        int            `json:"allowed_decision_count"`
	BlockedDecisionCount        int            `json:"blocked_decision_count"`
	AuthorityDebtCount          int            `json:"authority_debt_count"`
	PublicRepoMode              bool           `json:"public_repo_mode"`
	ReasoningTierGrantsApproval bool           `json:"reasoning_tier_grants_approval"`
	SelfAuthorityAllowed        bool           `json:"self_authority_allowed"`
	PublicSafetyOK              bool           `json:"public_safety_ok"`
	ConfidenceCap               string         `json:"confidence_cap"`
	EvidenceQualityDebtCount    int            `json:"evidence_quality_debt_count"`
	IncidentDebtCount           int            `json:"incident_debt_count"`
	ControlPlaneDebtCount       int            `json:"control_plane_debt_count"`
	TranslationDebtCount        int            `json:"translation_debt_count"`
	AllowedDecisions            []string       `json:"allowed_decisions"`
	BlockedDecisions            []string       `json:"blocked_decisions"`
	ByRisk                      map[string]int `json:"by_risk"`
	CheckedAt                   string         `json:"checked_at"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	confidenceStatus, err := confidence.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	evidenceQualityStatus, err := evidencequality.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	incidentStatus, err := incidents.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	controlPlaneStatus, err := controlplane.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	translationStatus, err := translation.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	publicSafetyStatus, err := security.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	return Assess(policy, Inputs{
		Confidence:      confidenceStatus,
		EvidenceQuality: evidenceQualityStatus,
		Incidents:       incidentStatus,
		ControlPlane:    controlPlaneStatus,
		Translation:     translationStatus,
		PublicSafety:    publicSafetyStatus,
	}), nil
}

func Assess(policy Policy, inputs Inputs) Status {
	status := Status{
		PolicyPath:                  PolicyRelativePath,
		InputCount:                  len(policy.RequiredInputs),
		DecisionCount:               len(policy.Decisions),
		PublicRepoMode:              policy.PublicRepoHighRiskBlocked,
		ReasoningTierGrantsApproval: policy.ReasoningTierGrantsApproval,
		SelfAuthorityAllowed:        policy.SelfAuthorityAllowed,
		PublicSafetyOK:              inputs.PublicSafety.OK,
		ConfidenceCap:               normalizeToken(inputs.Confidence.LevelCap),
		EvidenceQualityDebtCount:    inputs.EvidenceQuality.ReassessmentDebtCount,
		IncidentDebtCount:           inputs.Incidents.IncidentDebtCount,
		ControlPlaneDebtCount:       inputs.ControlPlane.ManifestDebtCount,
		TranslationDebtCount:        inputs.Translation.OpenDebtCount + inputs.Translation.ForbiddenLossCount,
		ByRisk:                      map[string]int{},
		CheckedAt:                   time.Now().UTC().Format(time.RFC3339),
	}
	if status.ConfidenceCap == "" {
		status.ConfidenceCap = "unknown"
	}
	status.AuthorityDebtCount = status.EvidenceQualityDebtCount +
		status.IncidentDebtCount +
		status.ControlPlaneDebtCount +
		status.TranslationDebtCount

	status.Outcome, status.ActiveRule = authorityOutcome(status, inputs)
	for _, decision := range normalizedDecisions(policy.Decisions) {
		status.ByRisk[decision.Risk]++
		if decisionAllowed(status.Outcome, decision) {
			status.AllowedDecisions = append(status.AllowedDecisions, decision.Key)
		} else {
			status.BlockedDecisions = append(status.BlockedDecisions, decision.Key)
		}
	}
	sort.Strings(status.AllowedDecisions)
	sort.Strings(status.BlockedDecisions)
	status.AllowedDecisionCount = len(status.AllowedDecisions)
	status.BlockedDecisionCount = len(status.BlockedDecisions)
	return status
}

func authorityOutcome(status Status, inputs Inputs) (string, string) {
	if !status.PublicSafetyOK {
		return "blocked", "public_safety_not_ok"
	}
	if inputs.Confidence.Blocked {
		return "blocked", "confidence_blocked"
	}
	if status.ConfidenceCap == "low" || status.ConfidenceCap == "unknown" {
		return "blocked", "confidence_cap_low"
	}
	if inputs.Translation.ForbiddenLossCount > 0 {
		return "blocked", "forbidden_translation_loss"
	}
	if status.AuthorityDebtCount > 0 {
		switch {
		case status.EvidenceQualityDebtCount > 0:
			return "review_required", "evidence_quality_debt"
		case status.IncidentDebtCount > 0:
			return "review_required", "incident_debt"
		case status.ControlPlaneDebtCount > 0:
			return "review_required", "control_plane_debt"
		case status.TranslationDebtCount > 0:
			return "review_required", "translation_debt"
		}
	}
	return "limited", "public_repo_high_risk_block"
}

func decisionAllowed(outcome string, decision Decision) bool {
	switch outcome {
	case "blocked":
		return decision.AllowedWhenBlocked
	case "review_required":
		return decision.AllowedWhenBlocked || (decision.PublicRepoAllowed && !decision.RequiresHumanReview && decision.Risk == "low")
	default:
		return decision.PublicRepoAllowed && decision.Risk != "high"
	}
}

func ReadPolicy(root string) (Policy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(PolicyRelativePath)))
	if err != nil {
		return Policy{}, err
	}
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		return Policy{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("authority policy context = %q", policy.Context)
	}
	if !policy.PublicStatusRedacted {
		return fmt.Errorf("authority public status must stay redacted")
	}
	if policy.SelfAuthorityAllowed || policy.ReasoningTierGrantsApproval {
		return fmt.Errorf("authority policy must not allow self-authority or reasoning-tier approval")
	}
	if !policy.PublicRepoHighRiskBlocked {
		return fmt.Errorf("authority policy must block high-risk decisions in public repo mode")
	}
	inputs := normalizeList(policy.RequiredInputs)
	for _, input := range []string{"confidence_assessor", "evidence_quality", "incident_lifecycle", "control_plane", "translation", "public_safety"} {
		if !contains(inputs, input) {
			return fmt.Errorf("authority input %q is missing", input)
		}
	}
	tiers := mapByKey(policy.ReasoningTiers)
	for _, tier := range []string{"r0_compiler", "r1_low", "r2_medium", "r3_high", "r4_governance"} {
		if _, ok := tiers[tier]; !ok {
			return fmt.Errorf("authority reasoning tier %q is missing", tier)
		}
	}
	roles := mapRolePermissions(policy.RolePermissions)
	for _, role := range []string{"producer", "independent_reviewer", "adversarial_reviewer", "deterministic_verifier", "governance_steward"} {
		if _, ok := roles[role]; !ok {
			return fmt.Errorf("authority role %q is missing", role)
		}
	}
	attributes := normalizeList(policy.DomainAttributes)
	for _, attribute := range []string{"agent_reliability", "reasoning_tier", "ontology_maturity", "evidence_quality", "security_impact", "data_sensitivity", "change_risk", "verification_scope", "lease_status", "quarantine_state", "human_review_capacity"} {
		if !contains(attributes, attribute) {
			return fmt.Errorf("authority domain attribute %q is missing", attribute)
		}
	}
	decisions := normalizedDecisions(policy.Decisions)
	if len(decisions) == 0 {
		return fmt.Errorf("authority decisions are required")
	}
	decisionMap := mapDecisions(decisions)
	for _, key := range []string{"major_ontology_change", "security_boundary_change", "production_change", "evidence_pruning", "quarantine_release", "high_risk_automation"} {
		decision, ok := decisionMap[key]
		if !ok {
			return fmt.Errorf("authority high-risk decision %q is missing", key)
		}
		if decision.PublicRepoAllowed || decision.Risk != "high" {
			return fmt.Errorf("authority high-risk decision %q must stay blocked", key)
		}
	}
	for _, decision := range decisions {
		if decision.Key == "" {
			return fmt.Errorf("authority decision key is required")
		}
		if !contains([]string{"low", "medium", "high"}, decision.Risk) {
			return fmt.Errorf("authority decision %q has invalid risk", decision.Key)
		}
		if decision.Risk == "high" && decision.PublicRepoAllowed {
			return fmt.Errorf("authority high-risk decision %q must not be public-repo allowed", decision.Key)
		}
	}
	outcomes := normalizeList(policy.Outcomes)
	for _, outcome := range []string{"limited", "review_required", "blocked"} {
		if !contains(outcomes, outcome) {
			return fmt.Errorf("authority outcome %q is missing", outcome)
		}
	}
	classes := normalizeList(policy.AuthorityDebtClasses)
	for _, class := range []string{"public_safety", "confidence_cap", "evidence_quality", "incident", "control_plane", "translation", "human_review"} {
		if !contains(classes, class) {
			return fmt.Errorf("authority debt class %q is missing", class)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"outcome", "active_rule", "allowed_decision_count", "blocked_decision_count", "authority_debt_count", "public_repo_mode", "reasoning_tier_grants_approval", "self_authority_allowed", "public_safety_ok", "confidence_cap", "allowed_decisions", "blocked_decisions", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("authority public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj authority status") {
		return fmt.Errorf("authority status command is missing")
	}
	return nil
}

func normalizedDecisions(decisions []Decision) []Decision {
	normalized := make([]Decision, 0, len(decisions))
	for _, decision := range decisions {
		decision.Key = normalizeToken(decision.Key)
		decision.Risk = normalizeToken(decision.Risk)
		if decision.Key == "" {
			continue
		}
		normalized = append(normalized, decision)
	}
	sort.Slice(normalized, func(i, j int) bool {
		return normalized[i].Key < normalized[j].Key
	})
	return normalized
}

func mapDecisions(decisions []Decision) map[string]Decision {
	mapped := map[string]Decision{}
	for _, decision := range decisions {
		mapped[decision.Key] = decision
	}
	return mapped
}

func mapByKey(tiers []ReasoningTier) map[string]ReasoningTier {
	mapped := map[string]ReasoningTier{}
	for _, tier := range tiers {
		tier.Key = normalizeToken(tier.Key)
		if tier.Key != "" {
			mapped[tier.Key] = tier
		}
	}
	return mapped
}

func mapRolePermissions(roles []RolePermission) map[string]RolePermission {
	mapped := map[string]RolePermission{}
	for _, role := range roles {
		role.Role = normalizeToken(role.Role)
		if role.Role != "" {
			mapped[role.Role] = role
		}
	}
	return mapped
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
