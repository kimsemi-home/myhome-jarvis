package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/codexcost"

func summarizeCodexCostScaling(
	packet codexcost.ScalingPacket,
) CodexCostScalingSummary {
	granting := grantingScalingOptionCount(packet.ScalingOptions)
	return CodexCostScalingSummary{
		PublicSafe:                    packet.PublicSafe && !packet.CanApplyExpansion && granting == 0,
		Decision:                      packet.Decision,
		Recommendation:                packet.Recommendation,
		NextSafeAction:                packet.NextSafeAction,
		CanApplyExpansion:             packet.CanApplyExpansion,
		BudgetState:                   packet.BudgetHeadroom.BudgetState,
		TotalUnits:                    packet.BudgetHeadroom.TotalUnits,
		RemainingToWarningUnits:       packet.BudgetHeadroom.RemainingToWarning,
		RemainingToReviewUnits:        packet.BudgetHeadroom.RemainingToReview,
		WarningUsedPercent:            packet.BudgetHeadroom.WarningUsedPercent,
		ReviewUsedPercent:             packet.BudgetHeadroom.ReviewUsedPercent,
		AttributionCoveragePercent:    packet.EvidencePosture.AttributionCoveragePercent,
		AcceptedChangeCount:           packet.EvidencePosture.AcceptedChangeCount,
		CacheSavingsUnits:             packet.EvidencePosture.CacheSavingsUnits,
		ValueProxyUnits:               packet.EvidencePosture.ValueProxyUnits,
		CostPerAcceptedChange:         packet.EvidencePosture.CostPerAcceptedChange,
		SustainabilityPosture:         packet.EvidencePosture.SustainabilityPosture,
		TrendPosture:                  packet.EvidencePosture.TrendPosture,
		ReviewGateCount:               packet.EvidencePosture.ReviewGateCount,
		StorageArchivePattern:         packet.StorageEvidence.Pattern,
		StorageArchiveReady:           packet.StorageEvidence.Ready,
		NoiseBudgetReady:              packet.StorageEvidence.NoiseBudgetReady,
		ArchiveManifestBudgetBreaches: packet.StorageEvidence.ManifestBudgetBreaches,
		ArchiveManifestInvalidEntries: packet.StorageEvidence.ManifestInvalidEntries,
		ConfigIsEvidence:              packet.StorageEvidence.ConfigIsEvidence,
		ScalingOptionCount:            len(packet.ScalingOptions),
		RecommendedScalingOptionKeys:  recommendedScalingOptionKeys(packet.ScalingOptions),
		GrantingScalingOptionCount:    granting,
	}
}
