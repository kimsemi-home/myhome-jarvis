package codexsustainability

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func StatusForRoot(root string) (Status, error) {
	return statusForRootAt(root, time.Now().UTC())
}

func statusForRootAt(root string, now time.Time) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy, now)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateEvidenceLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return finalizeStatus(policy, status, now), nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()
	status.Exists = true
	if err := scanLedger(policy, file, &status); err != nil {
		return Status{}, err
	}
	return finalizeStatus(policy, status, now), nil
}

func scanLedger(policy Policy, file *os.File, status *Status) error {
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		scanRecord(policy, scanner.Text(), status)
	}
	return scanner.Err()
}

func scanRecord(policy Policy, line string, status *Status) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	var record Record
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		status.InvalidRecordCount++
		return
	}
	normalized, err := normalizeRecord(policy, record)
	if err != nil {
		classifyRecordError(status, record, err)
		return
	}
	applyRecord(status, normalized)
}
