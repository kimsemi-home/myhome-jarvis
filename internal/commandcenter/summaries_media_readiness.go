package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/mediareadiness"

func summarizeMediaReadiness(status mediareadiness.Status) MediaReadinessSummary {
	return MediaReadinessSummary{
		PublicSafe:              status.PublicSafe,
		CaseCount:               status.CaseCount,
		AvailableCount:          status.AvailableCount,
		DegradedCount:           status.DegradedCount,
		MaxPlanningLatencyMS:    status.MaxPlanningLatencyMS,
		TargetPlanningLatencyMS: status.TargetPlanningLatencyMS,
		LocalLauncherAvailable:  status.LocalLauncherAvailable,
	}
}
