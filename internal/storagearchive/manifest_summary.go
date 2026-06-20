package storagearchive

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func readManifestSummary(
	root string,
	policy domain.StoragePolicy,
) (manifestSummary, error) {
	path := filepath.Join(root, filepath.FromSlash(policy.LogArchive.ManifestPath))
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return manifestSummary{}, nil
	}
	if err != nil {
		return manifestSummary{}, err
	}
	defer file.Close()
	return scanManifest(file)
}

func scanManifest(file *os.File) (manifestSummary, error) {
	summary := manifestSummary{
		Present:          true,
		LatestBySource:   map[string]manifestEntry{},
		ArchivedBySource: map[string]manifestEntry{},
	}
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		applyManifestLine(&summary, line)
	}
	if err := scanner.Err(); err != nil {
		return manifestSummary{}, err
	}
	summary.CompressionRatio = compressionRatioPercent(
		summary.ArchivedInputBytes,
		summary.ArchivedOutputBytes,
	)
	return summary, nil
}
