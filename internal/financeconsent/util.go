package financeconsent

import "strings"

func normalizeToken(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func missingValues(values []string, required []string) []string {
	missing := []string{}
	for _, item := range required {
		if !contains(values, item) {
			missing = append(missing, item)
		}
	}
	return missing
}
