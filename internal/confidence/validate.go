package confidence

import "fmt"

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("confidence policy context = %q", policy.Context)
	}
	if policy.AssessorKey != "confidence_assessor" {
		return fmt.Errorf("confidence assessor key must be confidence_assessor")
	}
	if !policy.ConfidenceIsCap || policy.SelfReportAllowed {
		return fmt.Errorf("confidence must be an external cap and self-reporting must stay disabled")
	}
	if !policy.PublicStatusRedacted || policy.RawEvidencePublicAllowed {
		return fmt.Errorf("confidence public status must be redacted")
	}
	levels := normalizeList(policy.Levels)
	if err := requireValues("confidence level", levels, requiredLevels); err != nil {
		return err
	}
	if err := requireValues("confidence input", normalizeList(policy.Inputs), requiredInputs); err != nil {
		return err
	}
	if err := validateCapRules(policy.CapRules, levels); err != nil {
		return err
	}
	return validatePolicySurface(policy)
}

var requiredLevels = []string{"blocked", "low", "medium", "high"}

var requiredInputs = []string{
	"evidence_graph",
	"learning_ledger",
	"quality_gate",
	"public_safety",
}
