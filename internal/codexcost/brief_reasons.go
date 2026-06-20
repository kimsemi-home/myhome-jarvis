package codexcost

func briefReasons(roi ROISummary) []string {
	reasons := []string{}
	if roi.BudgetState == "review_required" {
		reasons = append(reasons, "current_review_required")
	} else if roi.BudgetState == "warning" {
		reasons = append(reasons, "current_warning")
	}
	if roi.ReviewGateCount > 0 {
		reasons = append(reasons, "sustainability_review_gate")
	}
	switch roi.SustainabilityPosture {
	case "review_required":
		reasons = append(reasons, "sustainability_review_required")
	case "blocked":
		reasons = append(reasons, "sustainability_blocked")
	}
	if !roi.StorageArchiveReady {
		reasons = append(reasons, "storage_archive_not_ready")
	}
	if !roi.NoiseBudgetReady {
		reasons = append(reasons, "noise_budget_not_ready")
	}
	if !roi.ConfigIsEvidence {
		reasons = append(reasons, "archive_config_not_evidence")
	}
	if roi.ArchiveManifestBudgetBreaches > 0 {
		reasons = append(reasons, "archive_noise_budget_breach")
	}
	if roi.ArchiveManifestInvalidEntries > 0 {
		reasons = append(reasons, "archive_manifest_invalid")
	}
	if roi.InvalidAttributionCount > 0 {
		reasons = append(reasons, "invalid_attribution_records")
	}
	if roi.TotalUnits > 0 && roi.AttributionCoveragePercent < 80 {
		reasons = append(reasons, "attribution_coverage_low")
	}
	return reasons
}
