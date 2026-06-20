package commandcenter

func applyVisionReadiness(status *Status) {
	for _, key := range status.Vision.PillarKeys {
		switch visionPillarReadiness(key, *status) {
		case "ready":
			status.Vision.ReadyPillarKeys = append(status.Vision.ReadyPillarKeys, key)
		case "blocked":
			status.Vision.BlockedPillarKeys = append(status.Vision.BlockedPillarKeys, key)
		default:
			status.Vision.GatedPillarKeys = append(status.Vision.GatedPillarKeys, key)
		}
	}
	status.Vision.ReadyPillarCount = len(status.Vision.ReadyPillarKeys)
	status.Vision.GatedPillarCount = len(status.Vision.GatedPillarKeys)
	status.Vision.BlockedPillarCount = len(status.Vision.BlockedPillarKeys)
}

func visionPillarReadiness(key string, status Status) string {
	switch key {
	case "local_media_concierge":
		if status.MediaReadiness.PublicSafe && status.MediaReadiness.PlaybackReady &&
			status.MediaReadiness.DegradedCount == 0 {
			return "ready"
		}
		return "blocked"
	case "household_finance_copilot":
		if status.FinanceConsent.ReadinessState == "ready_read_only" {
			return "ready"
		}
		return "blocked"
	case "shorts_factory_control_plane":
		if !status.RepoFactory.PublicSafe || status.RepoFactory.MissingCreationGateCount > 0 {
			return "blocked"
		}
		if status.RepoFactory.RepoCreationBlockedUntilReview {
			return "gated"
		}
		return "ready"
	case "monetization_console":
		if status.Monetization.MonetizationDebtCount > 0 {
			return "blocked"
		}
		if status.Monetization.ExperimentCount == 0 ||
			status.Monetization.ReviewRequiredCount > 0 {
			return "gated"
		}
		return "ready"
	case "codex_cost_governor":
		if status.Cost.BudgetState == "ok" && status.Cost.ReviewRequiredCount == 0 &&
			status.Cost.MissingEvidenceCount == 0 &&
			status.CodexCostBrief.Decision == "allow" {
			return "ready"
		}
		return "gated"
	case "self_improvement_loop":
		if !status.PDCA.Ready || status.Incidents.IncidentDebtCount > 0 ||
			status.LocalRuntime.HealthDebtCount > 0 {
			return "blocked"
		}
		if authorityBlocked(status.Authority.Outcome) || status.Review.ReviewDebtCount > 0 {
			return "gated"
		}
		return "ready"
	default:
		return "gated"
	}
}
