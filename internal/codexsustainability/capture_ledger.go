package codexsustainability

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func capturedVersionExists(root string, policy Policy, version string) bool {
	if strings.TrimSpace(version) == "" {
		return false
	}
	file, err := os.Open(ledgerPath(root, policy))
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	if err != nil {
		return false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var record Record
		if json.Unmarshal([]byte(scanner.Text()), &record) == nil &&
			record.TrendBaselineVersion == version {
			return true
		}
	}
	return false
}

func appendRecords(root string, policy Policy, records []Record) error {
	path := ledgerPath(root, policy)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, record := range records {
		if _, err := normalizeRecord(policy, record); err != nil {
			return err
		}
		body, err := json.Marshal(record)
		if err != nil {
			return err
		}
		if _, err := file.Write(append(body, '\n')); err != nil {
			return err
		}
	}
	return nil
}

func ledgerPath(root string, policy Policy) string {
	return filepath.Join(root, filepath.FromSlash(policy.PrivateEvidenceLedger))
}
