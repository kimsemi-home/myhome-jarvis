package knowledge

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func indexFiles(root string, roots []string) ([]string, error) {
	seen := map[string]bool{}
	var files []string
	for _, relRoot := range roots {
		found, err := indexRoot(root, clean(relRoot), seen)
		if err != nil {
			return nil, err
		}
		files = append(files, found...)
	}
	sort.Strings(files)
	return files, nil
}

func indexRoot(root string, relRoot string, seen map[string]bool) ([]string, error) {
	if relRoot == "" || strings.Contains(relRoot, "..") || filepath.IsAbs(filepath.FromSlash(relRoot)) {
		return nil, fmt.Errorf("invalid knowledge index root %q", relRoot)
	}
	path := filepath.Join(root, filepath.FromSlash(relRoot))
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) && strings.HasPrefix(relRoot, "data/private/") {
			return nil, nil
		}
		return nil, err
	}
	if !info.IsDir() {
		if indexableFile(path) {
			return appendUnique(nil, seen, path), nil
		}
		return nil, nil
	}
	return walkIndexRoot(path, seen)
}
