package codexcost

type Brief struct {
	PolicyPath                      string   `json:"policy_path"`
	PublicSafe                      bool     `json:"public_safe"`
	Decision                        string   `json:"decision"`
	Reasons                         []string `json:"reasons,omitempty"`
	Recommendation                  string   `json:"recommendation"`
	NextSafeAction                  string   `json:"next_safe_action"`
	BudgetState                     string   `json:"budget_state"`
	TotalUnits                      int64    `json:"total_units"`
	WarningUnitThreshold            int64    `json:"warning_unit_threshold"`
	ReviewUnitThreshold             int64    `json:"review_unit_threshold"`
	AttributionCoveragePercent      int      `json:"attribution_coverage_percent"`
	TrackedScopeCount               int      `json:"tracked_scope_count"`
	ScopeCount                      int      `json:"scope_count"`
	SustainabilityPosture           string   `json:"sustainability_posture"`
	TrendPosture                    string   `json:"trend_posture"`
	ReviewGateCount                 int      `json:"review_gate_count"`
	AcceptedChangeCount             int64    `json:"accepted_change_count"`
	CacheSavingsUnits               int64    `json:"cache_savings_units"`
	ValueProxyUnits                 int64    `json:"value_proxy_units"`
	CostPerAcceptedChange           int64    `json:"cost_per_accepted_change"`
	StorageArchivePattern           string   `json:"storage_archive_pattern"`
	StorageArchiveReady             bool     `json:"storage_archive_ready"`
	NoiseBudgetReady                bool     `json:"noise_budget_ready"`
	MaxNoiseRatioPercent            int      `json:"max_noise_ratio_percent"`
	ArchiveManifestEntryCount       int      `json:"archive_manifest_entry_count"`
	ArchiveManifestBudgetBreaches   int      `json:"archive_manifest_budget_breaches"`
	ArchiveManifestInvalidEntries   int      `json:"archive_manifest_invalid_entries"`
	ArchiveManifestCompressionRatio int      `json:"archive_manifest_compression_ratio_percent"`
	ConfigIsEvidence                bool     `json:"config_is_evidence"`
	ForbiddenPublicFieldCount       int      `json:"forbidden_public_field_count"`
	CheckedAt                       string   `json:"checked_at"`
}
