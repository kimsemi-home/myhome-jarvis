package knowledge

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadRegistry(root string) (Registry, error) {
	registry, err := readRegistryUnchecked(root)
	if err != nil {
		return Registry{}, err
	}
	if failures := registryFailures(root, registry); len(failures) > 0 {
		return Registry{}, errors.New(strings.Join(failures, "; "))
	}
	return registry, nil
}

func readRegistryUnchecked(root string) (Registry, error) {
	path := filepath.Join(root, filepath.FromSlash(RegistryRelativePath))
	file, err := os.Open(path)
	if err != nil {
		return Registry{}, err
	}
	defer file.Close()

	var registry Registry
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&registry); err != nil {
		return Registry{}, err
	}
	return registry, nil
}
