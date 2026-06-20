package codexcost

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
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy, time.Now().UTC())
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateUsageLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return Status{}, err
	}
	defer file.Close()
	status.Exists = true

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		if err := scanRecord(policy, scanner.Text(), &status); err != nil {
			return Status{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	status.BudgetState = budgetState(policy, status.TotalUnits)
	return status, nil
}

func scanRecord(policy Policy, line string, status *Status) error {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	var record Record
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		status.InvalidRecordCount++
		return nil
	}
	normalized, err := normalizeRecord(policy, record)
	if err != nil {
		classifyRecordError(status, err)
		return nil
	}
	applyRecord(status, normalized)
	return nil
}
