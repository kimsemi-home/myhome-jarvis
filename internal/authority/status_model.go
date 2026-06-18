package authority

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/confidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/controlplane"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidencequality"
	"github.com/kimsemi-home/myhome-jarvis/internal/incidents"
	"github.com/kimsemi-home/myhome-jarvis/internal/review"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/translation"
)

type Inputs struct {
	Confidence      confidence.Status      `json:"-"`
	EvidenceQuality evidencequality.Status `json:"-"`
	Incidents       incidents.Status       `json:"-"`
	ControlPlane    controlplane.Status    `json:"-"`
	Translation     translation.Status     `json:"-"`
	Review          review.Status          `json:"-"`
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
	HumanReviewDebtCount        int            `json:"human_review_debt_count"`
	HumanReviewCapacityState    string         `json:"human_review_capacity_state"`
	AllowedDecisions            []string       `json:"allowed_decisions"`
	BlockedDecisions            []string       `json:"blocked_decisions"`
	ByRisk                      map[string]int `json:"by_risk"`
	CheckedAt                   string         `json:"checked_at"`
}
