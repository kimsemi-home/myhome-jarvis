package storagearchive

type SourceHealth struct {
	SourceKey               string `json:"source_key"`
	LatestState             string `json:"latest_state"`
	LatestBudgetVerdict     string `json:"latest_budget_verdict"`
	RecordCount             int    `json:"record_count"`
	NoiseCount              int    `json:"noise_count"`
	NoiseRatioPercent       int    `json:"noise_ratio_percent"`
	CompressionRatioPercent int    `json:"compression_ratio_percent"`
	ArchiveEvidencePresent  bool   `json:"archive_evidence_present"`
	HashCacheKeyPresent     bool   `json:"hash_cache_key_present"`
	BudgetOK                bool   `json:"budget_ok"`
	HealthDebt              bool   `json:"health_debt"`
	LatestObservedAt        string `json:"latest_observed_at,omitempty"`
	LatestArchivedAt        string `json:"latest_archived_at,omitempty"`
}
