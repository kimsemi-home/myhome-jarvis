package domain

import (
	"sort"
	"strings"
)

func recordCategory(categories map[string]bool, category string) {
	if strings.TrimSpace(category) != "" {
		categories[category] = true
	}
}

func sortedKeys(values map[string]bool) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func clampScore(value int) int {
	switch {
	case value < 0:
		return 0
	case value > 100:
		return 100
	default:
		return value
	}
}
