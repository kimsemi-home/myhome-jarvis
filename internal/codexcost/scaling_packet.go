package codexcost

func ScalingPacketForRoot(root string) (ScalingPacket, error) {
	brief, err := BriefForRoot(root)
	if err != nil {
		return ScalingPacket{}, err
	}
	return buildScalingPacket(brief), nil
}

func buildScalingPacket(brief Brief) ScalingPacket {
	return ScalingPacket{
		Context:           "CodexCostScalingPacket",
		Version:           "v1",
		PublicSafe:        true,
		Redaction:         "codex-cost-scaling-public-handoff",
		PolicyPath:        brief.PolicyPath,
		Decision:          brief.Decision,
		Recommendation:    brief.Recommendation,
		NextSafeAction:    brief.NextSafeAction,
		BudgetHeadroom:    budgetHeadroom(brief),
		EvidencePosture:   scalingEvidence(brief),
		StorageEvidence:   scalingStorage(brief),
		ScalingOptions:    scalingOptions(brief.Decision),
		CanApplyExpansion: false,
		CheckedAt:         brief.CheckedAt,
	}
}

func budgetHeadroom(brief Brief) BudgetHeadroom {
	return BudgetHeadroom{
		BudgetState:          brief.BudgetState,
		TotalUnits:           brief.TotalUnits,
		WarningUnitThreshold: brief.WarningUnitThreshold,
		ReviewUnitThreshold:  brief.ReviewUnitThreshold,
		RemainingToWarning:   remainingUnits(brief.WarningUnitThreshold, brief.TotalUnits),
		RemainingToReview:    remainingUnits(brief.ReviewUnitThreshold, brief.TotalUnits),
		WarningUsedPercent:   usedPercent(brief.TotalUnits, brief.WarningUnitThreshold),
		ReviewUsedPercent:    usedPercent(brief.TotalUnits, brief.ReviewUnitThreshold),
	}
}

func scalingEvidence(brief Brief) ScalingEvidence {
	return ScalingEvidence{
		AttributionCoveragePercent: brief.AttributionCoveragePercent,
		AcceptedChangeCount:        brief.AcceptedChangeCount,
		CacheSavingsUnits:          brief.CacheSavingsUnits,
		ValueProxyUnits:            brief.ValueProxyUnits,
		CostPerAcceptedChange:      brief.CostPerAcceptedChange,
		SustainabilityPosture:      brief.SustainabilityPosture,
		TrendPosture:               brief.TrendPosture,
		ReviewGateCount:            brief.ReviewGateCount,
	}
}
