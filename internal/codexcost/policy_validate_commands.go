package codexcost

import "fmt"

func validatePolicyCommands(policy Policy) error {
	for _, command := range []string{
		"mhj codex-cost status",
		"mhj codex-cost record <json-payload>",
		"mhj codex-cost guard <json-payload>",
		"mhj codex-cost attribute <json-payload>",
		"mhj codex-cost roi",
		"mhj codex-cost brief",
		"mhj codex-cost scaling-packet",
	} {
		if !contains(policy.Commands, command) {
			return fmt.Errorf("codex cost command %q is missing", command)
		}
	}
	return nil
}
