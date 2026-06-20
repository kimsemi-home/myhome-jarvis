package commandcenter

func blockedGateKeys(gates []GateSummary) []string {
	keys := make([]string, 0, len(gates))
	for _, gate := range gates {
		keys = append(keys, gate.Key)
	}
	return keys
}

func capabilityKeysForGates(gates []string) []string {
	seen := map[string]bool{}
	keys := []string{}
	for _, gate := range gates {
		for _, capability := range capabilitiesForGate(gate) {
			if !seen[capability] {
				seen[capability] = true
				keys = append(keys, capability)
			}
		}
	}
	return keys
}

func capabilitiesForGate(gate string) []string {
	switch gate {
	case "finance_consent":
		return []string{"household_finance_copilot"}
	case "codex_sustainability", "cost":
		return []string{"codex_cost_governor"}
	case "monetization":
		return []string{"monetization_console"}
	case "repo_factory":
		return []string{"shorts_factory_control_plane"}
	default:
		return []string{"self_improvement_loop"}
	}
}

func workItemGuardrails() []string {
	return []string{
		"public_safe_by_default",
		"private_data_stays_private",
		"no_self_approval",
	}
}
