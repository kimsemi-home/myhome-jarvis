package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

type RunPolicyEvidence struct {
	CompressionArchivePattern    string   `json:"compression_archive_pattern"`
	Compression                  string   `json:"compression"`
	NoiseBudgetEnabled           bool     `json:"noise_budget_enabled"`
	MaxNoiseRatioPercent         int      `json:"max_noise_ratio_percent"`
	MaxLowSignalRecordsPerWindow int      `json:"max_low_signal_records_per_window"`
	NoiseBudgetWindow            string   `json:"noise_budget_window"`
	DedupeKeyFields              []string `json:"dedupe_key_fields"`
	ConfigEvidenceField          string   `json:"config_evidence_field"`
	ConfigHashInputs             []string `json:"config_hash_inputs"`
	ConfigEvidenceSHA256         string   `json:"config_evidence_sha256"`
	ConfigIsEvidence             bool     `json:"config_is_evidence"`
	BreachBlocksArchive          bool     `json:"breach_blocks_archive"`
}

func newRunPolicyEvidence(
	policy domain.StoragePolicy,
	evidence configEvidenceRef,
) RunPolicyEvidence {
	return RunPolicyEvidence{
		CompressionArchivePattern:    policy.LogArchive.Mode,
		Compression:                  policy.LogArchive.Compression,
		NoiseBudgetEnabled:           policy.EvidenceNoiseBudget.Enabled,
		MaxNoiseRatioPercent:         policy.EvidenceNoiseBudget.MaxNoiseRatioPercent,
		MaxLowSignalRecordsPerWindow: policy.EvidenceNoiseBudget.MaxLowSignalRecordsPerWindow,
		NoiseBudgetWindow:            policy.EvidenceNoiseBudget.Window,
		DedupeKeyFields:              append([]string{}, policy.EvidenceNoiseBudget.DedupeKeyFields...),
		ConfigEvidenceField:          evidence.Field,
		ConfigHashInputs:             append([]string{}, evidence.Inputs...),
		ConfigEvidenceSHA256:         evidence.SHA256,
		ConfigIsEvidence:             policy.LogArchive.ConfigIsEvidence,
		BreachBlocksArchive:          policy.EvidenceNoiseBudget.BreachBlocksArchive,
	}
}
