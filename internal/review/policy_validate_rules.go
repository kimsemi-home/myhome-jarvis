package review

import "fmt"

func validateOverloadPolicy(policy Policy) error {
	rules := mapOverloadRules(policy.OverloadPolicy)
	for _, key := range requiredOverloadRules {
		if _, ok := rules[key]; !ok {
			return fmt.Errorf("review overload policy %q is missing", key)
		}
	}
	for _, key := range frozenOverloadRules {
		if rules[key].AllowedWhenOverloaded {
			return fmt.Errorf("review overload policy must freeze %q", key)
		}
	}
	return nil
}

func validatePolicySummary(policy Policy) error {
	if err := requireAll("public summary", policy.PublicSummaryFields, requiredSummaryFields); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj review status") {
		return fmt.Errorf("review status command is missing")
	}
	return nil
}
