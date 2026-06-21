package authority

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func RecordApprovalDecision(root string, payload []byte) (ApprovalDecisionResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ApprovalDecisionResult{}, err
	}
	packet, err := externalevidence.RepoSplitDecisionPacketForRoot(root)
	if err != nil {
		return ApprovalDecisionResult{}, err
	}
	request, err := decodeApprovalDecisionRequest(payload)
	if err != nil {
		return ApprovalDecisionResult{}, err
	}
	record, err := normalizeApprovalDecisionRequest(policy, request, packet, time.Now().UTC())
	if err != nil {
		return ApprovalDecisionResult{}, err
	}
	if err := appendApprovalDecision(root, policy, record); err != nil {
		return ApprovalDecisionResult{}, err
	}
	return resultForApprovalDecision(record), nil
}

func decodeApprovalDecisionRequest(payload []byte) (ApprovalDecisionRequest, error) {
	var request ApprovalDecisionRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return ApprovalDecisionRequest{}, fmt.Errorf("invalid approval payload: %w", err)
	}
	return request, nil
}
