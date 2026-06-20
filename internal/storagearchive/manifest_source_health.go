package storagearchive

func recordSourceManifestEntry(summary *manifestSummary, entry manifestEntry) {
	if entry.SourceKey == "" {
		return
	}
	if entry.At >= summary.LatestBySource[entry.SourceKey].At {
		summary.LatestBySource[entry.SourceKey] = entry
	}
	if entry.State == "archived" &&
		entry.At >= summary.ArchivedBySource[entry.SourceKey].At {
		summary.ArchivedBySource[entry.SourceKey] = entry
	}
}
