package incidents

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
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("incident policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateIncidentLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateIncidentLedger, ".jsonl") {
		return fmt.Errorf("incident ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted || policy.RawIncidentPublicAllowed {
		return fmt.Errorf("incident policy must be private append-only with redacted public status")
	}
	if policy.QuarantineStaleAfterHours <= 0 {
		return fmt.Errorf("incident quarantine stale threshold must be positive")
	}
	return validatePolicyLists(policy)
}
