package security

func checkHistoryContent(root string, commits []string, pattern historyPattern, report *HistoryReport) error {
	const batchSize = 64
	for start := 0; start < len(commits); start += batchSize {
		end := start + batchSize
		if end > len(commits) {
			end = len(commits)
		}
		if err := checkHistoryContentBatch(root, commits[start:end], pattern, report); err != nil {
			return err
		}
	}
	return nil
}

func checkHistoryContentBatch(root string, commits []string, pattern historyPattern, report *HistoryReport) error {
	args := []string{"grep", "-n", "-I", "-E", "-e", pattern.Pattern}
	args = append(args, commits...)
	args = append(args, "--")
	lines, err := gitLinesAllowNoMatches(root, args...)
	if err != nil {
		return err
	}
	for _, line := range lines {
		commit, path, lineNumber, ok := parseGitGrepLine(line)
		if ok {
			report.addHistory(commit, path, lineNumber, pattern.Code, pattern.Message)
		}
	}
	return nil
}
