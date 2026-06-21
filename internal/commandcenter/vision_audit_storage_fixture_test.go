package commandcenter

func readyStorageArchiveSummary() StorageArchiveSummary {
	return StorageArchiveSummary{
		PublicSafe:                true,
		CompressionArchivePattern: "compress_then_archive",
		Compression:               "gzip",
		PrivateLogSourceCount:     11,
		ArchiveReady:              true,
		NoiseBudgetReady:          true,
		MaxNoiseRatioPercent:      20,
		MaxLowSignalRecords:       10,
		NoiseBudgetWindow:         "per_quality_run",
		DedupeKeyFields: []string{
			"source", "kind", "evidence_ref",
		},
		ConfigEvidenceField: "evidence_noise_budget",
		ConfigHashInputs: []string{
			"private_log_sources", "log_archive", "evidence_noise_budget",
		},
		ConfigEvidenceSHA256: "5b92806ec9bc649cc6a4b0988ba98e2de751d53abfa99d7bb2eb1967851e79af",
		ConfigIsEvidence:     true,
		BreachBlocksArchive:  true,
	}
}
