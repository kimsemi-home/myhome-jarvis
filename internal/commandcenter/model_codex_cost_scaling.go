package commandcenter

type CodexCostScalingSummary struct {
	PublicSafe                    bool     `json:"public_safe"`
	Decision                      string   `json:"decision"`
	Recommendation                string   `json:"recommendation"`
	NextSafeAction                string   `json:"next_safe_action"`
	CanApplyExpansion             bool     `json:"can_apply_expansion"`
	BudgetState                   string   `json:"budget_state"`
	TotalUnits                    int64    `json:"total_units"`
	RemainingToWarningUnits       int64    `json:"remaining_to_warning_units"`
	RemainingToReviewUnits        int64    `json:"remaining_to_review_units"`
	WarningUsedPercent            int      `json:"warning_used_percent"`
	ReviewUsedPercent             int      `json:"review_used_percent"`
	AttributionCoveragePercent    int      `json:"attribution_coverage_percent"`
	AcceptedChangeCount           int64    `json:"accepted_change_count"`
	CacheSavingsUnits             int64    `json:"cache_savings_units"`
	ValueProxyUnits               int64    `json:"value_proxy_units"`
	CostPerAcceptedChange         int64    `json:"cost_per_accepted_change"`
	SustainabilityPosture         string   `json:"sustainability_posture"`
	TrendPosture                  string   `json:"trend_posture"`
	ReviewGateCount               int      `json:"review_gate_count"`
	StorageArchivePattern         string   `json:"storage_archive_pattern"`
	StorageArchiveReady           bool     `json:"storage_archive_ready"`
	NoiseBudgetReady              bool     `json:"noise_budget_ready"`
	ArchiveManifestBudgetBreaches int      `json:"archive_manifest_budget_breaches"`
	ArchiveManifestInvalidEntries int      `json:"archive_manifest_invalid_entries"`
	ConfigIsEvidence              bool     `json:"config_is_evidence"`
	ScalingOptionCount            int      `json:"scaling_option_count"`
	RecommendedScalingOptionKeys  []string `json:"recommended_scaling_option_keys"`
	GrantingScalingOptionCount    int      `json:"granting_scaling_option_count"`
}
