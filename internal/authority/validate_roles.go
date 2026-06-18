package authority

import "fmt"

func validatePolicyRoles(policy Policy) error {
	tiers := mapByKey(policy.ReasoningTiers)
	for _, tier := range []string{"r0_compiler", "r1_low", "r2_medium", "r3_high", "r4_governance"} {
		if _, ok := tiers[tier]; !ok {
			return fmt.Errorf("authority reasoning tier %q is missing", tier)
		}
	}
	roles := mapRolePermissions(policy.RolePermissions)
	for _, role := range []string{
		"producer", "independent_reviewer", "adversarial_reviewer",
		"deterministic_verifier", "governance_steward",
	} {
		if _, ok := roles[role]; !ok {
			return fmt.Errorf("authority role %q is missing", role)
		}
	}
	return requireAll("authority domain attribute", normalizeList(policy.DomainAttributes), []string{
		"agent_reliability", "reasoning_tier", "ontology_maturity",
		"evidence_quality", "security_impact", "data_sensitivity",
		"change_risk", "verification_scope", "lease_status",
		"quarantine_state", "human_review_capacity",
	})
}
