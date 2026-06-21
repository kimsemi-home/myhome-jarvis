package authority

import "fmt"

func validatePolicyCommands(policy Policy) error {
	for _, command := range requiredPolicyCommands() {
		if !contains(policy.Commands, command) {
			return fmt.Errorf("authority command %q is missing", command)
		}
	}
	return nil
}

func requiredPolicyCommands() []string {
	return []string{
		"mhj authority status",
		"mhj authority-review status",
		"mhj authority-review brief",
		"mhj authority-review request",
		"mhj authority-review evidence",
		"mhj authority-review queue",
		"mhj authority-review record <json-payload>",
		"mhj authority-review approval-status",
		"mhj authority-review approve <json-payload>",
	}
}
