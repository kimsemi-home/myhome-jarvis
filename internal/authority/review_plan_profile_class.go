package authority

func reviewClassForProfile(profile AssistantProfile) string {
	switch {
	case profile.PublicRepoChangeGateRequired:
		return "public_repo_change_review"
	case profile.WorkflowChangeGateRequired:
		return "workflow_change_review"
	case profile.PublicSafetyGateRequired:
		return "public_safety_review"
	case profile.RequiresHumanReview:
		return "human_review"
	default:
		return "none"
	}
}

func reviewRequestable(policy Policy, status Status, plan ReviewPlanStatus) bool {
	return reviewPlanPublicSafe(policy) &&
		status.PublicSafetyOK &&
		status.Outcome == "limited" &&
		status.ActiveRule == "public_repo_high_risk_block" &&
		plan.ReviewCapacityState == "available" &&
		plan.HighRiskBlockedDecisionCount > 0 &&
		!policy.SelfAuthorityAllowed
}
