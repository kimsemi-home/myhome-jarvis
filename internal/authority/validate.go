package authority

import "fmt"

func validatePolicy(policy Policy) error {
	for _, validate := range []func(Policy) error{
		validatePolicyBasics,
		validatePolicyRoles,
		validatePolicyDecisions,
		validatePolicySummary,
	} {
		if err := validate(policy); err != nil {
			return err
		}
	}
	return nil
}

func validatePolicyBasics(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("authority policy context = %q", policy.Context)
	}
	if !policy.PublicStatusRedacted {
		return fmt.Errorf("authority public status must stay redacted")
	}
	if policy.SelfAuthorityAllowed || policy.ReasoningTierGrantsApproval {
		return fmt.Errorf("authority policy must not allow self-authority or reasoning-tier approval")
	}
	if !policy.PublicRepoHighRiskBlocked {
		return fmt.Errorf("authority policy must block high-risk decisions in public repo mode")
	}
	return requireAll("authority input", normalizeList(policy.RequiredInputs), []string{
		"confidence_assessor", "evidence_quality", "incident_lifecycle",
		"control_plane", "translation", "human_review", "public_safety",
	})
}
