package codexcost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func AttributeCost(root string, payload []byte) (AttributionResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return AttributionResult{}, err
	}
	request, err := decodeAttributionRequest(payload)
	if err != nil {
		return AttributionResult{}, err
	}
	record, err := normalizeAttribution(policy, request, time.Now().UTC())
	if err != nil {
		return AttributionResult{}, err
	}
	if err := appendAttribution(root, policy, record); err != nil {
		return AttributionResult{}, err
	}
	return attributionResult(record), nil
}

func decodeAttributionRequest(payload []byte) (AttributionRequest, error) {
	var request AttributionRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return AttributionRequest{}, fmt.Errorf("invalid attribution payload: %w", err)
	}
	return request, nil
}

func attributionResult(record AttributionRecord) AttributionResult {
	return AttributionResult{Scope: record.Scope, UnitKind: record.UnitKind,
		Amount: record.Amount, Basis: record.Basis,
		SubjectHash: record.SubjectHash, CostRef: record.CostRef,
		EvidenceRefCount: len(record.EvidenceRefs), RecordedAt: record.At}
}
