package security

import (
	"regexp"
	"strings"
	"testing"
)

func TestCheckMetadataAllowsPublicGitHubNoreplyIdentity(t *testing.T) {
	report := HistoryReport{OK: true}
	owner := strings.Join([]string{"kim", "jooyoon"}, "")
	line := "commit\x1f" + owner + "\x1f115961382+" + owner + "@users.noreply.github.com\x1fGitHub\x1fnoreply@github.com\x1fmerge commit"
	checkMetadataLine(line, regexp.MustCompile(`(?i)`+privateIdentityHistoryPattern()), &report)
	if len(report.Findings) != 0 {
		t.Fatalf("public GitHub noreply identity was rejected: %+v", report.Findings)
	}
}

func TestCheckMetadataRejectsPrivateIdentityOutsidePublicNoreply(t *testing.T) {
	report := HistoryReport{OK: true}
	owner := strings.Join([]string{"kim", "jooyoon"}, "")
	line := "commit\x1f" + owner + "\x1fprivate@example.com\x1fAuthor\x1fauthor@example.com\x1fprivate identity"
	checkMetadataLine(line, regexp.MustCompile(`(?i)`+privateIdentityHistoryPattern()), &report)
	if !historyHasCode(report, "history_private_identity_metadata") {
		t.Fatalf("private metadata was not rejected: %+v", report.Findings)
	}
}
