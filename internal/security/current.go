package security

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func Check(root string) (Report, error) {
	report := Report{Root: ".", OK: true}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		return checkCurrentEntry(root, path, entry, walkErr, &report)
	})
	sortFindings(report.Findings)
	report.OK = len(report.Findings) == 0
	return report, err
}

func checkCurrentEntry(root string, path string, entry fs.DirEntry, walkErr error, report *Report) error {
	if walkErr != nil {
		return walkErr
	}
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	rel = filepath.ToSlash(rel)
	if rel == "." {
		return nil
	}
	if entry.IsDir() {
		return checkCurrentDir(entry, rel, report)
	}
	checkPath(rel, report)
	if shouldScanFileContent(rel) {
		return checkFileContent(path, rel, report)
	}
	return nil
}

func checkCurrentDir(entry fs.DirEntry, rel string, report *Report) error {
	if shouldSkipDir(entry.Name()) {
		return filepath.SkipDir
	}
	if strings.EqualFold(entry.Name(), "secrets") {
		report.add(rel, "forbidden_secret_dir", "secrets directories must not be tracked")
		return filepath.SkipDir
	}
	return nil
}
