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
		ConfigEvidenceField:       status.ConfigEvidenceField,
		ConfigEvidenceSHA256:      status.ConfigEvidenceSHA256,
		ConfigIsEvidence:          status.ConfigIsEvidence,
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
