package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"

func summarizeStorageArchiveSources(
	sources []storagearchive.SourceHealth,
) []StorageArchiveSourceHealth {
	summary := make([]StorageArchiveSourceHealth, 0, len(sources))
	for _, source := range sources {
		summary = append(summary, StorageArchiveSourceHealth{
			SourceKey:               source.SourceKey,
			LatestState:             source.LatestState,
			LatestBudgetVerdict:     source.LatestBudgetVerdict,
			RecordCount:             source.RecordCount,
			NoiseCount:              source.NoiseCount,
			NoiseRatioPercent:       source.NoiseRatioPercent,
			CompressionRatioPercent: source.CompressionRatioPercent,
			ArchiveEvidencePresent:  source.ArchiveEvidencePresent,
			HashCacheKeyPresent:     source.HashCacheKeyPresent,
			BudgetOK:                source.BudgetOK,
			HealthDebt:              source.HealthDebt,
			LatestObservedAt:        source.LatestObservedAt,
			LatestArchivedAt:        source.LatestArchivedAt,
		})
	}
	return summary
}
