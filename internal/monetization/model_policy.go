package monetization

type Policy struct {
	Context                  string   `json:"context"`
	Version                  string   `json:"version"`
	GeneratedArtifact        string   `json:"generated_artifact"`
	PrivateExperimentLedger  string   `json:"private_experiment_ledger"`
	AppendOnly               bool     `json:"append_only"`
	PublicStatusRedacted     bool     `json:"public_status_redacted"`
	RawRevenuePublicAllowed  bool     `json:"raw_revenue_public_allowed"`
	DecisionEvidenceRequired bool     `json:"decision_evidence_required"`
	CostEstimateRequired     bool     `json:"cost_estimate_required"`
	ExperimentStates         []string `json:"experiment_states"`
	DecisionKinds            []string `json:"decision_kinds"`
	ReviewStatuses           []string `json:"review_statuses"`
	ExpectedValueBands       []string `json:"expected_value_bands"`
	CostUnitKinds            []string `json:"cost_unit_kinds"`
	RequiredFields           []string `json:"required_fields"`
	AllowedEvidencePrefixes  []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields      []string `json:"public_summary_fields"`
	ForbiddenPublicFields    []string `json:"forbidden_public_fields"`
	Commands                 []string `json:"commands"`
}
