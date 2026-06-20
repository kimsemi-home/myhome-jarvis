package security

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func readStatusCache(root string, key statusCacheKey) (statusCacheRecord, bool, error) {
	content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(statusCachePath)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return statusCacheRecord{}, false, nil
		}
		return statusCacheRecord{}, false, err
	}
	var record statusCacheRecord
	if err := json.Unmarshal(content, &record); err != nil {
		return statusCacheRecord{}, false, nil
	}
	if !record.matches(key) {
		return statusCacheRecord{}, false, nil
	}
	return record, true, nil
}

func writeStatusCache(root string, record statusCacheRecord) error {
	path := filepath.Join(root, filepath.FromSlash(statusCachePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	content, err := json.Marshal(record)
	if err != nil {
		return err
	}
	content = append(content, '\n')
	return os.WriteFile(path, content, 0o600)
}

func (record statusCacheRecord) matches(key statusCacheKey) bool {
	return record.Version == statusCacheVersion &&
		record.Key == key.Key &&
		record.Head == key.Head &&
		record.InputHash == key.InputHash &&
		record.ValidationCommand == statusCacheValidation
}
