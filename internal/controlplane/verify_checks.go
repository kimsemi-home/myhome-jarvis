package controlplane

import (
	"encoding/json"
	"fmt"
)

func requireVerifierChecks(verifier VerificationPolicy) error {
	covered := map[string]bool{}
	for _, check := range verifier.Checks {
		if check.ID == "" || check.Evidence == "" {
			return fmt.Errorf("control-plane verifier check requires id and evidence")
		}
		covered[normalizeToken(check.ID)] = true
	}
	return requireAll("control-plane verifier check", coveredChecks(covered), []string{
		"policy-json-valid", "status-public-redacted", "lease-bounds-valid",
		"verifier-separation-required", "manifest-debt-evaluated",
	})
}

func coveredChecks(covered map[string]bool) []string {
	values := make([]string, 0, len(covered))
	for value := range covered {
		values = append(values, value)
	}
	return values
}

func verifyStatus(policy Policy, status Status) error {
	if status.ManifestDebtCount != 0 || status.VerifierViolationCount != 0 {
		return fmt.Errorf("control-plane manifest debt is not clear")
	}
	if !status.VerifierSeparationRequired ||
		status.MinLeaseSeconds <= 0 || status.MaxLeaseSeconds <= status.MinLeaseSeconds {
		return fmt.Errorf("control-plane verifier status is invalid")
	}
	body, err := json.Marshal(status)
	if err != nil {
		return err
	}
	if err := rejectSensitiveText(string(body)); err != nil {
		return err
	}
	return rejectForbiddenPolicyText(policy, string(body))
}
