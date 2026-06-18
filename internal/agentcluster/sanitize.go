package agentcluster

import (
	"sort"
	"strings"
)

func sanitizeRoles(roles []Role) []Role {
	clean := make([]Role, 0, len(roles))
	for _, role := range roles {
		role.Key = normalizeToken(role.Key)
		role.Label = publicText(role.Label)
		role.ReasoningTier = strings.TrimSpace(role.ReasoningTier)
		role.Authority = normalizeToken(role.Authority)
		role.MustProduce = normalizeList(role.MustProduce)
		role.MustNot = normalizeList(role.MustNot)
		if role.Key == "" || role.Label == "" {
			continue
		}
		clean = append(clean, role)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func sanitizeSidecars(sidecars []Sidecar) []Sidecar {
	clean := make([]Sidecar, 0, len(sidecars))
	for _, sidecar := range sidecars {
		sidecar.Key = normalizeToken(sidecar.Key)
		sidecar.Label = publicText(sidecar.Label)
		sidecar.Checks = normalizeList(sidecar.Checks)
		if sidecar.Key == "" || sidecar.Label == "" {
			continue
		}
		clean = append(clean, sidecar)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func sanitizeSignals(signals []Signal) []Signal {
	clean := make([]Signal, 0, len(signals))
	for _, signal := range signals {
		signal.Key = normalizeToken(signal.Key)
		signal.Label = publicText(signal.Label)
		signal.Status = normalizeToken(signal.Status)
		signal.Evidence = publicText(signal.Evidence)
		if signal.Key == "" || signal.Label == "" {
			continue
		}
		clean = append(clean, signal)
	}
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Key < clean[j].Key
	})
	return clean
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func publicText(value string) string {
	value = strings.TrimSpace(value)
	return strings.ReplaceAll(value, "\n", " ")
}
