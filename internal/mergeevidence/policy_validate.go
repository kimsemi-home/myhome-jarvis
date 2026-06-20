package mergeevidence

import "fmt"

var requiredGates = []string{
	"clean_branch", "required_checks_success", "public_safety_passed",
	"review_gate_clear", "generated_drift_clear",
}

var requiredEvidence = []string{
	"pr_url", "feature_commit", "merge_commit", "push_quality_run",
	"pr_quality_run", "main_quality_run", "linear_completion_comment",
	"public_safety_scan",
}

var requiredSummaryFields = []string{
	"context", "version", "policy_path", "default_behavior",
	"eligible_gate_count", "required_evidence_count",
	"missing_required_evidence_count", "merge_ready",
	"merge_blocked_until_evidence", "checked_at",
}

func ValidatePolicy(policy Policy) error {
	if policy.Context != "MergeEvidencePolicy" {
		return fmt.Errorf("merge evidence context = %q", policy.Context)
	}
	if policy.DefaultBehavior != "merge_when_eligible" || !policy.PublicStatusRedacted {
		return fmt.Errorf("merge evidence default behavior or redaction is invalid")
	}
	if policy.MergeWithoutReviewAllowed || policy.PersistPrivateEvidence {
		return fmt.Errorf("merge evidence must not allow unreviewed private evidence")
	}
	if err := validateGates(policy.Gates); err != nil {
		return err
	}
	if missing := missingStrings(policy.RequiredEvidence, requiredEvidence); len(missing) > 0 {
		return fmt.Errorf("merge evidence required item %q is missing", missing[0])
	}
	if missing := missingStrings(policy.PublicSummaryFields, requiredSummaryFields); len(missing) > 0 {
		return fmt.Errorf("merge evidence summary field %q is missing", missing[0])
	}
	if !contains(policy.Commands, "mhj merge-evidence status") {
		return fmt.Errorf("merge evidence status command is missing")
	}
	return nil
}

func validateGates(gates []Gate) error {
	if missing := missingStrings(gateKeys(gates), requiredGates); len(missing) > 0 {
		return fmt.Errorf("merge evidence gate %q is missing", missing[0])
	}
	for _, gate := range gates {
		if gate.Key == "" || gate.Label == "" || gate.Evidence == "" {
			return fmt.Errorf("merge evidence gate %q is incomplete", gate.Key)
		}
		if !gate.Required || !gate.BlocksMerge {
			return fmt.Errorf("merge evidence gate %q must be required and blocking", gate.Key)
		}
	}
	return nil
}
