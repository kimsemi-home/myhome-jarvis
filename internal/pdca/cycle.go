package pdca

import "strings"

func invalidCycle(policy Policy, cycle Cycle) bool {
	if cycle.CycleID == "" || cycle.At == "" || cycle.Owner == "" {
		return true
	}
	if !contains(policy.AllowedStatuses, cycle.Status) {
		return true
	}
	for _, ref := range []string{cycle.PlanRef, cycle.DoRef, cycle.CheckRef, cycle.ActRef} {
		if !safeRef(ref) {
			return true
		}
	}
	return false
}

func safeRef(ref string) bool {
	if ref == "" || strings.HasPrefix(ref, "/") || strings.Contains(ref, "..") {
		return false
	}
	for _, prefix := range []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", ".github/"} {
		if strings.HasPrefix(ref, prefix) {
			return true
		}
	}
	return false
}
