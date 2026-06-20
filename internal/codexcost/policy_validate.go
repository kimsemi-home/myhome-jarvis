package codexcost

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "CodexCostGovernor" {
		return fmt.Errorf("codex cost policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateUsageLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateUsageLedger, ".jsonl") {
		return fmt.Errorf("codex cost ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted ||
		policy.RawUsagePublicAllowed {
		return fmt.Errorf("codex cost policy must be private append-only and redacted")
	}
	if policy.WarningUnitThreshold <= 0 ||
		policy.ReviewUnitThreshold <= policy.WarningUnitThreshold {
		return fmt.Errorf("codex cost thresholds are invalid")
	}
	if err := requireAll("unit kind", policy.UnitKinds, requiredUnitKinds); err != nil {
		return err
	}
	if err := requireAll("loop scope", policy.LoopScopes, requiredLoopScopes); err != nil {
		return err
	}
	if err := requireAll("required field", policy.RequiredFields, requiredRecordFields); err != nil {
		return err
	}
	if err := requireAll("semantic hash input", policy.SemanticHashInputs, requiredSemanticHashInputs); err != nil {
		return err
	}
	if err := requireAll("public summary", policy.PublicSummaryFields, requiredSummaryFields); err != nil {
		return err
	}
	for _, command := range []string{
		"mhj codex-cost status",
		"mhj codex-cost record <json-payload>",
		"mhj codex-cost guard <json-payload>",
		"mhj codex-cost roi",
	} {
		if !contains(policy.Commands, command) {
			return fmt.Errorf("codex cost command %q is missing", command)
		}
	}
	return nil
}

func requireAll(label string, values []string, required []string) error {
	normalized := normalizeList(values)
	for _, value := range required {
		if !contains(normalized, value) {
			return fmt.Errorf("codex cost %s %q is missing", label, value)
		}
	}
	return nil
}
