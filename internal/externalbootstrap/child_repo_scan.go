package externalbootstrap

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func scanChildRepoPublicSafety(root string, status *ChildRepoStatus) error {
	pattern := regexp.MustCompile("(?i)" + childPublicSafetyPattern())
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
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
			return childScanDir(rel, entry)
		}
		return scanChildRepoFile(path, rel, pattern, status)
	})
	status.PublicSafetyOK = countChildFindings(*status, "public_safety") == 0
	return err
}

func childScanDir(rel string, entry fs.DirEntry) error {
	name := strings.ToLower(entry.Name())
	if name == ".git" || rel == ".mhj/cache" || strings.HasPrefix(rel, ".mhj/cache/") {
		return filepath.SkipDir
	}
	return nil
}

func scanChildRepoFile(path string, rel string, pattern *regexp.Regexp, status *ChildRepoStatus) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		if pattern.MatchString(scanner.Text()) {
			status.addFinding(rel, "public_safety_forbidden_text",
				"child repo contains private identity, local path, or secret-looking text")
			break
		}
	}
	return scanner.Err()
}
