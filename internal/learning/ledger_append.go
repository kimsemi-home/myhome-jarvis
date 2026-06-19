package learning

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func appendObservation(root string, policy Policy, observation Observation) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateLedger))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(observation)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = file.Write(data)
	return err
}
