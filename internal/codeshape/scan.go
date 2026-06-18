package codeshape

import (
	"io/fs"
	"path/filepath"
)

func scanRoots(root string, policy Policy, status *Status) error {
	legacy := legacyMap(policy.LegacyDebtFiles)
	for _, sourceRoot := range policy.SourceRoots {
		if err := validateRelPath(sourceRoot); err != nil {
			return err
		}
		base := filepath.Join(root, filepath.FromSlash(sourceRoot))
		err := filepath.WalkDir(base, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			rel = filepath.ToSlash(rel)
			if shouldSkip(policy, rel, entry) {
				if entry.IsDir() && rel != sourceRoot {
					return filepath.SkipDir
				}
				return nil
			}
			if entry.IsDir() || !wantedExtension(policy, rel) {
				return nil
			}
			lines, err := countLines(path)
			if err != nil {
				return err
			}
			recordFile(policy, legacy, status, rel, lines)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
