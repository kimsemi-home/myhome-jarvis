package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"

func storageArchiveDebtCount(status storagearchive.Status) int {
	count := 0
	if !status.PublicSafe {
		count++
	}
	if !status.ArchiveReady || !status.NoiseBudgetReady {
		count++
	}
	count += status.ManifestInvalidEntryCount
	count += status.ManifestBudgetBreachCount
	return count
}
