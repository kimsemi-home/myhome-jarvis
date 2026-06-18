package controlplane

import "fmt"

func validatePolicySummary(policy Policy) error {
	if err := requireAll("control-plane public summary", normalizeList(policy.PublicSummaryFields), []string{
		"count", "invalid_manifest_count", "manifest_debt_count",
		"verifier_violation_count", "by_decision_kind",
		"by_authority_profile", "by_lease_status", "checked_at",
	}); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj control-plane status") {
		return fmt.Errorf("control-plane status command is missing")
	}
	return nil
}

func requireAll(label string, values []string, required []string) error {
	for _, value := range required {
		if !contains(values, value) {
			return fmt.Errorf("%s %q is missing", label, value)
		}
	}
	return nil
}
