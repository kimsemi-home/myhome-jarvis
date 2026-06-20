package codexsustainability

type Policy struct {
	Context                              string   `json:"context"`
	Version                              string   `json:"version"`
	GeneratedArtifact                    string   `json:"generated_artifact"`
	PrivateEvidenceLedger                string   `json:"private_evidence_ledger"`
	AppendOnly                           bool     `json:"append_only"`
	PublicStatusRedacted                 bool     `json:"public_status_redacted"`
	RawEvidencePublicAllowed             bool     `json:"raw_evidence_public_allowed"`
	TrendBaselinesVersioned              bool     `json:"trend_baselines_versioned"`
	EvidenceMaxAgeHours                  int      `json:"evidence_max_age_hours"`
	TrendBaselineMaxAgeHours             int      `json:"trend_baseline_max_age_hours"`
	CostPerAcceptedChangeReviewThreshold int64    `json:"cost_per_accepted_change_review_threshold"`
	RecordKinds                          []string `json:"record_kinds"`
	Metrics                              []string `json:"metrics"`
	RequiredFields                       []string `json:"required_fields"`
	ProposalRequiredFields               []string `json:"proposal_required_fields"`
	AllowedEvidencePrefixes              []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields                  []string `json:"public_summary_fields"`
	ForbiddenPublicFields                []string `json:"forbidden_public_fields"`
	Commands                             []string `json:"commands"`
}
