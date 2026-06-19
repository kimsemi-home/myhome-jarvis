package learning

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Record(root string, payload []byte) (RecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return RecordResult{}, err
	}
	request, err := decodeRecordRequest(payload)
	if err != nil {
		return RecordResult{}, err
	}
	observation, err := normalizeObservation(policy, request)
	if err != nil {
		return RecordResult{}, err
	}
	if err := appendObservation(root, policy, observation); err != nil {
		return RecordResult{}, err
	}
	return RecordResult{
		ID:               observation.ID,
		Path:             policy.PrivateLedger,
		Kind:             observation.Kind,
		Stage:            observation.Stage,
		Status:           observation.Status,
		EvidenceRefCount: len(observation.EvidenceRefs),
		RecordedAt:       observation.At,
	}, nil
}

func decodeRecordRequest(payload []byte) (RecordRequest, error) {
	var request RecordRequest
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return RecordRequest{}, fmt.Errorf("invalid learning record payload: %w", err)
	}
	return request, nil
}
