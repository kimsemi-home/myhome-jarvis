package mergeevidence

type Policy struct {
	Context                   string   `json:"context"`
	Version                   string   `json:"version"`
	GeneratedArtifact         string   `json:"generated_artifact"`
	DefaultBehavior           string   `json:"default_behavior"`
	MergePreference           string   `json:"merge_preference"`
	PublicStatusRedacted      bool     `json:"public_status_redacted"`
	MergeWithoutReviewAllowed bool     `json:"merge_without_review_allowed"`
	PersistPrivateEvidence    bool     `json:"persist_private_evidence"`
	PostMergeEvidenceRequired bool     `json:"post_merge_evidence_required"`
	LinearCompletionRequired  bool     `json:"linear_completion_required"`
	MainQualityRunRequired    bool     `json:"main_quality_run_required"`
	PrivateDataScanRequired   bool     `json:"private_data_scan_required"`
	Gates                     []Gate   `json:"gates"`
	RequiredEvidence          []string `json:"required_evidence"`
	PublicSummaryFields       []string `json:"public_summary_fields"`
	ForbiddenPublicFields     []string `json:"forbidden_public_fields"`
	Commands                  []string `json:"commands"`
}

type Gate struct {
	Key         string `json:"key"`
	Label       string `json:"label"`
	Evidence    string `json:"evidence"`
	Required    bool   `json:"required"`
	BlocksMerge bool   `json:"blocks_merge"`
}

type Status struct {
	Context                      string   `json:"context"`
	Version                      string   `json:"version"`
	PolicyPath                   string   `json:"policy_path"`
	DefaultBehavior              string   `json:"default_behavior"`
	MergePreference              string   `json:"merge_preference"`
	PublicSafe                   bool     `json:"public_safe"`
	Redaction                    string   `json:"redaction"`
	PostMergeEvidenceRequired    bool     `json:"post_merge_evidence_required"`
	LinearCompletionRequired     bool     `json:"linear_completion_required"`
	MainQualityRunRequired       bool     `json:"main_quality_run_required"`
	PrivateDataScanRequired      bool     `json:"private_data_scan_required"`
	EligibleGateCount            int      `json:"eligible_gate_count"`
	RequiredEvidenceCount        int      `json:"required_evidence_count"`
	MissingGateCount             int      `json:"missing_gate_count"`
	MissingRequiredEvidenceCount int      `json:"missing_required_evidence_count"`
	MergeReady                   bool     `json:"merge_ready"`
	MergeBlockedUntilEvidence    bool     `json:"merge_blocked_until_evidence"`
	GateKeys                     []string `json:"gate_keys"`
	RequiredEvidence             []string `json:"required_evidence"`
	Commands                     []string `json:"commands"`
	CheckedAt                    string   `json:"checked_at"`
}
