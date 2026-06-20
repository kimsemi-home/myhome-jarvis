package commandcenter

func visionEvidenceRetention(summary StorageArchiveSummary) VisionEvidenceRetention {
	return VisionEvidenceRetention{
		PublicSafe:                summary.PublicSafe,
		CompressionArchivePattern: summary.CompressionArchivePattern,
		Compression:               summary.Compression,
		PrivateLogSourceCount:     summary.PrivateLogSourceCount,
		ArchiveReady:              summary.ArchiveReady,
		NoiseBudgetReady:          summary.NoiseBudgetReady,
		MaxNoiseRatioPercent:      summary.MaxNoiseRatioPercent,
		MaxLowSignalRecords:       summary.MaxLowSignalRecords,
		NoiseBudgetWindow:         summary.NoiseBudgetWindow,
		DedupeKeyFields:           append([]string{}, summary.DedupeKeyFields...),
		ConfigEvidenceField:       summary.ConfigEvidenceField,
		ConfigHashInputs:          append([]string{}, summary.ConfigHashInputs...),
		ConfigEvidenceSHA256:      summary.ConfigEvidenceSHA256,
		ConfigIsEvidence:          summary.ConfigIsEvidence,
		BreachBlocksArchive:       summary.BreachBlocksArchive,
		ManifestPresent:           summary.ManifestPresent,
		ManifestEntryCount:        summary.ManifestEntryCount,
		ManifestArchivedCount:     summary.ManifestArchivedCount,
		ManifestBudgetBreachCount: summary.ManifestBudgetBreachCount,
		ManifestInvalidEntryCount: summary.ManifestInvalidEntryCount,
		ManifestCompressionRatio:  summary.ManifestCompressionRatio,
	}
}
