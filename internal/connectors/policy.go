package connectors

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func readGeneratedPolicy(root string) (generatedPolicy, error) {
	body, err := os.ReadFile(generatedPolicyPath(root))
	if err != nil {
		return generatedPolicy{}, err
	}
	var policy generatedPolicy
	if err := json.Unmarshal(body, &policy); err != nil {
		return generatedPolicy{}, err
	}
	return policy, nil
}

func generatedPolicyPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(generatedConnectorPath))
}
