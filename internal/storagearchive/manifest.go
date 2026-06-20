package storagearchive

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func appendManifest(root string, manifestPath string, entries []manifestEntry) error {
	if len(entries) == 0 {
		return nil
	}
	path := filepath.Join(root, filepath.FromSlash(manifestPath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, entry := range entries {
		if err := encoder.Encode(entry); err != nil {
			return err
		}
	}
	return nil
}
