package controlplane

import (
	"path/filepath"
	"strings"
)

func cleanRef(value string) string {
	return filepath.ToSlash(strings.TrimSpace(value))
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
