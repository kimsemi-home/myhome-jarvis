package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func readVerificationJSON[T any](root string, rel string) (T, error) {
	var value T
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return value, err
	}
	if err := json.Unmarshal(body, &value); err != nil {
		return value, fmt.Errorf("%s: %w", rel, err)
	}
	return value, nil
}
