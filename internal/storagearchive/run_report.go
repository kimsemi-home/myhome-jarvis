package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func newRunReport(
	policy domain.StoragePolicy,
	checkedAt string,
	evidence configEvidenceRef,
) RunReport {
	return RunReport{
		PolicyPath:              PolicyRelativePath,
		ArchiveRoot:             policy.LogArchive.ArchiveRoot,
		ManifestPath:            policy.LogArchive.ManifestPath,
		SourceCount:             len(policy.PrivateLogSources),
		ConfigEvidenceField:     evidence.Field,
		ConfigHashInputs:        append([]string{}, evidence.Inputs...),
		ConfigEvidenceSHA256:    evidence.SHA256,
		PublicSafe:              true,
		RawPayloadPublicAllowed: policy.LogArchive.RawPayloadPublicAllowed,
		CheckedAt:               checkedAt,
	}
}

func applyResult(report *RunReport, result RunResult) {
	switch result.State {
	case "archived":
		report.ArchivedCount++
		report.ArchivedInputBytes += result.InputBytes
		report.ArchivedOutputBytes += result.OutputBytes
		report.CompressionRatioPercent = compressionRatioPercent(
			report.ArchivedInputBytes,
			report.ArchivedOutputBytes,
		)
	case "cached":
		report.CachedCount++
	case "budget_breach":
		report.BudgetBreachCount++
		report.SkippedCount++
	default:
		report.SkippedCount++
	}
}
