package commandcenter

type VisionAudit struct {
	Context                 string                   `json:"context"`
	Version                 string                   `json:"version"`
	PublicSafe              bool                     `json:"public_safe"`
	Redaction               string                   `json:"redaction"`
	PolicyPath              string                   `json:"policy_path"`
	Mission                 string                   `json:"mission"`
	OperatingMode           string                   `json:"operating_mode"`
	UniversalTermCount      int                      `json:"universal_term_count"`
	LinearEpicCount         int                      `json:"linear_epic_count"`
	RequirementCount        int                      `json:"requirement_count"`
	ReadyRequirementCount   int                      `json:"ready_requirement_count"`
	GatedRequirementCount   int                      `json:"gated_requirement_count"`
	BlockedRequirementCount int                      `json:"blocked_requirement_count"`
	OpenGateCount           int                      `json:"open_gate_count"`
	GoalComplete            bool                     `json:"goal_complete"`
	CompletionRule          string                   `json:"completion_rule"`
	NextSafeAction          string                   `json:"next_safe_action"`
	EvidenceRetention       VisionEvidenceRetention  `json:"evidence_retention"`
	Requirements            []VisionRequirementAudit `json:"requirements"`
	CheckedAt               string                   `json:"checked_at"`
}

type VisionRequirementAudit struct {
	CapabilityKey  string   `json:"capability_key"`
	State          string   `json:"state"`
	EvidenceRefs   []string `json:"evidence_refs"`
	GateRefs       []string `json:"gate_refs"`
	NextSafeAction string   `json:"next_safe_action"`
}

type VisionEvidenceRetention struct {
	PublicSafe                bool     `json:"public_safe"`
	CompressionArchivePattern string   `json:"compression_archive_pattern"`
	Compression               string   `json:"compression"`
	PrivateLogSourceCount     int      `json:"private_log_source_count"`
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
	ManifestBudgetBreachCount int      `json:"manifest_budget_breach_count"`
	ManifestInvalidEntryCount int      `json:"manifest_invalid_entry_count"`
	ManifestCompressionRatio  int      `json:"manifest_compression_ratio_percent"`
}
