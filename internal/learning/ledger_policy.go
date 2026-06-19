package learning

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const PolicyRelativePath = "generated/learning.generated.json"

type Policy struct {
	Context                     string   `json:"context"`
	Version                     string   `json:"version"`
	PrivateLedger               string   `json:"private_ledger"`
	GeneratedArtifact           string   `json:"generated_artifact"`
	AppendOnly                  bool     `json:"append_only"`
	PrivateJournalRequired      bool     `json:"private_journal_required"`
	PublicStatusRedacted        bool     `json:"public_status_redacted"`
	RawObservationPublicAllowed bool     `json:"raw_observation_public_allowed"`
	RequiredFields              []string `json:"required_fields"`
	AllowedKinds                []string `json:"allowed_kinds"`
	Lifecycle                   []string `json:"lifecycle"`
	AllowedStatuses             []string `json:"allowed_statuses"`
	AllowedEvidencePrefixes     []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields         []string `json:"public_summary_fields"`
	ForbiddenPrivateMarkers     []string `json:"forbidden_private_markers"`
	Commands                    []string `json:"commands"`
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
