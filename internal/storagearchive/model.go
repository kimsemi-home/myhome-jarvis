package storagearchive

type Status struct {
	PolicyPath                   string   `json:"policy_path"`
	Compression                  string   `json:"compression"`
	ArchiveRoot                  string   `json:"archive_root"`
	ArchiveExtension             string   `json:"archive_extension"`
	ManifestPath                 string   `json:"manifest_path"`
	PrivateLogSourceCount        int      `json:"private_log_source_count"`
	PrivateLogSourceKeys         []string `json:"private_log_source_keys"`
	Lifecycle                    []string `json:"lifecycle"`
	NoiseBudgetEnabled           bool     `json:"noise_budget_enabled"`
	MaxNoiseRatioPercent         int      `json:"max_noise_ratio_percent"`
	MaxLowSignalRecordsPerWindow int      `json:"max_low_signal_records_per_window"`
	NoiseBudgetWindow            string   `json:"noise_budget_window"`
	DedupeKeyFields              []string `json:"dedupe_key_fields"`
	ConfigEvidenceField          string   `json:"config_evidence_field"`
	ConfigHashInputs             []string `json:"config_hash_inputs"`
	ConfigEvidenceSHA256         string   `json:"config_evidence_sha256"`
	BreachBlocksArchive          bool     `json:"breach_blocks_archive"`
	ConfigIsEvidence             bool     `json:"config_is_evidence"`
	RawPayloadPublicAllowed      bool     `json:"raw_payload_public_allowed"`
	PublicSafe                   bool     `json:"public_safe"`
	ArchiveReady                 bool     `json:"archive_ready"`
	NoiseBudgetReady             bool     `json:"noise_budget_ready"`
	CompressionArchivePattern    string   `json:"compression_archive_pattern"`
	ManifestPresent              bool     `json:"manifest_present"`
	ManifestEntryCount           int      `json:"manifest_entry_count"`
	ManifestArchivedCount        int      `json:"manifest_archived_count"`
	ManifestSkippedCount         int      `json:"manifest_skipped_count"`
	ManifestBudgetBreachCount    int      `json:"manifest_budget_breach_count"`
	ManifestInvalidEntryCount    int      `json:"manifest_invalid_entry_count"`
	ManifestArchivedInputBytes   int64    `json:"manifest_archived_input_bytes"`
	ManifestArchivedOutputBytes  int64    `json:"manifest_archived_output_bytes"`
	ManifestCompressionRatio     int      `json:"manifest_compression_ratio_percent"`
	ManifestLastEntryAt          string   `json:"manifest_last_entry_at,omitempty"`
	ManifestLastArchivedAt       string   `json:"manifest_last_archived_at,omitempty"`
	ManifestLastBudgetBreachAt   string   `json:"manifest_last_budget_breach_at,omitempty"`
	CheckedAt                    string   `json:"checked_at"`
}
