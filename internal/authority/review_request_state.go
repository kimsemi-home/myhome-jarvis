package authority

func reviewRequestState(plan ReviewPlanStatus) string {
	if plan.ReviewRequestable {
		return "ready"
	}
	return "blocked"
}

func reviewRequestNextHandling(plan ReviewPlanStatus) string {
	if plan.ReviewRequestable {
		return "attach_packet_to_human_review_evidence"
	}
	if plan.NextSafeAction != "" {
		return plan.NextSafeAction
	}
	return "resolve_authority_debt"
}

func reviewRequestEvidenceFields() []string {
	return []string{
		"policy_path",
		"required_review_classes",
		"blocked_decision_counts",
		"profile_gate_counts",
		"non_approval_flags",
	}
}
