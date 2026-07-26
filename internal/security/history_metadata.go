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
	for index, field := range parts[1:] {
		if !isPublicGitHubIdentityField(parts[1:], index) && privateIdentity.MatchString(field) {
			report.addHistory(commit, "(commit metadata)", 0, "history_private_identity_metadata", "git commit metadata must not contain private identity markers")
			return
		}
		if secretHistoryPattern.MatchString(field) {
			report.addHistory(commit, "(commit metadata)", 0, "history_secret_metadata", "git commit metadata must not contain secret-looking literal values")
			return
		}
	}
}

func isPublicGitHubIdentityField(fields []string, index int) bool {
	emailIndex := -1
	switch index {
	case 0, 1:
		emailIndex = 1
	case 2, 3:
		emailIndex = 3
	}
	return emailIndex >= 0 && emailIndex < len(fields) && strings.HasSuffix(
		strings.ToLower(strings.TrimSpace(fields[emailIndex])),
		"@users.noreply.github.com",
	)
}
