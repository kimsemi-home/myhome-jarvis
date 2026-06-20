package mediareadiness

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const PolicyRelativePath = "generated/media_readiness.generated.json"

func ReadPolicy(root string) (Policy, error) {
	body, err := os.ReadFile(filepath.Join(root, PolicyRelativePath))
	if err != nil {
		return Policy{}, err
	}
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		return Policy{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}
