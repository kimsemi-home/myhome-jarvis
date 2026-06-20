package codexcost

func briefDecision(reasons []string) string {
	for _, reason := range reasons {
		switch reason {
		case "current_review_required",
			"sustainability_review_gate",
			"sustainability_review_required",
			"sustainability_blocked",
			"storage_archive_not_ready",
			"noise_budget_not_ready",
			"archive_config_not_evidence",
			"archive_noise_budget_breach",
			"archive_manifest_invalid":
			return "review_required"
		}
	}
	if len(reasons) > 0 {
		return "warn"
	}
	return "allow"
}

func briefRecommendation(decision string, roi ROISummary) string {
	switch decision {
	case "review_required":
		return "review_governance_before_scaling"
	case "warn":
		return "tighten_cost_evidence_before_scaling"
	default:
		if roi.ValueProxyUnits > roi.TotalUnits {
			return "cache_value_supports_scaling"
		}
		return "continue_monitoring"
	}
}

func briefNextSafeAction(decision string, roi ROISummary) string {
	switch decision {
	case "review_required":
		return "hold_paid_or_external_expansion"
	case "warn":
		return "improve_attribution_and_recheck"
	default:
		if roi.TotalUnits == 0 {
			return "record_first_cost_sample"
		}
		return "continue_local_first_loop"
	}
}
