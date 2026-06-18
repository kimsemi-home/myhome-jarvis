package translation

import "fmt"

func validatePolicySummary(policy Policy) error {
	if err := requireAll("translation public summary", normalizeList(policy.PublicSummaryFields), []string{
		"open_debt_count",
		"forbidden_loss_count",
		"invalid_manifest_count",
		"missing_manifest_count",
		"checked_at",
	}); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj translation status") {
		return fmt.Errorf("translation status command is missing")
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
