package codexcost

import "fmt"

func validateAttributionPolicy(policy Policy) error {
	if err := requireAll(
		"attribution field",
		policy.AttributionRequiredFields,
		requiredAttributionFields,
	); err != nil {
		return err
	}
	if err := requireAll(
		"attribution cost ref input",
		policy.AttributionCostRefInputs,
		requiredAttributionCostRefInputs,
	); err != nil {
		return err
	}
	if policy.AttributionSubjectMaxLength <= 0 ||
		policy.AttributionSubjectMaxLength > 200 {
		return fmt.Errorf("codex cost attribution subject length is invalid")
	}
	if policy.AttributionCostRefMaxLength <= 0 ||
		policy.AttributionCostRefMaxLength > 120 {
		return fmt.Errorf("codex cost attribution cost ref length is invalid")
	}
	return nil
}
