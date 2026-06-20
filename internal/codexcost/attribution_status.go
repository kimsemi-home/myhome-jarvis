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

func AttributionStatusForRoot(root string) (AttributionStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return AttributionStatus{}, err
	}
	status := newAttributionStatus(policy)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateAttributionLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return finalizeAttributionStatus(status), nil
	}
	if err != nil {
		return AttributionStatus{}, err
	}
	defer file.Close()
	status.Exists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		scanAttributionRecord(policy, scanner.Text(), &status)
	}
	if err := scanner.Err(); err != nil {
		return AttributionStatus{}, err
	}
	return finalizeAttributionStatus(status), nil
}

func newAttributionStatus(policy Policy) AttributionStatus {
	return AttributionStatus{LedgerPath: policy.PrivateAttributionLedger,
		ByScope: map[string]int64{}, SubjectCountByScope: map[string]int{},
		subjectsByScope: map[string]map[string]bool{},
		costRefUnits:    map[string]int64{},
		CheckedAt:       time.Now().UTC().Format(time.RFC3339)}
}

func scanAttributionRecord(
	policy Policy,
	line string,
	status *AttributionStatus,
) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	var record AttributionRecord
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		status.InvalidRecordCount++
		return
	}
	normalized, err := normalizeAttributionRecord(policy, record)
	if err != nil {
		classifyAttributionError(status, err)
		return
	}
	applyAttribution(status, normalized)
}
