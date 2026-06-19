package supervisor

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteDaemonState(root string, state DaemonState) (string, error) {
	if strings.TrimSpace(root) == "" {
		return "", errors.New("root is required")
	}
	state, err := normalizeState(state)
	if err != nil {
		return "", err
	}
	path := statePath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return "", err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return "", err
	}
	return daemonStateRelativePath, nil
}

func ReadDaemonState(root string) (DaemonState, error) {
	data, err := os.ReadFile(statePath(root))
	if err != nil {
		return DaemonState{}, err
	}
	var state DaemonState
	if err := json.Unmarshal(data, &state); err != nil {
		return DaemonState{}, fmt.Errorf("read daemon state: %w", err)
	}
	if state.Name == "" {
		state.Name = daemonStateName
	}
	return state, nil
}
