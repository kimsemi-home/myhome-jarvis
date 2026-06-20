package codexcost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func RecordUsage(root string, payload []byte) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	request, err := decodeRecordRequest(payload)
	if err != nil {
		return RecordResult{}, err
	}
	record, err := normalizeUsageRecord(policy, request, time.Now().UTC())
	if err != nil {
		return RecordResult{}, err
	}
	if err := appendRecord(root, policy, record); err != nil {
		return RecordResult{}, err
	}
	status, err := StatusForRoot(root)
	if err != nil {
		return RecordResult{}, err
	}
	return resultForRecord(record, status), nil
}

func decodeRecordRequest(payload []byte) (RecordRequest, error) {
	var request RecordRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return RecordRequest{}, fmt.Errorf("invalid codex cost record payload: %w", err)
	}
	return request, nil
}

func resultForRecord(record Record, status Status) RecordResult {
	return RecordResult{
		Scope:            record.Scope,
		UnitKind:         record.UnitKind,
		Amount:           record.Amount,
		Status:           record.Status,
		EvidenceRefCount: len(record.EvidenceRefs),
		BudgetState:      status.BudgetState,
		RecordedAt:       record.At,
	}
}
