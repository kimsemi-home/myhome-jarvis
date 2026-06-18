package security

import "strings"

func checkHistoryEnvPath(commit string, rel string, report *HistoryReport) {
	if rel == ".env" || (strings.HasPrefix(rel, ".env.") && rel != ".env.example") {
		report.addHistory(commit, rel, 0, "history_forbidden_env_file", "environment files must never appear in git history")
	}
}

func checkHistoryPrivatePath(commit string, rel string, report *HistoryReport) {
	if (strings.HasPrefix(rel, "data/private/") || strings.HasPrefix(rel, "data/lake/")) &&
		!isAllowedPrivatePlaceholder(rel) {
		report.addHistory(commit, rel, 0, "history_private_data_path", "private data and lake files must never appear in git history")
	}
	if hasPathSegment(rel, "secrets") {
		report.addHistory(commit, rel, 0, "history_forbidden_secret_dir", "secrets directories must never appear in git history")
	}
}
