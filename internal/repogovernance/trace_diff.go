package repogovernance

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func StagedFiles(root string) ([]string, error) {
	return gitFiles(root, "diff", "--cached", "--name-only", "--diff-filter=ACMR")
}

func RangeFiles(root, base string) ([]string, error) {
	if base == "" {
		return nil, errors.New("base ref is required")
	}
	return gitFiles(root, "diff", "--name-only", "--diff-filter=ACMR", base+"...HEAD")
}

func gitFiles(root string, args ...string) ([]string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var files []string
	for _, line := range strings.Split(string(out), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			files = append(files, filepath.ToSlash(line))
		}
	}
	sort.Strings(files)
	return files, nil
}

func CheckChanges(manifest Manifest, files []string) error {
	for _, group := range manifest.Groups {
		source := match(files, group.DocumentSources)
		generated := match(files, group.GeneratedDocuments)
		technical := match(files, append(append([]string{}, group.Code...), group.Tests...))
		if (source || generated || technical) && (source != generated || source != technical) {
			return fmt.Errorf("one-sided change in %s: source=%t generated=%t technical=%t", group.ID, source, generated, technical)
		}
	}
	return nil
}

func match(files, paths []string) bool {
	for _, file := range files {
		for _, path := range paths {
			if file == path || strings.HasPrefix(file, strings.TrimSuffix(path, "/")+"/") {
				return true
			}
		}
	}
	return false
}
