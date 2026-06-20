package codexsustainability

func reviewGateCount(policy Policy, status Status) int {
	count := 0
	if status.EvidenceFreshness != "fresh" {
		count++
	}
	if status.TrendPosture == "missing" || status.TrendPosture == "stale" ||
		status.TrendPosture == "slower_than_trend" {
		count++
	}
	if status.EstimatedCostUnits > 0 && status.AcceptedChangeCount == 0 {
		count++
	}
	if status.CostPerAcceptedChange >
		policy.CostPerAcceptedChangeReviewThreshold {
		count++
	}
	if status.MissingEvidenceCount > 0 {
		count += status.MissingEvidenceCount
	}
	if status.OptimizationClaimWithoutEvidenceCount > 0 {
		count += status.OptimizationClaimWithoutEvidenceCount
	}
	if status.ValidationFailureCount+status.HumanReviewDebtCount+status.ReworkCount > 0 {
		count++
	}
	return count
}
