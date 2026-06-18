package review

import "fmt"

func validatePolicySets(policy Policy) error {
	checks := []struct {
		label string
		have  []string
		want  []string
	}{
		{"risk", policy.AllowedRisks, requiredRisks},
		{"queue class", policy.QueueClasses, requiredQueueClasses},
		{"status", policy.AllowedStatuses, requiredStatuses},
		{"requester role", policy.RequesterRoles, requiredRequesterRoles},
		{"reviewer role", policy.ReviewerRoles, requiredReviewerRoles},
		{"required field", policy.RequiredFields, requiredReviewFields},
	}
	for _, check := range checks {
		if err := requireAll(check.label, check.have, check.want); err != nil {
			return err
		}
	}
	return validatePriorityOrder(policy)
}

func validatePriorityOrder(policy Policy) error {
	priority := normalizeList(policy.PriorityOrder)
	for _, class := range normalizeList(policy.QueueClasses) {
		if !contains(priority, class) {
			return fmt.Errorf("review priority order missing %q", class)
		}
	}
	if len(policy.PriorityOrder) < 2 ||
		normalizeToken(policy.PriorityOrder[0]) != "security_incident" ||
		normalizeToken(policy.PriorityOrder[1]) != "production_incident" {
		return fmt.Errorf("review priority must start with security and production incidents")
	}
	return nil
}

func requireAll(label string, have []string, want []string) error {
	normalized := normalizeList(have)
	for _, item := range want {
		if !contains(normalized, item) {
			return fmt.Errorf("review %s %q is missing", label, item)
		}
	}
	return nil
}
