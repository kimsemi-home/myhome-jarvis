package authority

import "fmt"

func validateRiskyProfiles(profiles map[string]AssistantProfile) error {
	finance := profiles["household_finance_copilot"]
	if !finance.RequiresHumanReview || finance.ExternalWritesAllowed {
		return fmt.Errorf("finance assistant profile must be review-only")
	}
	repo := profiles["shorts_factory_control_plane"]
	if !repo.RequiresHumanReview || !repo.PublicSafetyGateRequired ||
		!repo.PublicRepoChangeGateRequired || !repo.WorkflowChangeGateRequired {
		return fmt.Errorf("repo factory profile must require safety and review gates")
	}
	money := profiles["monetization_console"]
	if !money.RequiresHumanReview || !money.PublicSafetyGateRequired {
		return fmt.Errorf("monetization profile must require review and safety gate")
	}
	loop := profiles["self_improvement_loop"]
	if !loop.VerifierSeparationRequired || !loop.RequiresHumanReview {
		return fmt.Errorf("self-improvement profile must keep verifier separation")
	}
	return nil
}
