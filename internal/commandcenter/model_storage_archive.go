package commandcenter

type StorageArchiveSummary struct {
	PublicSafe                bool     `json:"public_safe"`
	CompressionArchivePattern string   `json:"compression_archive_pattern"`
	Compression               string   `json:"compression"`
	PrivateLogSourceCount     int      `json:"private_log_source_count"`
	PrivateLogSourceKeys      []string `json:"private_log_source_keys"`
	ArchiveReady              bool     `json:"archive_ready"`
	NoiseBudgetReady          bool     `json:"noise_budget_ready"`
	MaxNoiseRatioPercent      int      `json:"max_noise_ratio_percent"`
	MaxLowSignalRecords       int      `json:"max_low_signal_records_per_window"`
	NoiseBudgetWindow         string   `json:"noise_budget_window"`
	DedupeKeyFields           []string `json:"dedupe_key_fields"`
	ConfigEvidenceField       string   `json:"config_evidence_field"`
	ConfigHashInputs          []string `json:"config_hash_inputs"`
	ConfigEvidenceSHA256      string   `json:"config_evidence_sha256"`
	ConfigIsEvidence          bool     `json:"config_is_evidence"`
	BreachBlocksArchive       bool     `json:"breach_blocks_archive"`
	ManifestPresent           bool     `json:"manifest_present"`
	ManifestEntryCount        int      `json:"manifest_entry_count"`
	ManifestArchivedCount     int      `json:"manifest_archived_count"`
	ManifestSkippedCount      int      `json:"manifest_skipped_count"`
	ManifestBudgetBreachCount int      `json:"manifest_budget_breach_count"`
	ManifestInvalidEntryCount int      `json:"manifest_invalid_entry_count"`
	ManifestCompressionRatio  int      `json:"manifest_compression_ratio_percent"`
	ManifestLastArchivedAt    string   `json:"manifest_last_archived_at,omitempty"`
}
