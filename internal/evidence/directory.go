package evidence

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func countDirectoryFiles(root string, rel string) (int, bool, error) {
	path := filepath.Join(root, filepath.FromSlash(rel))
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	if !info.IsDir() {
		return 0, true, fmt.Errorf("evidence source is not a directory")
	}
	count := 0
	err = filepath.WalkDir(path, func(current string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type().IsRegular() {
			count++
		}
		return nil
	})
	return count, true, err
}
