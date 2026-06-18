package translation

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	for _, validate := range []func(Policy) error{
		validatePolicyBasics,
		validatePolicyRequiredValues,
		validatePolicySummary,
	} {
		if err := validate(policy); err != nil {
			return err
		}
	}
	return nil
}

func validatePolicyBasics(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("translation policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateLossLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateLossLedger, ".jsonl") {
		return fmt.Errorf("translation loss ledger must stay in data/private JSONL")
	}
	if !strings.HasPrefix(policy.PrivateManifestRoot, "data/private/") {
		return fmt.Errorf("translation manifest root must stay in data/private")
	}
	if !policy.ManifestRequired || !policy.PublicStatusRedacted || policy.RawLossPublicAllowed {
		return fmt.Errorf("translation policy must require private manifests and redacted public status")
	}
	return nil
}
