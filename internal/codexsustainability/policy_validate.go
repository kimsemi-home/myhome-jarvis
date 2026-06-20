package codexsustainability

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "CodexSustainabilityEvidenceLoop" {
		return fmt.Errorf("codex sustainability context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateEvidenceLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateEvidenceLedger, ".jsonl") {
		return fmt.Errorf("codex sustainability ledger must stay private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted ||
		policy.RawEvidencePublicAllowed || !policy.TrendBaselinesVersioned {
		return fmt.Errorf("codex sustainability policy must be private and versioned")
	}
	if policy.EvidenceMaxAgeHours <= 0 || policy.TrendBaselineMaxAgeHours <= 0 ||
		policy.CostPerAcceptedChangeReviewThreshold <= 0 {
		return fmt.Errorf("codex sustainability thresholds are invalid")
	}
	for _, item := range policyRequirements(policy) {
		if err := requireAll(item.label, item.values, item.required); err != nil {
			return err
		}
	}
	if !contains(policy.Commands, "mhj codex-sustainability status") {
		return fmt.Errorf("codex sustainability status command is missing")
	}
	return nil
}

type requirement struct {
	label    string
	values   []string
	required []string
}
