package confidence

import "fmt"

func validateCapRules(rules []CapRule, levels []string) error {
	if len(rules) == 0 {
		return fmt.Errorf("confidence cap rules are required")
	}
	for _, rule := range rules {
		if normalizeToken(rule.Key) == "" || normalizeToken(rule.When) == "" {
			return fmt.Errorf("confidence cap rule key and condition are required")
		}
		if !contains(levels, normalizeToken(rule.Cap)) {
			return fmt.Errorf("confidence cap rule %q has invalid cap", rule.Key)
		}
	}
	return nil
}

func validatePolicySurface(policy Policy) error {
	summary := normalizeList(policy.PublicSummaryFields)
	for _, field := range requiredSummaryFields {
		if !contains(summary, field) {
			return fmt.Errorf("confidence public summary missing %q", field)
		}
	}
	if !contains(policy.Commands, "mhj confidence status") {
		return fmt.Errorf("confidence status command is missing")
	}
	return nil
}

func requireValues(name string, values []string, required []string) error {
	for _, item := range required {
		if !contains(values, item) {
			return fmt.Errorf("%s %q is missing", name, item)
		}
	}
	return nil
}

var requiredSummaryFields = []string{
	"level_cap",
	"self_report_allowed",
	"evidence_link_count",
	"public_safety_ok",
	"checked_at",
}
