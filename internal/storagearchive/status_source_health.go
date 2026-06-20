package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func withSourceHealth(
	status Status,
	policy domain.StoragePolicy,
	manifest manifestSummary,
) Status {
	status.SourceHealth = make([]SourceHealth, 0, len(policy.PrivateLogSources))
	for _, source := range policy.PrivateLogSources {
		if source.Key == "" {
			continue
		}
		health := sourceHealthForKey(source.Key, manifest)
		status.SourceHealth = append(status.SourceHealth, health)
		if health.HealthDebt {
			status.SourceHealthDebtCount++
		}
	}
	return status
}

func sourceHealthForKey(key string, manifest manifestSummary) SourceHealth {
	entry, ok := manifest.LatestBySource[key]
	if !ok {
		return SourceHealth{
			SourceKey:           key,
			LatestState:         "missing_manifest_evidence",
			LatestBudgetVerdict: "unknown",
			HealthDebt:          true,
		}
	}
	archived := manifest.ArchivedBySource[key]
	return sourceHealthFromEntries(key, entry, archived)
}

func sourceHealthFromEntries(
	key string,
	entry manifestEntry,
	archived manifestEntry,
) SourceHealth {
	archivePresent := archived.State == "archived" && archived.ArchivePath != ""
	hashKeyPresent := archived.InputSHA256 != "" && archived.ConfigEvidenceSHA256 != ""
	verdict := knownOrUnknown(entry.BudgetVerdict)
	state := knownOrUnknown(entry.State)
	return SourceHealth{
		SourceKey:               key,
		LatestState:             state,
		LatestBudgetVerdict:     verdict,
		RecordCount:             entry.RecordCount,
		NoiseCount:              entry.NoiseCount,
		NoiseRatioPercent:       entry.NoiseRatioPercent,
		CompressionRatioPercent: sourceCompressionRatio(entry, archived),
		ArchiveEvidencePresent:  archivePresent,
		HashCacheKeyPresent:     hashKeyPresent,
		BudgetOK:                verdict == "ok",
		HealthDebt:              verdict != "ok" || !archivePresent,
		LatestObservedAt:        entry.At,
		LatestArchivedAt:        archivedAt(archivePresent, archived),
	}
}
