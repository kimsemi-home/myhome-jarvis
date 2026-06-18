package controlplane

import (
	"sort"
	"strings"
)

func normalizeRefs(values []string) []string {
	seen := map[string]bool{}
	refs := make([]string, 0, len(values))
	for _, value := range values {
		ref := cleanRef(value)
		if ref == "" || seen[ref] {
			continue
		}
		seen[ref] = true
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	return refs
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	normalized := make([]string, 0, len(values))
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

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func publicText(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}
