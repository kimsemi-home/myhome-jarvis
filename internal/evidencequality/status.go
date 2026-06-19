package evidencequality

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
	checkedAt := time.Now().UTC()
	status := newStatus(policy, checkedAt)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateSnapshotLedger)))
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
		if err := scanSnapshot(policy, checkedAt, scanner.Text(), &status); err != nil {
			return Status{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		return Status{}, err
	}
	status.ReassessmentDebtCount = reassessmentDebt(status)
	return status, nil
}

func scanSnapshot(policy Policy, checkedAt time.Time, line string, status *Status) error {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	var snapshot Snapshot
	if err := json.Unmarshal([]byte(line), &snapshot); err != nil {
		status.InvalidSnapshotCount++
		return nil
	}
	normalized, err := normalizeSnapshot(policy, snapshot)
	if err != nil {
		classifySnapshotError(status, err)
		return nil
	}
	applySnapshot(status, policy, normalized, checkedAt)
	return nil
}
