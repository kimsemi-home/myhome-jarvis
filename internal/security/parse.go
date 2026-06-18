package security

import (
	"path/filepath"
	"strconv"
	"strings"
)

func parseGitGrepLine(line string) (string, string, int, bool) {
	if len(line) < 42 || line[40] != ':' {
		return "", "", 0, false
	}
	commit := line[:40]
	rest := line[41:]
	pathEnd := strings.IndexByte(rest, ':')
	if pathEnd < 0 {
		return "", "", 0, false
	}
	path := filepath.ToSlash(rest[:pathEnd])
	rest = rest[pathEnd+1:]
	lineEnd := strings.IndexByte(rest, ':')
	if lineEnd < 0 {
		return "", "", 0, false
	}
	lineNumber, err := strconv.Atoi(rest[:lineEnd])
	if err != nil {
		return "", "", 0, false
	}
	return commit, path, lineNumber, true
}

func splitOutputLines(output string) []string {
	trimmed := strings.TrimRight(output, "\n")
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, "\n")
}
