package codexcost

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"
)

func GuardLoop(root string, payload []byte) (GuardResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return GuardResult{}, err
	}
	request, err := decodeGuardRequest(payload)
	if err != nil {
		return GuardResult{}, err
	}
	guard, err := normalizeGuardRequest(policy, request)
	if err != nil {
		return GuardResult{}, err
	}
	costStatus, err := StatusForRoot(root)
	if err != nil {
		return GuardResult{}, err
	}
	sustainability, err := codexsustainability.StatusForRoot(root)
	if err != nil {
		return GuardResult{}, err
	}
	return evaluateGuard(policy, guard, costStatus, sustainability), nil
}

func decodeGuardRequest(payload []byte) (GuardRequest, error) {
	var request GuardRequest
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return GuardRequest{}, fmt.Errorf("invalid codex cost guard payload: %w", err)
	}
	return request, nil
}
