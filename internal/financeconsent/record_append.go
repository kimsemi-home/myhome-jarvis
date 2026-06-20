package financeconsent

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func appendRecord(root string, policy Policy, record Record) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateConsentLedger))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}
	_, err = file.Write(append(body, '\n'))
	return err
}
