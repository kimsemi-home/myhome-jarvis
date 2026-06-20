package monetization

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

func scanLedger(reader io.Reader, policy Policy, status *Status) error {
	seen := map[string]bool{}
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		scanLine(policy, status, seen, scanner.Text())
	}
	return scanner.Err()
}

func scanLine(policy Policy, status *Status, seen map[string]bool, line string) {
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
		classifyRecordError(status, err)
		return
	}
	applyRecord(status, seen, normalized)
}

func applyRecord(status *Status, seen map[string]bool, record Record) {
	if !seen[record.ExperimentID] {
		status.ExperimentCount++
		seen[record.ExperimentID] = true
	}
	status.DecisionCount++
	status.ByState[record.State]++
	status.ByDecisionKind[record.DecisionKind]++
	status.ByReviewStatus[record.ReviewStatus]++
	status.ByExpectedValueBand[record.ExpectedValueBand]++
	if record.ReviewStatus == "requested" || record.State == "review_required" {
		status.ReviewRequiredCount++
	}
	if record.ExpectedValueBand == "unknown" {
		status.ExpectedValueUnknownCount++
	}
	updateLastObserved(status, record.At)
}
