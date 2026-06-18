package translation

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func inspectManifests(root string, policy Policy, status *Status) error {
	manifestRoot := filepath.Join(root, filepath.FromSlash(policy.PrivateManifestRoot))
	info, err := os.Stat(manifestRoot)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	status.ManifestRootExists = true
	if !info.IsDir() {
		status.InvalidManifestCount++
		return nil
	}
	return filepath.WalkDir(manifestRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		return inspectManifestEntry(path, entry, walkErr, policy, status)
	})
}

func inspectManifestEntry(path string, entry fs.DirEntry, walkErr error, policy Policy, status *Status) error {
	if walkErr != nil {
		return walkErr
	}
	if entry.IsDir() || filepath.Ext(path) != ".json" {
		return nil
	}
	status.ManifestCount++
	return inspectManifestFile(path, policy, status)
}
