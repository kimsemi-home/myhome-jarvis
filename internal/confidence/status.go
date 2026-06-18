package confidence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

const PolicyRelativePath = "generated/confidence.generated.json"

type Policy struct {
	Context                  string    `json:"context"`
	Version                  string    `json:"version"`
	GeneratedArtifact        string    `json:"generated_artifact"`
	AssessorKey              string    `json:"assessor_key"`
	ConfidenceIsCap          bool      `json:"confidence_is_cap"`
	SelfReportAllowed        bool      `json:"self_report_allowed"`
	PublicStatusRedacted     bool      `json:"public_status_redacted"`
	RawEvidencePublicAllowed bool      `json:"raw_evidence_public_allowed"`
	Levels                   []string  `json:"levels"`
	Inputs                   []string  `json:"inputs"`
	CapRules                 []CapRule `json:"cap_rules"`
	PublicSummaryFields      []string  `json:"public_summary_fields"`
	ForbiddenPublicFields    []string  `json:"forbidden_public_fields"`
	Commands                 []string  `json:"commands"`
}

type CapRule struct {
	Key       string `json:"key"`
	When      string `json:"when"`
	Cap       string `json:"cap"`
	Triggered bool   `json:"triggered,omitempty"`
}

type Inputs struct {
	Evidence     evidence.Status   `json:"-"`
	Quality      qualitylog.Status `json:"-"`
	PublicSafety security.Status   `json:"-"`
}

type Status struct {
	PolicyPath               string `json:"policy_path"`
	AssessorKey              string `json:"assessor_key"`
	LevelCap                 string `json:"level_cap"`
	Blocked                  bool   `json:"blocked"`
	SelfReportAllowed        bool   `json:"self_report_allowed"`
	EvidenceLinkCount        int    `json:"evidence_link_count"`
	DanglingEvidenceRefCount int    `json:"dangling_evidence_ref_count"`
	OpenLearningCount        int    `json:"open_learning_count"`
	QualityRecorded          bool   `json:"quality_recorded"`
	QualityOK                bool   `json:"quality_ok"`
	PublicSafetyOK           bool   `json:"public_safety_ok"`
	ActiveRule               string `json:"active_rule"`
	CheckedAt                string `json:"checked_at"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	evidenceStatus, err := evidence.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	qualityStatus, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	securityStatus, err := security.StatusForRoot(root)
	if err != nil {
		return Status{}, err
	}
	return Assess(policy, Inputs{
		Evidence:     evidenceStatus,
		Quality:      qualityStatus,
		PublicSafety: securityStatus,
	}), nil
}

func Assess(policy Policy, inputs Inputs) Status {
	status := Status{
		PolicyPath:               PolicyRelativePath,
		AssessorKey:              policy.AssessorKey,
		LevelCap:                 "high",
		SelfReportAllowed:        policy.SelfReportAllowed,
		EvidenceLinkCount:        inputs.Evidence.EdgeCount,
		DanglingEvidenceRefCount: inputs.Evidence.DanglingEvidenceRefCount,
		OpenLearningCount:        inputs.Evidence.OpenLearningCount,
		QualityRecorded:          inputs.Quality.Exists && inputs.Quality.Last != nil,
		PublicSafetyOK:           inputs.PublicSafety.OK,
		CheckedAt:                time.Now().UTC().Format(time.RFC3339),
	}
	if inputs.Quality.Last != nil {
		status.QualityOK = inputs.Quality.Last.OK
	}
	for _, rule := range policy.CapRules {
		rule.Key = normalizeToken(rule.Key)
		rule.When = normalizeToken(rule.When)
		rule.Cap = normalizeToken(rule.Cap)
		rule.Triggered = ruleTriggered(rule.When, status)
		if rule.Triggered {
			status.LevelCap = minLevel(status.LevelCap, rule.Cap)
			if status.ActiveRule == "" {
				status.ActiveRule = rule.Key
			}
		}
	}
	if status.LevelCap == "" {
		status.LevelCap = "high"
	}
	status.Blocked = status.LevelCap == "blocked"
	return status
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
		return fmt.Errorf("confidence policy context = %q", policy.Context)
	}
	if policy.AssessorKey != "confidence_assessor" {
		return fmt.Errorf("confidence assessor key must be confidence_assessor")
	}
	if !policy.ConfidenceIsCap || policy.SelfReportAllowed {
		return fmt.Errorf("confidence must be an external cap and self-reporting must stay disabled")
	}
	if !policy.PublicStatusRedacted || policy.RawEvidencePublicAllowed {
		return fmt.Errorf("confidence public status must be redacted")
	}
	levels := normalizeList(policy.Levels)
	for _, level := range []string{"blocked", "low", "medium", "high"} {
		if !contains(levels, level) {
			return fmt.Errorf("confidence level %q is missing", level)
		}
	}
	inputs := normalizeList(policy.Inputs)
	for _, input := range []string{"evidence_graph", "learning_ledger", "quality_gate", "public_safety"} {
		if !contains(inputs, input) {
			return fmt.Errorf("confidence input %q is missing", input)
		}
	}
	if len(policy.CapRules) == 0 {
		return fmt.Errorf("confidence cap rules are required")
	}
	for _, rule := range policy.CapRules {
		if normalizeToken(rule.Key) == "" || normalizeToken(rule.When) == "" {
			return fmt.Errorf("confidence cap rule key and condition are required")
		}
		if !contains(levels, normalizeToken(rule.Cap)) {
			return fmt.Errorf("confidence cap rule %q has invalid cap", rule.Key)
		}
	}
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range []string{"level_cap", "self_report_allowed", "evidence_link_count", "public_safety_ok", "checked_at"} {
		if !contains(summary, field) {
			return fmt.Errorf("confidence public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj confidence status") {
		return fmt.Errorf("confidence status command is missing")
	}
	return nil
}

func ruleTriggered(condition string, status Status) bool {
	switch condition {
	case "public_safety_not_ok":
		return !status.PublicSafetyOK
	case "latest_quality_failed":
		return status.QualityRecorded && !status.QualityOK
	case "evidence_edge_count_zero":
		return status.EvidenceLinkCount == 0
	case "dangling_evidence_ref_count_positive":
		return status.DanglingEvidenceRefCount > 0
	case "open_learning_count_positive":
		return status.OpenLearningCount > 0
	case "latest_quality_missing":
		return !status.QualityRecorded
	case "evidence_links_and_verification_clear":
		return status.PublicSafetyOK && status.QualityRecorded && status.QualityOK && status.EvidenceLinkCount > 0 && status.DanglingEvidenceRefCount == 0 && status.OpenLearningCount == 0
	default:
		return false
	}
}

func minLevel(left string, right string) string {
	if levelRank(right) < levelRank(left) {
		return right
	}
	return left
}

func levelRank(level string) int {
	switch normalizeToken(level) {
	case "blocked":
		return 0
	case "low":
		return 1
	case "medium":
		return 2
	case "high":
		return 3
	default:
		return 3
	}
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
