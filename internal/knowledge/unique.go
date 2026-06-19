package knowledge

import (
	"path/filepath"
	"strings"
)

func appendUnique(values []string, seen map[string]bool, value string) []string {
	if seen[value] {
		return values
	}
	seen[value] = true
	return append(values, value)
}

func appendUniqueRel(values []string, seen map[string]bool, value string) []string {
	value = filepath.ToSlash(strings.TrimSpace(value))
	if value == "" || seen[value] {
		return values
	}
	seen[value] = true
	return append(values, value)
}
