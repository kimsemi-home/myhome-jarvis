package codexsustainability

import "time"

func finalizeStatus(policy Policy, status Status, now time.Time) Status {
	status.MedianCycleMinutes = median(status.cycleMinutes)
	if status.AcceptedChangeCount > 0 {
		status.CostPerAcceptedChange = status.EstimatedCostUnits / status.AcceptedChangeCount
	}
	if status.maxProposalCostPerAcceptedChange > status.CostPerAcceptedChange {
		status.CostPerAcceptedChange = status.maxProposalCostPerAcceptedChange
	}
	status.EvidenceFreshness = ageState(now, status.LastObservedAt, policy.EvidenceMaxAgeHours)
	status.TrendPosture = trendPosture(policy, status, now)
	status.ReviewGateCount = reviewGateCount(policy, status)
	status.SustainabilityPosture = sustainabilityPosture(status)
	return status
}

func sustainabilityPosture(status Status) string {
	if !status.Exists || status.EvidenceFreshness == "missing" ||
		status.TrendPosture == "missing" {
		return "blocked"
	}
	if status.ReviewGateCount > 0 {
		return "review_required"
	}
	return "sustainable"
}
