package monetization

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "MonetizationExperimentLedger" {
		return fmt.Errorf("monetization policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateExperimentLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateExperimentLedger, ".jsonl") {
		return fmt.Errorf("monetization ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted ||
		policy.RawRevenuePublicAllowed {
		return fmt.Errorf("monetization policy must be private append-only and redacted")
	}
	if !policy.DecisionEvidenceRequired || !policy.CostEstimateRequired {
		return fmt.Errorf("monetization decisions require evidence and cost")
	}
	for _, item := range policyRequirements(policy) {
		if err := requireAll(item.label, item.values, item.required); err != nil {
			return err
		}
	}
	if !contains(policy.Commands, "mhj monetization status") {
		return fmt.Errorf("monetization status command is missing")
	}
	if !contains(policy.Commands, "mhj monetization record <json-payload>") {
		return fmt.Errorf("monetization record command is missing")
	}
	return nil
}

type requirement struct {
	label    string
	values   []string
	required []string
}

func policyRequirements(policy Policy) []requirement {
	return []requirement{
		{"state", policy.ExperimentStates, requiredStates},
		{"decision kind", policy.DecisionKinds, requiredDecisions},
		{"review status", policy.ReviewStatuses, requiredReviews},
		{"expected value band", policy.ExpectedValueBands, requiredBands},
		{"cost unit kind", policy.CostUnitKinds, requiredCostUnits},
		{"required field", policy.RequiredFields, requiredFields},
		{"public summary", policy.PublicSummaryFields, requiredSummaryFields},
	}
}
