package storagearchive

import (
	"os"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func runSource(
	root string,
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	now string,
	evidence configEvidenceRef,
	cache archiveCache,
) (RunResult, manifestEntry, error) {
	path := filepath.Join(root, filepath.FromSlash(source.Path))
	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return skippedResult(source, "missing", evidence),
			skippedEntry(now, source, "missing", evidence), nil
	}
	if err != nil {
		return RunResult{}, manifestEntry{}, err
	}
	scan, err := scanSource(content, policy.EvidenceNoiseBudget)
	if err != nil {
		return RunResult{}, manifestEntry{}, err
	}
	return runScannedSource(root, policy, source, now, evidence, cache, scan)
}

func runScannedSource(
	root string,
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	now string,
	evidence configEvidenceRef,
	cache archiveCache,
	scan sourceScan,
) (RunResult, manifestEntry, error) {
	if scan.RecordCount == 0 {
		result, entry := scannedSkip(now, source, "empty", scan, evidence)
		return result, entry, nil
	}
	if !scan.BudgetOK && policy.EvidenceNoiseBudget.BreachBlocksArchive {
		result, entry := scannedSkip(now, source, "budget_breach", scan, evidence)
		return result, entry, nil
	}
	result, ok, err := cachedArchiveResult(root, policy, source, scan, evidence, cache)
	if err != nil || ok {
		return result, manifestEntry{}, err
	}
	return archiveSource(root, policy, source, scan, now, evidence)
}
