package confidence

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

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
