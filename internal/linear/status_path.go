package linear

import (
	"path/filepath"
	"strings"
)

func privateRelativePath(path string) string {
	if path == "" {
		return ""
	}
	slashed := strings.ReplaceAll(filepath.ToSlash(path), "\\", "/")
	for _, prefix := range []string{"data/private/", "data/lake/"} {
		if index := strings.Index(slashed, prefix); index >= 0 {
			return slashed[index:]
		}
	}
	if strings.HasPrefix(slashed, "/") || strings.Contains(slashed, ":/") {
		if index := strings.LastIndex(slashed, "/"); index >= 0 && index < len(slashed)-1 {
			return slashed[index+1:]
		}
		return strings.TrimLeft(slashed, "/")
	}
	return slashed
}
