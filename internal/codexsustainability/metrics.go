package codexsustainability

func applyMetric(status *Status, record Record) {
	switch record.Metric {
	case "codex_tokens", "codex_coin", "github_actions_minutes":
		status.EstimatedCostUnits += record.Amount
	case "elapsed_cycle_minutes":
		status.cycleMinutes = append(status.cycleMinutes, record.Amount)
	case "rework_count":
		status.ReworkCount += record.Amount
	case "cache_hit_count":
		status.CacheHitCount += record.Amount
	case "cache_miss_count":
		status.CacheMissCount += record.Amount
	case "validation_failure_count":
		status.ValidationFailureCount += record.Amount
	case "human_review_debt":
		status.HumanReviewDebtCount += record.Amount
	case "accepted_change_count":
		status.AcceptedChangeCount += record.Amount
	case "cache_savings_units":
		status.CacheSavingsUnits += record.Amount
	}
}
