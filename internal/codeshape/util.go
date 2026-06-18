package codeshape

import (
	"fmt"
	"path/filepath"
	"strings"
)

func validateRelPath(path string) error {
	path = strings.TrimSpace(filepath.ToSlash(path))
	if path == "" || filepath.IsAbs(filepath.FromSlash(path)) || strings.Contains(path, "..") {
		return fmt.Errorf("code-shape path %q must be repo-relative", path)
	}
	return nil
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
