package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func skippedResult(source domain.PrivateLogSource, state string) RunResult {
	return RunResult{
		SourceKey:  source.Key,
		SourcePath: source.Path,
		State:      state,
		BudgetOK:   true,
	}
}

func scannedResult(source domain.PrivateLogSource, state string, scan sourceScan) RunResult {
	return RunResult{
		SourceKey:         source.Key,
		SourcePath:        source.Path,
		State:             state,
		InputBytes:        int64(len(scan.Content)),
		InputSHA256:       scan.InputSHA256,
		RecordCount:       scan.RecordCount,
		NoiseCount:        scan.NoiseCount,
		NoiseRatioPercent: scan.NoiseRatioPercent,
		BudgetOK:          scan.BudgetOK,
	}
}

func archivedResult(
	source domain.PrivateLogSource,
	scan sourceScan,
	archivePath string,
	outputBytes int64,
) RunResult {
	result := scannedResult(source, "archived", scan)
	result.ArchivePath = archivePath
	result.OutputBytes = outputBytes
	return result
}
