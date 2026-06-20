package codexcost

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const PolicyRelativePath = "generated/codex_cost.generated.json"

type Policy struct {
	Context                 string   `json:"context"`
	Version                 string   `json:"version"`
	GeneratedArtifact       string   `json:"generated_artifact"`
	PrivateUsageLedger      string   `json:"private_usage_ledger"`
	AppendOnly              bool     `json:"append_only"`
	PublicStatusRedacted    bool     `json:"public_status_redacted"`
	RawUsagePublicAllowed   bool     `json:"raw_usage_public_allowed"`
	SemanticHashInputs      []string `json:"semantic_hash_inputs"`
	UnitKinds               []string `json:"unit_kinds"`
	LoopScopes              []string `json:"loop_scopes"`
	RecordStatuses          []string `json:"record_statuses"`
	WarningUnitThreshold    int64    `json:"warning_unit_threshold"`
	ReviewUnitThreshold     int64    `json:"review_unit_threshold"`
	RequiredFields          []string `json:"required_fields"`
	AllowedEvidencePrefixes []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields     []string `json:"public_summary_fields"`
	ForbiddenPublicFields   []string `json:"forbidden_public_fields"`
	Commands                []string `json:"commands"`
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
