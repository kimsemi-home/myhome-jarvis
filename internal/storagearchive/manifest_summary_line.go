package storagearchive

import "encoding/json"

func applyManifestLine(summary *manifestSummary, line string) {
	var entry manifestEntry
	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		summary.InvalidEntryCount++
		return
	}
	summary.EntryCount++
	summary.LastEntryAt = latestRFC3339(summary.LastEntryAt, entry.At)
	recordSourceManifestEntry(summary, entry)
	if entry.State == "archived" {
		applyArchivedManifestEntry(summary, entry)
		return
	}
	summary.SkippedCount++
	if entry.BudgetVerdict == "breach" || entry.State == "budget_breach" {
		summary.BudgetBreachCount++
		summary.LastBudgetBreachAt = latestRFC3339(
			summary.LastBudgetBreachAt,
			entry.At,
		)
	}
}

func applyArchivedManifestEntry(summary *manifestSummary, entry manifestEntry) {
	summary.ArchivedCount++
	summary.ArchivedInputBytes += entry.InputBytes
	summary.ArchivedOutputBytes += entry.OutputBytes
	summary.LastArchivedAt = latestRFC3339(summary.LastArchivedAt, entry.At)
}

func latestRFC3339(current string, candidate string) string {
	if candidate > current {
		return candidate
	}
	return current
}
