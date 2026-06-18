package linear

import "strings"

func publicIssueKey(value string) string {
	value = strings.TrimSpace(value)
	parts := strings.Split(value, "-")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return ""
	}
	if !isPublicIssueKeyPrefix(parts[0]) || !isPublicIssueKeyNumber(parts[1]) {
		return ""
	}
	return strings.ToUpper(parts[0]) + "-" + parts[1]
}

func isPublicIssueKeyPrefix(value string) bool {
	for _, char := range value {
		if !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}

func isPublicIssueKeyNumber(value string) bool {
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
