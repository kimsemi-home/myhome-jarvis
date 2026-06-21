package authority

import (
	"fmt"
	"strings"
)

func normalizeApprovalTarget(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" || unsafeApprovalText(value) {
		return "", fmt.Errorf("approval target is not public-safe")
	}
	return value, nil
}

func normalizeReviewerBoundary(value string) (string, error) {
	value = normalizeToken(value)
	if value == "" || !strings.HasPrefix(value, "human_") ||
		strings.Contains(value, "codex") || unsafeApprovalText(value) {
		return "", fmt.Errorf("approval reviewer boundary must be human")
	}
	return value, nil
}

func unsafeApprovalText(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	for _, forbidden := range []string{
		"data/private", "raw_private", "raw_rationale",
		"token", "secret", "credential", "cookie",
		"private_key", "id_rsa", "http://", "https://",
		"/" + "users" + "/", "local_absolute_path",
	} {
		if strings.Contains(value, forbidden) {
			return true
		}
	}
	return false
}
