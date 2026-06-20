package storagearchive

import (
	"os"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func cachedArchiveResult(
	root string,
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	scan sourceScan,
	evidence configEvidenceRef,
	cache archiveCache,
) (RunResult, bool, error) {
	archivePath := archiveRelativePath(policy, source, scan)
	entry, ok := cache.hit(source.Key, scan.InputSHA256, evidence.SHA256)
	if !ok || entry.ArchivePath != archivePath {
		return RunResult{}, false, nil
	}
	info, err := os.Stat(filepath.Join(root, filepath.FromSlash(archivePath)))
	if os.IsNotExist(err) {
		return RunResult{}, false, nil
	}
	if err != nil {
		return RunResult{}, false, err
	}
	return cachedResult(source, scan, archivePath, info.Size(), evidence), true, nil
}
