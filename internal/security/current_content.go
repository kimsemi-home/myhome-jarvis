package security

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func shouldScanFileContent(rel string) bool {
	rel = filepath.ToSlash(rel)
	if strings.HasPrefix(rel, "data/private/") || strings.HasPrefix(rel, "data/lake/") {
		return false
	}
	if strings.ToLower(filepath.Base(rel)) == "cargo.lock" {
		return true
	}
	ext := strings.ToLower(filepath.Ext(rel))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".pdf", ".zip", ".gz", ".tar", ".tgz", ".parquet":
		return false
	default:
		return true
	}
}

func checkFileContent(path string, rel string, report *Report) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return scanFileContent(file, rel, report)
}

func scanFileContent(file *os.File, rel string, report *Report) error {
	privateIdentity := regexp.MustCompile(`(?i)` + privateIdentityHistoryPattern())
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	line := 0
	for scanner.Scan() {
		line++
		checkCurrentTextLine(rel, line, scanner.Text(), privateIdentity, report)
	}
	return scanner.Err()
}
