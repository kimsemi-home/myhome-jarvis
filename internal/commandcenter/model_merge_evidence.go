package commandcenter

type MergeEvidenceSummary struct {
	PublicSafe                   bool   `json:"public_safe"`
	DefaultBehavior              string `json:"default_behavior"`
	EligibleGateCount            int    `json:"eligible_gate_count"`
	RequiredEvidenceCount        int    `json:"required_evidence_count"`
	MissingGateCount             int    `json:"missing_gate_count"`
	MissingRequiredEvidenceCount int    `json:"missing_required_evidence_count"`
	MergeReady                   bool   `json:"merge_ready"`
	MergeBlockedUntilEvidence    bool   `json:"merge_blocked_until_evidence"`
}
