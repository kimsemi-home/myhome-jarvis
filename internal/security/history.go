package security

import "os/exec"

func CheckHistory(root string) (HistoryReport, error) {
	report := HistoryReport{Root: ".", OK: true}
	if _, err := exec.LookPath("git"); err != nil {
		return report, err
	}
	commits, err := gitLines(root, "rev-list", "--all")
	if err != nil {
		return report, err
	}
	if len(commits) == 0 {
		return report, nil
	}
	if err := checkHistoryPaths(root, &report); err != nil {
		return report, err
	}
	if err := checkHistoryMetadata(root, &report); err != nil {
		return report, err
	}
	for _, pattern := range historyPatterns() {
		if err := checkHistoryContent(root, commits, pattern, &report); err != nil {
			return report, err
		}
	}
	sortHistoryFindings(report.Findings)
	report.OK = len(report.Findings) == 0
	return report, nil
}
