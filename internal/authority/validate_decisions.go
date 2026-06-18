package authority

import "fmt"

func validatePolicyDecisions(policy Policy) error {
	decisions := normalizedDecisions(policy.Decisions)
	if len(decisions) == 0 {
		return fmt.Errorf("authority decisions are required")
	}
	decisionMap := mapDecisions(decisions)
	for _, key := range []string{
		"major_ontology_change", "security_boundary_change", "production_change",
		"evidence_pruning", "quarantine_release", "high_risk_automation",
	} {
		decision, ok := decisionMap[key]
		if !ok {
			return fmt.Errorf("authority high-risk decision %q is missing", key)
		}
		if decision.PublicRepoAllowed || decision.Risk != "high" {
			return fmt.Errorf("authority high-risk decision %q must stay blocked", key)
		}
	}
	return validateDecisionRisks(decisions)
}

func validateDecisionRisks(decisions []Decision) error {
	for _, decision := range decisions {
		if decision.Key == "" {
			return fmt.Errorf("authority decision key is required")
		}
		if !contains([]string{"low", "medium", "high"}, decision.Risk) {
			return fmt.Errorf("authority decision %q has invalid risk", decision.Key)
		}
		if decision.Risk == "high" && decision.PublicRepoAllowed {
			return fmt.Errorf("authority high-risk decision %q must not be public-repo allowed", decision.Key)
		}
	}
	return nil
}
