package domain

import (
	"sort"
	"strings"
)

func normalizeOwner(owner string) string {
	switch strings.TrimSpace(strings.ToLower(owner)) {
	case "user":
		return "user"
	case "spouse":
		return "spouse"
	case "household":
		return "household"
	default:
		return "unknown"
	}
}

func ownerBreakdownOrder[T any](owners map[string]*T) []string {
	keys := make([]string, 0, len(owners))
	for key := range owners {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return ownerRank(keys[i]) < ownerRank(keys[j])
	})
	return keys
}

func ownerRank(owner string) int {
	switch owner {
	case "user":
		return 0
	case "spouse":
		return 1
	case "household":
		return 2
	default:
		return 3
	}
}
