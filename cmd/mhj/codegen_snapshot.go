package main

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
)

func generatedSnapshot(root string) (map[string][]byte, error) {
	generatedRoot := filepath.Join(root, "generated")
	files := map[string][]byte{}
	err := filepath.WalkDir(generatedRoot, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(rel)] = body
		return nil
	})
	return files, err
}

func changedGeneratedFiles(before map[string][]byte, after map[string][]byte) []string {
	seen := map[string]bool{}
	var changed []string
	for path, body := range before {
		seen[path] = true
		if next, ok := after[path]; !ok || !bytes.Equal(body, next) {
			changed = append(changed, path)
		}
	}
	for path := range after {
		if !seen[path] {
			changed = append(changed, path)
		}
	}
	sort.Strings(changed)
	return changed
}
