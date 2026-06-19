package evidencequality

import (
	"fmt"
	"strings"
)

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("evidence quality policy context = %q", policy.Context)
	}
	if !strings.HasPrefix(policy.PrivateSnapshotLedger, "data/private/") ||
		!strings.HasSuffix(policy.PrivateSnapshotLedger, ".jsonl") {
		return fmt.Errorf("evidence quality ledger must stay in data/private JSONL")
	}
	if !policy.AppendOnly || !policy.PublicStatusRedacted ||
		policy.RawSnapshotPublicAllowed {
		return fmt.Errorf("evidence quality policy must be private append-only with redacted public status")
	}
	if policy.StaleAfterHours <= 0 {
		return fmt.Errorf("evidence quality stale threshold must be positive")
	}
	if err := requireAll("quality level", policy.QualityLevels, requiredQualityLevels); err != nil {
		return err
	}
	if err := requireAll("mapping confidence", policy.MappingConfidenceLevels, requiredMappingLevels); err != nil {
		return err
	}
	if err := requireAll("purpose", policy.AllowedPurposes, requiredPurposes); err != nil {
		return err
	}
	if err := requireAll("reassessment reason", policy.ReassessmentReasons, requiredReassessmentReasons); err != nil {
		return err
	}
	if err := requireAll("required field", policy.RequiredFields, requiredSnapshotFields); err != nil {
		return err
	}
	if err := requireAll("public summary", policy.PublicSummaryFields, requiredSummaryFields); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj evidence-quality status") {
		return fmt.Errorf("evidence quality status command is missing")
	}
	return nil
}

func requireAll(label string, values []string, required []string) error {
	normalized := normalizeList(values)
	for _, value := range required {
		if !contains(normalized, value) {
			return fmt.Errorf("evidence quality %s %q is missing", label, value)
		}
	}
	return nil
}
