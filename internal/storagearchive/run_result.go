package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func skippedResult(
	source domain.PrivateLogSource,
	state string,
	evidence configEvidenceRef,
) RunResult {
	return RunResult{
		SourceKey:            source.Key,
		SourcePath:           source.Path,
		State:                state,
		ConfigEvidenceSHA256: evidence.SHA256,
		ConfigEvidenceField:  evidence.Field,
		BudgetOK:             true,
	}
}

func scannedResult(
	source domain.PrivateLogSource,
	state string,
	scan sourceScan,
	evidence configEvidenceRef,
) RunResult {
	return RunResult{
		SourceKey:            source.Key,
		SourcePath:           source.Path,
		State:                state,
		InputBytes:           int64(len(scan.Content)),
		InputSHA256:          scan.InputSHA256,
		ConfigEvidenceSHA256: evidence.SHA256,
		ConfigEvidenceField:  evidence.Field,
		RecordCount:          scan.RecordCount,
		NoiseCount:           scan.NoiseCount,
		NoiseRatioPercent:    scan.NoiseRatioPercent,
		BudgetOK:             scan.BudgetOK,
	}
}

func archivedResult(
	source domain.PrivateLogSource,
	scan sourceScan,
	archivePath string,
	outputBytes int64,
	evidence configEvidenceRef,
) RunResult {
	result := scannedResult(source, "archived", scan, evidence)
	result.ArchivePath = archivePath
	result.OutputBytes = outputBytes
	result.CompressionRatioPercent = compressionRatioPercent(result.InputBytes, outputBytes)
	return result
}
