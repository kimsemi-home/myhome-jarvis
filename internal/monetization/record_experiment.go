package monetization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func RecordExperiment(root string, payload []byte) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	request, err := decodeRecordRequest(payload)
	if err != nil {
		return RecordResult{}, err
	}
	record, err := normalizeRecordRequest(policy, request, time.Now().UTC())
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
		return RecordRequest{}, fmt.Errorf("invalid monetization record payload: %w", err)
	}
	return request, nil
}
