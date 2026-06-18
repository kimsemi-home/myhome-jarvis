package authority

import (
	"sort"
	"strings"
)

func normalizedDecisions(decisions []Decision) []Decision {
	normalized := make([]Decision, 0, len(decisions))
	for _, decision := range decisions {
		decision.Key = normalizeToken(decision.Key)
		decision.Risk = normalizeToken(decision.Risk)
		if decision.Key == "" {
			continue
		}
		normalized = append(normalized, decision)
	}
	sort.Slice(normalized, func(i, j int) bool {
		return normalized[i].Key < normalized[j].Key
	})
	return normalized
}

func mapDecisions(decisions []Decision) map[string]Decision {
	mapped := map[string]Decision{}
	for _, decision := range decisions {
		mapped[decision.Key] = decision
	}
	return mapped
}

func mapByKey(tiers []ReasoningTier) map[string]ReasoningTier {
	mapped := map[string]ReasoningTier{}
	for _, tier := range tiers {
		tier.Key = normalizeToken(tier.Key)
		if tier.Key != "" {
			mapped[tier.Key] = tier
		}
	}
	return mapped
}

func mapRolePermissions(roles []RolePermission) map[string]RolePermission {
	mapped := map[string]RolePermission{}
	for _, role := range roles {
		role.Role = normalizeToken(role.Role)
		if role.Role != "" {
			mapped[role.Role] = role
		}
	}
	return mapped
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}
