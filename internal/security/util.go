package security

import (
	"path/filepath"
	"strings"
)

func hasPathSegment(rel string, segment string) bool {
	for _, part := range strings.Split(filepath.ToSlash(rel), "/") {
		if strings.EqualFold(part, segment) {
			return true
		}
	}
	return false
}

func isAllowedPrivatePlaceholder(rel string) bool {
	switch filepath.ToSlash(rel) {
	case "data/private/.keep", "data/private/.gitkeep", "data/lake/.keep", "data/lake/.gitkeep":
		return true
	default:
		return false
	}
}
