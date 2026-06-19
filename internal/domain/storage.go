package domain

import (
	"encoding/json"
	"os"
)

func ReadStoragePolicy(path string) (StoragePolicy, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return StoragePolicy{}, err
	}
	var policy StoragePolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		return StoragePolicy{}, err
	}
	return policy, nil
}
