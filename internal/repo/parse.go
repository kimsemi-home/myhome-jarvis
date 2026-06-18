package repo

import (
	"path/filepath"
	"strings"
)

func parsePorcelain(data string) ([]Change, []string) {
	var tracked []Change
	var untracked []string
	fields := strings.Split(data, "\x00")
	for index := 0; index < len(fields); index++ {
		field := fields[index]
		if len(field) < 4 {
			continue
		}
		code := field[:2]
		path := filepath.ToSlash(strings.TrimLeft(field[2:], " \t"))
		if code == "??" {
			untracked = append(untracked, path)
			continue
		}
		tracked = append(tracked, Change{Code: code, Path: path})
		if code[0] == 'R' || code[0] == 'C' {
			index++
		}
	}
	return tracked, untracked
}

func parseIgnoredPrivate(data string) []string {
	var paths []string
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "!! ") {
			continue
		}
		path := filepath.ToSlash(strings.TrimSpace(strings.TrimPrefix(line, "!! ")))
		if strings.HasPrefix(path, "data/private/") || strings.HasPrefix(path, "data/lake/") {
			paths = append(paths, path)
		}
	}
	return paths
}
