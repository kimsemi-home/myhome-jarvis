package repofactory

import "fmt"

func validatePolicy(policy Policy) error {
	if policy.Context != "PublicSafeRepoFactory" {
		return fmt.Errorf("repo factory policy context = %q", policy.Context)
	}
	if !policy.PublicRepoDefault || !policy.CodexProjectRequired {
		return fmt.Errorf("repo factory must default to public Codex-ready repos")
	}
	if policy.RepoCreationAllowedWithoutReview ||
		policy.PrivateAssetsPublicAllowed ||
		policy.LocalPathsPublicAllowed {
		return fmt.Errorf("repo factory must block unreviewed creation and private public output")
	}
	if !policy.AuthorityReviewRequired || !policy.PublicSafetyEvidenceRequired {
		return fmt.Errorf("repo factory requires authority review and public safety evidence")
	}
	if err := validateTemplateFiles(policy); err != nil {
		return err
	}
	if err := validateCreationGates(policy.CreationGates); err != nil {
		return err
	}
	if missing := containsAll(policy.BootstrapChecklist, requiredChecklistItems); len(missing) > 0 {
		return fmt.Errorf("repo factory bootstrap checklist missing %q", missing[0])
	}
	if missing := containsAll(policy.PublicSummaryFields, requiredSummaryFields); len(missing) > 0 {
		return fmt.Errorf("repo factory public summary field missing %q", missing[0])
	}
	if !contains(policy.Commands, "mhj repo-factory status") {
		return fmt.Errorf("repo factory status command is missing")
	}
	if !contains(policy.Commands, "mhj repo-factory decision-packet") {
		return fmt.Errorf("repo factory decision packet command is missing")
	}
	return nil
}

func validateCreationGates(gates []CreationGate) error {
	if missing := containsAll(gateKeys(gates), requiredCreationGates); len(missing) > 0 {
		return fmt.Errorf("repo factory creation gate missing %q", missing[0])
	}
	for _, gate := range gates {
		if gate.Key == "" || gate.Evidence == "" {
			return fmt.Errorf("repo factory creation gate %q is incomplete", gate.Key)
		}
		if !gate.Required || !gate.BlocksRepoCreation {
			return fmt.Errorf("repo factory creation gate %q must be required and blocking", gate.Key)
		}
	}
	return nil
}
