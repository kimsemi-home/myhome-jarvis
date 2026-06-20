package commandcenter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const visionPolicyPath = "generated/assistant_vision.generated.json"

type visionPolicy struct {
	Context           string         `json:"context"`
	Version           string         `json:"version"`
	Mission           string         `json:"mission"`
	OperatingMode     string         `json:"operating_mode"`
	CapabilityPillars []visionPillar `json:"capability_pillars"`
	Guardrails        []string       `json:"guardrails"`
}

type visionPillar struct {
	Key string `json:"key"`
}

func readVisionPolicy(root string) (visionPolicy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(visionPolicyPath)))
	if err != nil {
		return visionPolicy{}, err
	}
	var policy visionPolicy
	if err := json.Unmarshal(body, &policy); err != nil {
		return visionPolicy{}, err
	}
	return policy, nil
}
