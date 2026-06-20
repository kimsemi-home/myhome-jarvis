package codexcost

type ROISummary struct {
	PolicyPath                   string   `json:"policy_path"`
	LedgerPath                   string   `json:"ledger_path"`
	AttributionLedgerPath        string   `json:"attribution_ledger_path"`
	ScopeCount                   int      `json:"scope_count"`
	TrackedScopeCount            int      `json:"tracked_scope_count"`
	TotalUnits                   int64    `json:"total_units"`
	AttributedUnits              int64    `json:"attributed_units"`
	AttributionEntryUnits        int64    `json:"attribution_entry_units"`
	AttributionRecordCount       int      `json:"attribution_record_count"`
	AttributionCostRefCount      int      `json:"attribution_cost_ref_count"`
	AttributionCoveragePercent   int      `json:"attribution_coverage_percent"`
	InvalidAttributionCount      int      `json:"invalid_attribution_count"`
	BudgetState                  string   `json:"budget_state"`
	SustainabilityPosture        string   `json:"sustainability_posture"`
	TrendPosture                 string   `json:"trend_posture"`
	ReviewGateCount              int      `json:"review_gate_count"`
	AcceptedChangeCount          int64    `json:"accepted_change_count"`
	LedgerAcceptedChangeCount    int64    `json:"ledger_accepted_change_count"`
	MergeAcceptedChangeCount     int64    `json:"merge_accepted_change_count"`
	AcceptedChangeEvidenceSource string   `json:"accepted_change_evidence_source"`
	AcceptedChangeLogLimit       int      `json:"accepted_change_log_limit"`
	CacheSavingsUnits            int64    `json:"cache_savings_units"`
	ValueProxyUnits              int64    `json:"value_proxy_units"`
	CostPerAcceptedChange        int64    `json:"cost_per_accepted_change"`
	ValueProxyMethod             string   `json:"value_proxy_method"`
	StorageArchivePattern        string   `json:"storage_archive_pattern"`
	StorageArchiveReady          bool     `json:"storage_archive_ready"`
	NoiseBudgetReady             bool     `json:"noise_budget_ready"`
	MaxNoiseRatioPercent         int      `json:"max_noise_ratio_percent"`
	ConfigEvidenceField          string   `json:"config_evidence_field"`
	ConfigIsEvidence             bool     `json:"config_is_evidence"`
	PrivateLogSourceKeys         []string `json:"private_log_source_keys"`
	ForbiddenPublicFieldCount    int      `json:"forbidden_public_field_count"`
	Rows                         []ROIRow `json:"rows"`
	CheckedAt                    string   `json:"checked_at"`
}

type ROIRow struct {
	Scope                    string `json:"scope"`
	CostUnits                int64  `json:"cost_units"`
	AttributedCostUnits      int64  `json:"attributed_cost_units"`
	AttributionSubjectCount  int    `json:"attribution_subject_count"`
	CostSharePercent         int    `json:"cost_share_percent"`
	Status                   string `json:"status"`
	ValueProxyUnits          int64  `json:"value_proxy_units"`
	CostPerAcceptedChange    int64  `json:"cost_per_accepted_change"`
	BudgetState              string `json:"budget_state"`
	SustainabilityPosture    string `json:"sustainability_posture"`
	ReviewGateCount          int    `json:"review_gate_count"`
	StorageArchivePattern    string `json:"storage_archive_pattern"`
	NoiseBudgetReady         bool   `json:"noise_budget_ready"`
	EvidenceConfigIsEvidence bool   `json:"evidence_config_is_evidence"`
	Recommendation           string `json:"recommendation"`
}
