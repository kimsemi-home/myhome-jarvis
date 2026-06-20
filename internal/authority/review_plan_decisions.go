package authority

func applyReviewDecisionCounts(policy Policy, status Status, plan *ReviewPlanStatus) {
	blocked := stringSet(status.BlockedDecisions)
	for _, decision := range normalizedDecisions(policy.Decisions) {
		if !blocked[decision.Key] {
			continue
		}
		if decision.Risk == "high" {
			plan.HighRiskBlockedDecisionCount++
		}
		if decision.RequiresHumanReview {
			plan.ReviewRequiredDecisionCount++
			plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "human_review")
		}
	}
	if plan.HighRiskBlockedDecisionCount > 0 {
		plan.RequiredReviewClasses = append(plan.RequiredReviewClasses, "high_risk_public_repo_review")
	}
}

func stringSet(values []string) map[string]bool {
	mapped := map[string]bool{}
	for _, value := range values {
		mapped[value] = true
	}
	return mapped
}
