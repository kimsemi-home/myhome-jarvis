package security

import (
	"path/filepath"
	"strings"
)

func shouldSkipDir(name string) bool {
	switch name {
	case ".git", "target", "build", "dist", "bin", ".dart_tool":
		return true
	default:
		return false
	}
}

func shouldSkipFile(rel string) bool {
	switch filepath.Base(rel) {
	case ".git", ".flutter-plugins", ".flutter-plugins-dependencies", ".packages":
		return true
	default:
		return false
	}
}

func checkPath(rel string, report *Report) {
	base := strings.ToLower(filepath.Base(rel))
	ext := strings.ToLower(filepath.Ext(rel))
	if isEnvFile(rel) {
		report.add(rel, "forbidden_env_file", "environment files must stay local")
	}
	if isPrivateOrLake(rel) {
		return
	}
	checkLanguagePath(rel, ext, report)
	checkDependencyPath(rel, base, report)
	checkSensitivePath(rel, base, report)
}

func isEnvFile(rel string) bool {
	return rel == ".env" || (strings.HasPrefix(rel, ".env.") && rel != ".env.example")
}

func isPrivateOrLake(rel string) bool {
	return strings.HasPrefix(rel, "data/private/") || strings.HasPrefix(rel, "data/lake/")
}
