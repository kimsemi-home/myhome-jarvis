package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"

func summarizeStorageArchive(status storagearchive.Status) StorageArchiveSummary {
	return StorageArchiveSummary{
		PublicSafe:                status.PublicSafe,
		CompressionArchivePattern: status.CompressionArchivePattern,
		Compression:               status.Compression,
		PrivateLogSourceCount:     status.PrivateLogSourceCount,
		PrivateLogSourceKeys:      status.PrivateLogSourceKeys,
		ArchiveReady:              status.ArchiveReady,
		NoiseBudgetReady:          status.NoiseBudgetReady,
		MaxNoiseRatioPercent:      status.MaxNoiseRatioPercent,
		MaxLowSignalRecords:       status.MaxLowSignalRecordsPerWindow,
		NoiseBudgetWindow:         status.NoiseBudgetWindow,
		DedupeKeyFields:           append([]string{}, status.DedupeKeyFields...),
		ConfigEvidenceField:       status.ConfigEvidenceField,
		ConfigHashInputs:          append([]string{}, status.ConfigHashInputs...),
		ConfigEvidenceSHA256:      status.ConfigEvidenceSHA256,
		ConfigIsEvidence:          status.ConfigIsEvidence,
		BreachBlocksArchive:       status.BreachBlocksArchive,
		ManifestPresent:           status.ManifestPresent,
		ManifestEntryCount:        status.ManifestEntryCount,
		ManifestArchivedCount:     status.ManifestArchivedCount,
		ManifestSkippedCount:      status.ManifestSkippedCount,
		ManifestBudgetBreachCount: status.ManifestBudgetBreachCount,
		ManifestInvalidEntryCount: status.ManifestInvalidEntryCount,
		ManifestCompressionRatio:  status.ManifestCompressionRatio,
		ManifestLastArchivedAt:    status.ManifestLastArchivedAt,
	}
}
