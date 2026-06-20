package commandcenter

type MergeEvidenceSummary struct {
	PublicSafe                   bool   `json:"public_safe"`
	DefaultBehavior              string `json:"default_behavior"`
	MergePreference              string `json:"merge_preference"`
	PostMergeEvidenceRequired    bool   `json:"post_merge_evidence_required"`
	LinearCompletionRequired     bool   `json:"linear_completion_required"`
	MainQualityRunRequired       bool   `json:"main_quality_run_required"`
	PrivateDataScanRequired      bool   `json:"private_data_scan_required"`
	EligibleGateCount            int    `json:"eligible_gate_count"`
	RequiredEvidenceCount        int    `json:"required_evidence_count"`
	MissingGateCount             int    `json:"missing_gate_count"`
	MissingRequiredEvidenceCount int    `json:"missing_required_evidence_count"`
	MergeReady                   bool   `json:"merge_ready"`
	MergeBlockedUntilEvidence    bool   `json:"merge_blocked_until_evidence"`
}
