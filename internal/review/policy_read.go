package review

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	if err := validatePolicyShape(policy); err != nil {
		return err
	}
	if err := validatePolicySets(policy); err != nil {
		return err
	}
	if err := validateOverloadPolicy(policy); err != nil {
		return err
	}
	return validatePolicySummary(policy)
}

func validatePolicyShape(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("review policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateReviewQueue, "data/private/") ||
		!strings.HasSuffix(policy.PrivateReviewQueue, ".jsonl") {
		return fmt.Errorf("review queue must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawReviewPublicAllowed {
		return fmt.Errorf("review policy must be private append-only with redacted public status")
	}
	if policy.MaxOpenReviews <= 0 || policy.MaxHighRiskOpenReviews < 0 ||
		policy.MinBackupReviewers <= 0 {
		return fmt.Errorf("review capacity thresholds are invalid")
	}
	return nil
}
