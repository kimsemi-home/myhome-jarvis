package authority

import "fmt"

func validatePolicySummary(policy Policy) error {
	if err := requireAll("authority outcome", normalizeList(policy.Outcomes), []string{
		"limited", "review_required", "blocked",
	}); err != nil {
		return err
	}
	if err := requireAll("authority debt class", normalizeList(policy.AuthorityDebtClasses), []string{
		"public_safety", "confidence_cap", "evidence_quality",
		"incident", "control_plane", "translation", "human_review",
	}); err != nil {
		return err
	}
	if err := requireAll("authority public summary", normalizeList(policy.PublicSummaryFields), []string{
		"outcome", "active_rule", "allowed_decision_count",
		"blocked_decision_count", "authority_debt_count", "public_repo_mode",
		"reasoning_tier_grants_approval", "self_authority_allowed",
		"public_safety_ok", "confidence_cap", "human_review_debt_count",
		"human_review_capacity_state", "allowed_decisions",
		"blocked_decisions", "profile_count", "review_required_profile_count",
		"public_safety_gated_profile_count", "self_approval_blocked_profile_count",
		"profile_keys", "review_required_profiles", "checked_at",
	}); err != nil {
		return err
	}
	if !contains(policy.Commands, "mhj authority status") {
		return fmt.Errorf("authority status command is missing")
	}
	if !contains(policy.Commands, "mhj authority-review status") {
		return fmt.Errorf("authority review status command is missing")
	}
	return nil
}

func requireAll(label string, values []string, required []string) error {
	for _, value := range required {
		if !contains(values, value) {
			return fmt.Errorf("%s %q is missing", label, value)
		}
	}
	return nil
}
