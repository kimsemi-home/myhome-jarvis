package storagearchive

import "github.com/kimsemi-home/myhome-jarvis/internal/domain"

func skippedEntry(
	at string,
	source domain.PrivateLogSource,
	state string,
	evidence configEvidenceRef,
) manifestEntry {
	return manifestEntry{
		At:                   at,
		SourceKey:            source.Key,
		SourcePath:           source.Path,
		State:                state,
		ConfigEvidenceSHA256: evidence.SHA256,
		ConfigEvidenceField:  evidence.Field,
		BudgetVerdict:        "ok",
		RawPayloadStored:     false,
	}
}

func scannedEntry(at string, result RunResult) manifestEntry {
	return manifestEntry{
		At:                      at,
		SourceKey:               result.SourceKey,
		SourcePath:              result.SourcePath,
		ArchivePath:             result.ArchivePath,
		State:                   result.State,
		InputBytes:              result.InputBytes,
		OutputBytes:             result.OutputBytes,
		InputSHA256:             result.InputSHA256,
		ConfigEvidenceSHA256:    result.ConfigEvidenceSHA256,
		ConfigEvidenceField:     result.ConfigEvidenceField,
		RecordCount:             result.RecordCount,
		NoiseCount:              result.NoiseCount,
		NoiseRatioPercent:       result.NoiseRatioPercent,
		CompressionRatioPercent: result.CompressionRatioPercent,
		BudgetVerdict:           budgetVerdict(result),
		RawPayloadStored:        result.State == "archived",
	}
}

func budgetVerdict(result RunResult) string {
	if result.BudgetOK {
		return "ok"
	}
	return "breach"
}
