package mediareadiness

import "fmt"

func ValidatePolicy(policy Policy) error {
	if policy.Context != "MediaReadinessBenchmark" {
		return fmt.Errorf("media readiness context = %q", policy.Context)
	}
	if policy.Version == "" || policy.BenchmarkKind != "command_planning" {
		return fmt.Errorf("media readiness version or benchmark kind is invalid")
	}
	if !policy.PublicStatusRedacted || policy.ExecuteCommands {
		return fmt.Errorf("media readiness must be public-redacted and dry-run")
	}
	if policy.PersistPayloads || policy.PersistURLs {
		return fmt.Errorf("media readiness must not persist payloads or URLs")
	}
	if policy.TargetPlanningLatencyMS <= 0 {
		return fmt.Errorf("media readiness target latency must be positive")
	}
	return validateCases(policy.Cases)
}

func validateCases(cases []BenchmarkCase) error {
	required := map[string]bool{"youtube_launch": false, "youtube_search": false, "ott_netflix": false}
	for _, item := range cases {
		if item.ID == "" || item.Capability == "" || item.Command == "" {
			return fmt.Errorf("media readiness case must declare id, capability, command")
		}
		if _, ok := required[item.ID]; ok {
			required[item.ID] = true
		}
	}
	for key, seen := range required {
		if !seen {
			return fmt.Errorf("media readiness case %q is missing", key)
		}
	}
	return nil
}
