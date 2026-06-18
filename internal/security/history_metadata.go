package security

import (
	"regexp"
	"strings"
)

func checkHistoryMetadata(root string, report *HistoryReport) error {
	lines, err := gitLines(root, "log", "--all", "--format=%H%x1f%an%x1f%ae%x1f%cn%x1f%ce%x1f%s")
	if err != nil {
		return err
	}
	privateIdentity := regexp.MustCompile(`(?i)` + privateIdentityHistoryPattern())
	for _, line := range lines {
		checkMetadataLine(line, privateIdentity, report)
	}
	return nil
}

func checkMetadataLine(line string, privateIdentity *regexp.Regexp, report *HistoryReport) {
	parts := strings.Split(line, "\x1f")
	if len(parts) < 6 {
		return
	}
	commit := parts[0]
	for _, field := range parts[1:] {
		if privateIdentity.MatchString(field) {
			report.addHistory(commit, "(commit metadata)", 0, "history_private_identity_metadata", "git commit metadata must not contain private identity markers")
			return
		}
		if secretHistoryPattern.MatchString(field) {
			report.addHistory(commit, "(commit metadata)", 0, "history_secret_metadata", "git commit metadata must not contain secret-looking literal values")
			return
		}
	}
}
