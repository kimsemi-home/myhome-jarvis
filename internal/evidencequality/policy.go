package evidencequality

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const PolicyRelativePath = "generated/evidence_quality.generated.json"

type Policy struct {
	Context                  string   `json:"context"`
	Version                  string   `json:"version"`
	GeneratedArtifact        string   `json:"generated_artifact"`
	PrivateSnapshotLedger    string   `json:"private_snapshot_ledger"`
	AppendOnly               bool     `json:"append_only"`
	PublicStatusRedacted     bool     `json:"public_status_redacted"`
	RawSnapshotPublicAllowed bool     `json:"raw_snapshot_public_allowed"`
	StaleAfterHours          int      `json:"stale_after_hours"`
	QualityLevels            []string `json:"quality_levels"`
	MappingConfidenceLevels  []string `json:"mapping_confidence_levels"`
	AllowedPurposes          []string `json:"allowed_purposes"`
	ReassessmentReasons      []string `json:"reassessment_reasons"`
	RequiredFields           []string `json:"required_fields"`
	AllowedEvidencePrefixes  []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields      []string `json:"public_summary_fields"`
	ForbiddenPublicFields    []string `json:"forbidden_public_fields"`
	Commands                 []string `json:"commands"`
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
