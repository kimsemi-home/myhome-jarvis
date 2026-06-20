package domain

type StoragePolicy struct {
	FixtureFormat       string              `json:"fixture_format"`
	LakeLayers          []string            `json:"lake_layers"`
	Datasets            []string            `json:"datasets"`
	LongTermFormat      string              `json:"long_term_format"`
	Compression         string              `json:"compression"`
	PrivateRoot         string              `json:"private_root"`
	PrivateLogSources   []PrivateLogSource  `json:"private_log_sources"`
	LogArchive          LogArchivePolicy    `json:"log_archive"`
	EvidenceNoiseBudget EvidenceNoiseBudget `json:"evidence_noise_budget"`
}

type PrivateLogSource struct {
	Key    string `json:"key"`
	Path   string `json:"path"`
	Format string `json:"format"`
}

type LogArchivePolicy struct {
	Mode                    string   `json:"mode"`
	Compression             string   `json:"compression"`
	ArchiveRoot             string   `json:"archive_root"`
	ArchiveExtension        string   `json:"archive_extension"`
	ManifestPath            string   `json:"manifest_path"`
	RawPayloadPublicAllowed bool     `json:"raw_payload_public_allowed"`
	ConfigIsEvidence        bool     `json:"config_is_evidence"`
	Lifecycle               []string `json:"lifecycle"`
}

type EvidenceNoiseBudget struct {
	Enabled                      bool     `json:"enabled"`
	MaxNoiseRatioPercent         int      `json:"max_noise_ratio_percent"`
	MaxLowSignalRecordsPerWindow int      `json:"max_low_signal_records_per_window"`
	Window                       string   `json:"window"`
	DedupeKeyFields              []string `json:"dedupe_key_fields"`
	ConfigEvidenceField          string   `json:"config_evidence_field"`
	BreachBlocksArchive          bool     `json:"breach_blocks_archive"`
}

type RecommendationsSummary struct {
	Count int                  `json:"count"`
	Items []RecommendationItem `json:"items"`
}

type RecommendationItem struct {
	Kind                       string `json:"kind"`
	Title                      string `json:"title"`
	Rationale                  string `json:"rationale"`
	Score                      int    `json:"score"`
	Currency                   string `json:"currency"`
	EstimatedMonthlyMinorUnits int64  `json:"estimated_monthly_minor_units"`
	EvidenceCount              int    `json:"evidence_count"`
}

type HouseholdSummary struct {
	Scopes []HouseholdScopeSummary `json:"scopes"`
}

type HouseholdScopeSummary struct {
	Scope                   string `json:"scope"`
	Label                   string `json:"label"`
	Currency                string `json:"currency"`
	FinanceRecords          int    `json:"finance_records"`
	FinanceCreditMinorUnits int64  `json:"finance_credit_minor_units"`
	FinanceDebitMinorUnits  int64  `json:"finance_debit_minor_units"`
	FinanceNetMinorUnits    int64  `json:"finance_net_minor_units"`
	PurchaseRecords         int    `json:"purchase_records"`
	PurchaseSpendMinorUnits int64  `json:"purchase_spend_minor_units"`
}
