package codexsustainability

import "fmt"

func policyRequirements(policy Policy) []requirement {
	return []requirement{
		{"record kind", policy.RecordKinds, requiredRecordKinds},
		{"metric", policy.Metrics, requiredMetrics},
		{"required field", policy.RequiredFields, requiredFields},
		{"proposal field", policy.ProposalRequiredFields, requiredProposalFields},
		{"public summary", policy.PublicSummaryFields, requiredSummaryFields},
	}
}

func requireAll(label string, values []string, required []string) error {
	normalized := normalizeList(values)
	for _, value := range required {
		if !contains(normalized, value) {
			return fmt.Errorf("codex sustainability %s %q is missing", label, value)
		}
	}
	return nil
}
