package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func newRunReport(policy domain.StoragePolicy, checkedAt string) RunReport {
	return RunReport{
		PolicyPath:              PolicyRelativePath,
		ArchiveRoot:             policy.LogArchive.ArchiveRoot,
		ManifestPath:            policy.LogArchive.ManifestPath,
		SourceCount:             len(policy.PrivateLogSources),
		PublicSafe:              true,
		RawPayloadPublicAllowed: policy.LogArchive.RawPayloadPublicAllowed,
		CheckedAt:               checkedAt,
	}
}

func applyResult(report *RunReport, result RunResult) {
	switch result.State {
	case "archived":
		report.ArchivedCount++
	case "budget_breach":
		report.BudgetBreachCount++
		report.SkippedCount++
	default:
		report.SkippedCount++
	}
}
