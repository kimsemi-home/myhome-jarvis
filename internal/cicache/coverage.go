package cicache

import (
	"path/filepath"
	"strings"
)

func generatedCoverageCount(artifacts []string, patterns []string) int {
	count := 0
	for _, artifact := range artifacts {
		if strings.HasPrefix(artifact, "generated/") &&
			coveredByAny(artifact, patterns) {
			count++
		}
	}
	return count
}

func generatedArtifactCount(artifacts []string) int {
	count := 0
	for _, artifact := range artifacts {
		if strings.HasPrefix(artifact, "generated/") {
			count++
		}
	}
	return count
}

func coveredByAny(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if covers(pattern, path) {
			return true
		}
	}
	return false
}

func covers(pattern string, path string) bool {
	if pattern == path {
		return true
	}
	if strings.HasSuffix(pattern, "/**") {
		return strings.HasPrefix(path, strings.TrimSuffix(pattern, "**"))
	}
	if strings.HasSuffix(pattern, "/*.json") {
		return filepath.Dir(path) == strings.TrimSuffix(pattern, "/*.json") &&
			strings.HasSuffix(path, ".json")
	}
	return false
}
