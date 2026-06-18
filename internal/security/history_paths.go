package security

import (
	"path/filepath"
	"strings"
)

func checkHistoryPaths(root string, report *HistoryReport) error {
	lines, err := gitLines(root, "log", "--all", "--name-only", "--pretty=format:__MHJ_COMMIT__%H")
	if err != nil {
		return err
	}
	commit := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "__MHJ_COMMIT__") {
			commit = strings.TrimPrefix(line, "__MHJ_COMMIT__")
			continue
		}
		rel := filepath.ToSlash(strings.TrimSpace(line))
		if rel != "" {
			checkHistoryPath(commit, rel, report)
		}
	}
	return nil
}

func checkHistoryPath(commit string, rel string, report *HistoryReport) {
	base := strings.ToLower(filepath.Base(rel))
	ext := strings.ToLower(filepath.Ext(rel))
	checkHistoryEnvPath(commit, rel, report)
	checkHistoryPrivatePath(commit, rel, report)
	checkHistoryLanguagePath(commit, rel, ext, report)
	checkHistoryDependencyPath(commit, rel, base, report)
	checkHistorySensitivePath(commit, rel, base, report)
}
