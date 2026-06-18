package controlplane

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	for _, validate := range []func(Policy) error{
		validatePolicyBasics,
		validatePolicyLists,
		validatePolicySummary,
	} {
		if err := validate(policy); err != nil {
			return err
		}
	}
	return nil
}

func validatePolicyBasics(policy Policy) error {
	if policy.Context != "AgentOps" {
		return fmt.Errorf("control-plane policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateManifestLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateManifestLedger, ".jsonl") {
		return fmt.Errorf("control-plane manifest ledger must stay in data/private JSONL")
	}
	if !policy.ManifestRequired || !policy.AppendOnly ||
		!policy.PublicStatusRedacted || policy.RawRationalePublicAllowed {
		return fmt.Errorf("control-plane policy must require private append-only redacted manifests")
	}
	if !policy.VerifierSeparationRequired {
		return fmt.Errorf("control-plane verifier separation must be required")
	}
	if policy.MinLeaseSeconds <= 0 || policy.MaxLeaseSeconds <= policy.MinLeaseSeconds {
		return fmt.Errorf("control-plane lease bounds are invalid")
	}
	return nil
}
