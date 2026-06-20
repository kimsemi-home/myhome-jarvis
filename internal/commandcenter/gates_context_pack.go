package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/contextpack"

func contextPackDebtCount(status contextpack.Status) int {
	count := 0
	if status.SplitCriteriaCount < 5 {
		count++
	}
	if status.ExportedArtifactCount < 6 {
		count++
	}
	if !status.PublicSafe {
		count++
	}
	return count
}
