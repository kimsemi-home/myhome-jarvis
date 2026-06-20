package storagearchive

func withManifestSummary(
	status Status,
	manifest manifestSummary,
) Status {
	status.ManifestPresent = manifest.Present
	status.ManifestEntryCount = manifest.EntryCount
	status.ManifestArchivedCount = manifest.ArchivedCount
	status.ManifestSkippedCount = manifest.SkippedCount
	status.ManifestBudgetBreachCount = manifest.BudgetBreachCount
	status.ManifestInvalidEntryCount = manifest.InvalidEntryCount
	status.ManifestArchivedInputBytes = manifest.ArchivedInputBytes
	status.ManifestArchivedOutputBytes = manifest.ArchivedOutputBytes
	status.ManifestCompressionRatio = manifest.CompressionRatio
	status.ManifestLastEntryAt = manifest.LastEntryAt
	status.ManifestLastArchivedAt = manifest.LastArchivedAt
	status.ManifestLastBudgetBreachAt = manifest.LastBudgetBreachAt
	return status
}
