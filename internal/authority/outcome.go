package authority

func authorityOutcome(status Status, inputs Inputs) (string, string) {
	if !status.PublicSafetyOK {
		return "blocked", "public_safety_not_ok"
	}
	if inputs.Confidence.Blocked {
		return "blocked", "confidence_blocked"
	}
	if status.ConfidenceCap == "low" || status.ConfidenceCap == "unknown" {
		return "blocked", "confidence_cap_low"
	}
	if inputs.Translation.ForbiddenLossCount > 0 {
		return "blocked", "forbidden_translation_loss"
	}
	if status.HumanReviewCapacityState == "overloaded" {
		return "review_required", "human_review_overloaded"
	}
	if status.AuthorityDebtCount > 0 {
		return debtOutcome(status)
	}
	return "limited", "public_repo_high_risk_block"
}

func debtOutcome(status Status) (string, string) {
	switch {
	case status.EvidenceQualityDebtCount > 0:
		return "review_required", "evidence_quality_debt"
	case status.IncidentDebtCount > 0:
		return "review_required", "incident_debt"
	case status.ControlPlaneDebtCount > 0:
		return "review_required", "control_plane_debt"
	case status.TranslationDebtCount > 0:
		return "review_required", "translation_debt"
	default:
		return "review_required", "human_review_debt"
	}
}

func decisionAllowed(outcome string, decision Decision) bool {
	switch outcome {
	case "blocked":
		return decision.AllowedWhenBlocked
	case "review_required":
		return decision.AllowedWhenBlocked ||
			(decision.PublicRepoAllowed && !decision.RequiresHumanReview && decision.Risk == "low")
	default:
		return decision.PublicRepoAllowed && decision.Risk != "high"
	}
}
