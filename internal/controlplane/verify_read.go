package controlplane

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func readVerificationPolicy(root string) (VerificationPolicy, error) {
	path := filepath.Join(root, filepath.FromSlash(VerificationRelativePath))
	body, err := os.ReadFile(path)
	if err != nil {
		return VerificationPolicy{}, err
	}
	var policy VerificationPolicy
	if err := json.Unmarshal(body, &policy); err != nil {
		return VerificationPolicy{}, err
	}
	return policy, nil
}
