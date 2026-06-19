package scheduler

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func WriteState(root string, state State) error {
	if strings.TrimSpace(state.Name) == "" {
		return errors.New("scheduler state name is required")
	}
	dir := filepath.Join(root, "data", "private", "scheduler")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(statePath(root, state.Name), data, 0o600)
}

func statePath(root string, name string) string {
	return filepath.Join(root, "data", "private", "scheduler", name+"-state.json")
}

func checkpointPath(root string, path string) string {
	cleaned := filepath.Clean(path)
	if root != "" {
		relative, err := filepath.Rel(root, cleaned)
		if isLocalRelative(relative, err) {
			return filepath.ToSlash(relative)
		}
	}
	return filepath.ToSlash(cleaned)
}

func isLocalRelative(relative string, err error) bool {
	return err == nil &&
		relative != "." &&
		relative != ".." &&
		!strings.HasPrefix(relative, ".."+string(filepath.Separator))
}
