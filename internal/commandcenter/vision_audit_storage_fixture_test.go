package commandcenter

func readyStorageArchiveSummary() StorageArchiveSummary {
	return StorageArchiveSummary{
		PublicSafe:                true,
		CompressionArchivePattern: "compress_then_archive",
		Compression:               "gzip",
		PrivateLogSourceCount:     10,
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
		ConfigEvidenceSHA256: "ed77dd74e96a24249b9aa17dc69cef044053c354d1bb62575be25a0835a2d3c9",
		ConfigIsEvidence:     true,
		BreachBlocksArchive:  true,
	}
}
