package evidence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadPolicy(root string) (Policy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(PolicyRelativePath)))
	if err != nil {
		return Policy{}, err
	}
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		return Policy{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("evidence graph policy context = %q", policy.Context)
	}
	if policy.PrivateRoot != "data/private" {
		return fmt.Errorf("evidence graph private root must be data/private")
	}
	if !policy.PrivateGraphRequired ||
		!policy.PublicStatusRedacted ||
		policy.RawEvidencePublicAllowed {
		return fmt.Errorf("evidence graph policy must require private graph evidence and redacted public status")
	}
	if err := validateGraphKinds(policy); err != nil {
		return err
	}
	if err := validatePrivateSources(policy); err != nil {
		return err
	}
	return validatePolicySummary(policy)
}
