package storagearchive

import (
	"os"
	"path/filepath"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func RunForRoot(root string) (RunReport, error) {
	policy, err := domain.ReadStoragePolicy(
		filepath.Join(root, filepath.FromSlash(PolicyRelativePath)),
	)
	if err != nil {
		return RunReport{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return RunReport{}, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	report := newRunReport(policy, now)
	entries := make([]manifestEntry, 0, len(policy.PrivateLogSources))
	for _, source := range policy.PrivateLogSources {
		result, entry, err := runSource(root, policy, source, now)
		if err != nil {
			return RunReport{}, err
		}
		report.Results = append(report.Results, result)
		entries = append(entries, entry)
		applyResult(&report, result)
	}
	if err := appendManifest(root, policy.LogArchive.ManifestPath, entries); err != nil {
		return RunReport{}, err
	}
	return report, nil
}

func runSource(
	root string,
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	now string,
) (RunResult, manifestEntry, error) {
	path := filepath.Join(root, filepath.FromSlash(source.Path))
	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return skippedResult(source, "missing"), skippedEntry(now, source, "missing"), nil
	}
	if err != nil {
		return RunResult{}, manifestEntry{}, err
	}
	scan, err := scanSource(content, policy.EvidenceNoiseBudget)
	if err != nil {
		return RunResult{}, manifestEntry{}, err
	}
	if scan.RecordCount == 0 {
		result, entry := scannedSkip(now, source, "empty", scan)
		return result, entry, nil
	}
	if !scan.BudgetOK && policy.EvidenceNoiseBudget.BreachBlocksArchive {
		result, entry := scannedSkip(now, source, "budget_breach", scan)
		return result, entry, nil
	}
	return archiveSource(root, policy, source, scan, now)
}
