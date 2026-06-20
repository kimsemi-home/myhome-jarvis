package storagearchive

import (
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
	evidence := configEvidenceRefForPolicy(policy)
	report := newRunReport(policy, now, evidence)
	cache, err := readArchiveCache(root, policy.LogArchive.ManifestPath)
	if err != nil {
		return RunReport{}, err
	}
	entries := make([]manifestEntry, 0, len(policy.PrivateLogSources))
	for _, source := range policy.PrivateLogSources {
		result, entry, err := runSource(root, policy, source, now, evidence, cache)
		if err != nil {
			return RunReport{}, err
		}
		report.Results = append(report.Results, result)
		if entry.At != "" {
			entries = append(entries, entry)
		}
		applyResult(&report, result)
	}
	if err := appendManifest(root, policy.LogArchive.ManifestPath, entries); err != nil {
		return RunReport{}, err
	}
	return report, nil
}
