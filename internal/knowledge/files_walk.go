package knowledge

import (
	"io/fs"
	"path/filepath"
)

func walkIndexRoot(path string, seen map[string]bool) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if skippedIndexDir(entry.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if indexableFile(path) {
			files = appendUnique(files, seen, path)
		}
		return nil
	})
	return files, err
}

func skippedIndexDir(name string) bool {
	switch name {
	case ".git", ".dart_tool", "build", "dist", "target", "bin":
		return true
	default:
		return false
	}
}
