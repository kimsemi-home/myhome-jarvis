package security

import (
	"regexp"
	"strings"
)

var secretHistoryPattern = regexp.MustCompile(`(?i)(BEGIN (RSA|OPENSSH|EC|DSA) PRIVATE KEY|Authorization:[[:space:]]*Bearer[[:space:]]+[A-Za-z0-9._~+/=-]{20,}|(api[_-]?key|secret|password|token)[[:space:]]*[:=][[:space:]]*[A-Za-z0-9._~+/=-]{20,})`)

func privateIdentityHistoryPattern() string {
	oldOwner := strings.Join([]string{"kim", "jooyoon"}, "")
	oldTeam := strings.Join([]string{"kim-joo", "-yoon"}, "")
	localUser := strings.Join([]string{"al", "ice"}, "")
	return strings.Join([]string{
		oldOwner,
		oldTeam,
		"github\\.com/" + oldOwner,
		"/" + "Users",
		"(^|[^[:alnum:]_])" + localUser + "([^[:alnum:]_]|$)",
	}, "|")
}

func historyPatterns() []historyPattern {
	return []historyPattern{
		{
			Code:    "history_private_identity",
			Pattern: privateIdentityHistoryPattern(),
			Message: "git history must not contain private user, path, or old-repository identity markers",
		},
		{
			Code:    "history_secret_literal",
			Pattern: `(BEGIN (RSA|OPENSSH|EC|DSA) PRIVATE KEY|Authorization:[[:space:]]*Bearer[[:space:]]+[A-Za-z0-9._~+/=-]{20,}|(api[_-]?key|secret|password|token)[[:space:]]*[:=][[:space:]]*[A-Za-z0-9._~+/=-]{20,})`,
			Message: "git history must not contain secret-looking literal values",
		},
	}
}
