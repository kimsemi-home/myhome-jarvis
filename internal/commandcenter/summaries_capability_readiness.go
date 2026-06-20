package commandcenter

func summarizeCapabilityReadiness(status Status) CapabilityReadinessSummary {
	return CapabilityReadinessSummary{
		PublicSafe:             capabilityReadinessPublicSafe(status),
		CapabilityCount:        status.Vision.CapabilityCount,
		ReadyCapabilityCount:   status.Vision.ReadyPillarCount,
		GatedCapabilityCount:   status.Vision.GatedPillarCount,
		BlockedCapabilityCount: status.Vision.BlockedPillarCount,
		ReadyCapabilityKeys:    append([]string{}, status.Vision.ReadyPillarKeys...),
		GatedCapabilityKeys:    append([]string{}, status.Vision.GatedPillarKeys...),
		BlockedCapabilityKeys:  append([]string{}, status.Vision.BlockedPillarKeys...),
		Media:                  summarizeCapabilityMedia(status),
		FinanceConsent:         summarizeCapabilityFinance(status),
		Monetization:           summarizeCapabilityMonetization(status),
		CodexCost:              summarizeCapabilityCodexCost(status),
	}
}

func capabilityReadinessPublicSafe(status Status) bool {
	return status.MediaReadiness.PublicSafe &&
		status.CodexCostBrief.PublicSafe &&
		status.CodexCostScaling.PublicSafe &&
		!status.CodexCostScaling.CanApplyExpansion &&
		status.CodexCostScaling.GrantingScalingOptionCount == 0 &&
		status.FinanceConsent.ForbiddenActionEnabledCount == 0
}

func capabilityState(key string, status Status) string {
	return visionPillarReadiness(key, status)
}
