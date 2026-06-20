package authority

import "fmt"

func validatePolicyProfiles(policy Policy) error {
	profiles := normalizedProfiles(policy.AssistantAuthorityProfiles)
	if err := requireAll("assistant authority profile", profileKeys(profiles), []string{
		"local_media_concierge", "household_finance_copilot",
		"shorts_factory_control_plane", "monetization_console",
		"codex_cost_governor", "self_improvement_loop",
	}); err != nil {
		return err
	}
	for _, profile := range profiles {
		if profile.SelfApprovalAllowed {
			return fmt.Errorf("assistant profile %q allows self approval", profile.Key)
		}
	}
	return validateRiskyProfiles(mapProfiles(profiles))
}

func profileKeys(profiles []AssistantProfile) []string {
	keys := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		keys = append(keys, profile.Key)
	}
	return keys
}

func mapProfiles(profiles []AssistantProfile) map[string]AssistantProfile {
	mapped := map[string]AssistantProfile{}
	for _, profile := range profiles {
		mapped[profile.Key] = profile
	}
	return mapped
}
