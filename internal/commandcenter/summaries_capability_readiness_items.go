package commandcenter

func summarizeCapabilityMedia(status Status) CapabilityMediaReadiness {
	media := status.MediaReadiness
	return CapabilityMediaReadiness{
		State:                   capabilityState("local_media_concierge", status),
		PublicSafe:              media.PublicSafe,
		PlaybackReady:           media.PlaybackReady,
		PlaybackAvailableCount:  media.PlaybackAvailableCount,
		DegradedCount:           media.DegradedCount,
		MaxPlanningLatencyMS:    media.MaxPlanningLatencyMS,
		TargetPlanningLatencyMS: media.TargetPlanningLatencyMS,
		LocalLauncherAvailable:  media.LocalLauncherAvailable,
	}
}

func summarizeCapabilityFinance(status Status) CapabilityFinanceReadiness {
	finance := status.FinanceConsent
	return CapabilityFinanceReadiness{
		State:                       capabilityState("household_finance_copilot", status),
		ReadinessState:              finance.ReadinessState,
		FinanceMode:                 finance.FinanceMode,
		ActiveConsentCount:          finance.ActiveConsentCount,
		MissingRequiredConsentCount: finance.MissingRequiredConsentCount,
		ReviewRequiredCount:         finance.ReviewRequiredCount,
		MissingEvidenceCount:        finance.MissingEvidenceCount,
		ForbiddenActionEnabledCount: finance.ForbiddenActionEnabledCount,
		ConsentDebtCount:            finance.ConsentDebtCount,
	}
}

func summarizeCapabilityMonetization(status Status) CapabilityMonetizationReadiness {
	monetization := status.Monetization
	return CapabilityMonetizationReadiness{
		State:                     capabilityState("monetization_console", status),
		ExperimentCount:           monetization.ExperimentCount,
		DecisionCount:             monetization.DecisionCount,
		ReviewRequiredCount:       monetization.ReviewRequiredCount,
		MissingEvidenceCount:      monetization.MissingEvidenceCount,
		MissingCostEstimateCount:  monetization.MissingCostEstimateCount,
		ExpectedValueUnknownCount: monetization.ExpectedValueUnknownCount,
		MonetizationDebtCount:     monetization.MonetizationDebtCount,
	}
}

func summarizeCapabilityCodexCost(status Status) CapabilityCodexCostReadiness {
	return CapabilityCodexCostReadiness{
		State:                      capabilityState("codex_cost_governor", status),
		PublicSafe:                 status.CodexCostBrief.PublicSafe,
		ScalingPublicSafe:          status.CodexCostScaling.PublicSafe,
		BudgetState:                status.Cost.BudgetState,
		TotalUnits:                 status.Cost.TotalUnits,
		ReviewRequiredCount:        status.Cost.ReviewRequiredCount,
		MissingEvidenceCount:       status.Cost.MissingEvidenceCount,
		BriefDecision:              status.CodexCostBrief.Decision,
		BriefNextSafeAction:        status.CodexCostBrief.NextSafeAction,
		CanApplyExpansion:          status.CodexCostScaling.CanApplyExpansion,
		ReviewGateCount:            status.CodexCostScaling.ReviewGateCount,
		GrantingScalingOptionCount: status.CodexCostScaling.GrantingScalingOptionCount,
	}
}
