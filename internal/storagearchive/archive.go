package storagearchive

import (
	"path"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func archiveSource(
	root string,
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	scan sourceScan,
	now string,
) (RunResult, manifestEntry, error) {
	archivePath := archiveRelativePath(policy, source, scan)
	outputBytes, err := writeGzip(root, archivePath, scan.Content)
	if err != nil {
		return RunResult{}, manifestEntry{}, err
	}
	result := archivedResult(source, scan, archivePath, outputBytes)
	return result, scannedEntry(now, result), nil
}

func scannedSkip(
	now string,
	source domain.PrivateLogSource,
	state string,
	scan sourceScan,
) (RunResult, manifestEntry) {
	result := scannedResult(source, state, scan)
	return result, scannedEntry(now, result)
}

func archiveRelativePath(
	policy domain.StoragePolicy,
	source domain.PrivateLogSource,
	scan sourceScan,
) string {
	name := safeName(source.Key)
	if len(scan.InputSHA256) >= 16 {
		name += "-" + scan.InputSHA256[:16]
	}
	return path.Join(policy.LogArchive.ArchiveRoot, name+policy.LogArchive.ArchiveExtension)
}

func safeName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "source"
	}
	return strings.Map(safeNameRune, name)
}

func safeNameRune(char rune) rune {
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' ||
		char >= '0' && char <= '9' || char == '-' || char == '_' {
		return char
	}
	return '-'
}
