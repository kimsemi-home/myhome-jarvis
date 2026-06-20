package codexcost

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

func evaluateGuard(
	policy Policy,
	guard GuardRequest,
	cost Status,
	sustainability codexsustainability.Status,
) GuardResult {
	projected := budgetState(policy, cost.TotalUnits+guard.EstimatedUnits)
	reasons := guardReasons(cost, projected, sustainability)
	return GuardResult{
		Decision:              guardDecision(reasons),
		Reasons:               reasons,
		Scope:                 guard.Scope,
		UnitKind:              guard.UnitKind,
		EstimatedUnits:        guard.EstimatedUnits,
		EstimatedMinutes:      guard.EstimatedMinutes,
		CurrentBudgetState:    cost.BudgetState,
		ProjectedBudgetState:  projected,
		SustainabilityPosture: sustainability.SustainabilityPosture,
		ReviewGateCount:       sustainability.ReviewGateCount,
		EvidenceRefCount:      len(guard.EvidenceRefs),
	}
}

func guardDecision(reasons []string) string {
	for _, reason := range reasons {
		if reason == "projected_review_threshold" ||
			reason == "current_review_required" ||
			reason == "sustainability_review_required" ||
			reason == "sustainability_blocked" {
			return "review_required"
		}
	}
	if len(reasons) > 0 {
		return "warn"
	}
	return "allow"
}
