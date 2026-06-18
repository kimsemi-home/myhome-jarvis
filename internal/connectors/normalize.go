package connectors

import (
	"sort"
	"strings"
)

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	var normalized []string
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func countReadOnlyOperations(operations []string) int {
	count := 0
	for _, operation := range operations {
		if isReadOnlyOperation(operation) {
			count++
		}
	}
	return count
}

func isReadOnlyOperation(operation string) bool {
	switch operation {
	case "read_fixture", "summarize", "recommend_review":
		return true
	default:
		return false
	}
}
