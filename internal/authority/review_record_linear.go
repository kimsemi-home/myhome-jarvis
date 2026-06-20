package authority

import (
	"fmt"
	"strings"
)

func normalizeLinearIssueRef(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", nil
	}
	value = strings.ToUpper(value)
	parts := strings.Split(value, "-")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("linear issue ref must be an issue identifier")
	}
	for _, char := range parts[0] {
		if char < 'A' || char > 'Z' {
			return "", fmt.Errorf("linear issue ref must be an issue identifier")
		}
	}
	for _, char := range parts[1] {
		if char < '0' || char > '9' {
			return "", fmt.Errorf("linear issue ref must be an issue identifier")
		}
	}
	return value, nil
}
