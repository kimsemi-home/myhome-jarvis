package monetization

import "strings"

func normalizeToken(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeList(values []string) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = normalizeToken(value)
		if value != "" {
			out = append(out, value)
		}
	}
	return out
}

func contains(values []string, value string) bool {
	value = normalizeToken(value)
	for _, candidate := range values {
		if normalizeToken(candidate) == value {
			return true
		}
	}
	return false
}
