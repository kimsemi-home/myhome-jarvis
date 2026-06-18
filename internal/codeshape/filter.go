package codeshape

import (
	"io/fs"
	"strings"
)

func shouldSkip(policy Policy, rel string, entry fs.DirEntry) bool {
	for _, prefix := range policy.ExcludedPrefixes {
		if strings.HasPrefix(rel, prefix) {
			return true
		}
	}
	if entry.IsDir() {
		switch entry.Name() {
		case ".git", ".dart_tool", "build", "node_modules", "target":
			return true
		}
	}
	return false
}

func wantedExtension(policy Policy, rel string) bool {
	for _, extension := range policy.Extensions {
		if strings.HasSuffix(rel, extension) {
			return true
		}
	}
	return false
}
