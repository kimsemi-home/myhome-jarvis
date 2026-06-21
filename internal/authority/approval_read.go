package authority

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func readApprovalDecisionLedger(
	root string,
	policy Policy,
) ([]ApprovalDecisionRecord, int, bool, error) {
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateApprovalDecisionLedger))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, 0, true, nil
	}
	if err != nil {
		return nil, 0, false, err
	}
	defer file.Close()
	records := []ApprovalDecisionRecord{}
	invalid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record, ok := parseApprovalRecord(scanner.Text())
		if !ok {
			invalid++
			continue
		}
		records = append(records, record)
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, false, err
	}
	return records, invalid, false, nil
}

func parseApprovalRecord(line string) (ApprovalDecisionRecord, bool) {
	var record ApprovalDecisionRecord
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return ApprovalDecisionRecord{}, false
	}
	return record, approvalRecordPublicSafe(record)
}
