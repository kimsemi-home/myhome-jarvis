package codexcost

type ScalingPacket struct {
	Context           string          `json:"context"`
	Version           string          `json:"version"`
	PublicSafe        bool            `json:"public_safe"`
	Redaction         string          `json:"redaction"`
	PolicyPath        string          `json:"policy_path"`
	Decision          string          `json:"decision"`
	Recommendation    string          `json:"recommendation"`
	NextSafeAction    string          `json:"next_safe_action"`
	BudgetHeadroom    BudgetHeadroom  `json:"budget_headroom"`
	EvidencePosture   ScalingEvidence `json:"evidence_posture"`
	StorageEvidence   ScalingStorage  `json:"storage_evidence"`
	ScalingOptions    []ScalingOption `json:"scaling_options"`
	CanApplyExpansion bool            `json:"can_apply_expansion"`
	CheckedAt         string          `json:"checked_at"`
}

type BudgetHeadroom struct {
	BudgetState          string `json:"budget_state"`
	TotalUnits           int64  `json:"total_units"`
	WarningUnitThreshold int64  `json:"warning_unit_threshold"`
	ReviewUnitThreshold  int64  `json:"review_unit_threshold"`
	RemainingToWarning   int64  `json:"remaining_to_warning_units"`
	RemainingToReview    int64  `json:"remaining_to_review_units"`
	WarningUsedPercent   int    `json:"warning_used_percent"`
	ReviewUsedPercent    int    `json:"review_used_percent"`
}

type ScalingEvidence struct {
	AttributionCoveragePercent int    `json:"attribution_coverage_percent"`
	AcceptedChangeCount        int64  `json:"accepted_change_count"`
	CacheSavingsUnits          int64  `json:"cache_savings_units"`
	ValueProxyUnits            int64  `json:"value_proxy_units"`
	CostPerAcceptedChange      int64  `json:"cost_per_accepted_change"`
	SustainabilityPosture      string `json:"sustainability_posture"`
	TrendPosture               string `json:"trend_posture"`
	ReviewGateCount            int    `json:"review_gate_count"`
}

type ScalingStorage struct {
	Pattern                  string `json:"compression_archive_pattern"`
	Ready                    bool   `json:"archive_ready"`
	NoiseBudgetReady         bool   `json:"noise_budget_ready"`
	MaxNoiseRatioPercent     int    `json:"max_noise_ratio_percent"`
	ManifestEntryCount       int    `json:"manifest_entry_count"`
	ManifestBudgetBreaches   int    `json:"manifest_budget_breach_count"`
	ManifestInvalidEntries   int    `json:"manifest_invalid_entry_count"`
	ManifestCompressionRatio int    `json:"manifest_compression_ratio_percent"`
	ConfigIsEvidence         bool   `json:"config_is_evidence"`
}

type ScalingOption struct {
	Key                           string `json:"key"`
	Label                         string `json:"label"`
	Recommended                   bool   `json:"recommended"`
	Effect                        string `json:"effect"`
	RequiresGuard                 bool   `json:"requires_guard"`
	RequiresHumanReview           bool   `json:"requires_human_review"`
	RequiresSeparateRecordCommand bool   `json:"requires_separate_record_command"`
	ThisPacketGrantsSpend         bool   `json:"this_packet_grants_spend"`
	AllowsPaidExpansion           bool   `json:"allows_paid_expansion"`
	AllowsExternalTooling         bool   `json:"allows_external_tooling"`
	AllowsWorkflowChanges         bool   `json:"allows_workflow_changes"`
	AllowsSelfApproval            bool   `json:"allows_self_approval"`
}
