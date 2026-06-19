package incidents

import "fmt"

func validatePolicyLists(policy Policy) error {
	validations := []struct {
		values []string
		want   []string
		name   string
	}{
		{policy.AllowedKinds, requiredIncidentKinds(), "incident kind"},
		{policy.Lifecycle, requiredLifecycleStages(), "incident lifecycle stage"},
		{policy.AllowedStatuses, requiredStatuses(), "incident status"},
		{policy.OwnerRoles, requiredOwnerRoles(), "incident owner role"},
		{policy.QuarantineStates, requiredQuarantineStates(), "incident quarantine state"},
		{policy.RequiredFields, requiredIncidentFields(), "incident required field"},
		{policy.PublicSummaryFields, requiredSummaryFields(), "incident public summary"},
	}
	for _, validation := range validations {
		if err := requirePolicyValues(validation.values, validation.want, validation.name); err != nil {
			return err
		}
	}
	if !contains(policy.Commands, "mhj incidents status") {
		return fmt.Errorf("incident status command is missing")
	}
	return nil
}

func requirePolicyValues(values []string, required []string, name string) error {
	normalized := normalizeList(values)
	for _, item := range required {
		if !contains(normalized, item) {
			return fmt.Errorf("%s %q is missing", name, item)
		}
	}
	return nil
}
