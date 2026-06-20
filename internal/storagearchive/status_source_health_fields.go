package storagearchive

func sourceCompressionRatio(entry manifestEntry, archived manifestEntry) int {
	if entry.CompressionRatioPercent > 0 {
		return entry.CompressionRatioPercent
	}
	return archived.CompressionRatioPercent
}

func knownOrUnknown(value string) string {
	if value == "" {
		return "unknown"
	}
	return value
}

func archivedAt(present bool, archived manifestEntry) string {
	if !present {
		return ""
	}
	return archived.At
}
