package authority

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReviewRecordLedgerForRoot(root string, policy Policy, requestID string) (ReviewRecordLedgerSummary, error) {
	summary := ReviewRecordLedgerSummary{
		RequestID:     requestID,
		LedgerState:   "missing",
		ApprovalState: "not_recorded",
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateReviewRequestLedger))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return summary, nil
	}
	if err != nil {
		return ReviewRecordLedgerSummary{}, err
	}
	defer file.Close()
	summary.LedgerState = "unmatched"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		applyLedgerLine(scanner.Text(), requestID, &summary)
	}
	if err := scanner.Err(); err != nil {
		return ReviewRecordLedgerSummary{}, err
	}
	if summary.Recorded {
		summary.LedgerState = "recorded_pending_review"
		summary.ApprovalState = "not_approved"
	}
	return summary, nil
}

func applyLedgerLine(line string, requestID string, summary *ReviewRecordLedgerSummary) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	var record ReviewRecord
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		summary.InvalidRecordCount++
		return
	}
	if record.RequestID != requestID {
		return
	}
	if !reviewRecordPending(record) {
		summary.InvalidRecordCount++
		return
	}
	summary.Recorded = true
	summary.RecordCount++
	summary.LastRecordedAt = record.At
}

func reviewRecordPending(record ReviewRecord) bool {
	return record.RequestState == "ready" &&
		record.QueueState == "pending_human_review" &&
		record.ApprovalState == "not_approved" &&
		record.PublicSafe &&
		!record.ApprovalGranted &&
		!record.ExternalWritesAllowed &&
		!record.SelfApprovalAllowed
}
