package codexcost

func BriefForRoot(root string) (Brief, error) {
	roi, err := ROISummaryForRoot(root)
	if err != nil {
		return Brief{}, err
	}
	cost, err := StatusForRoot(root)
	if err != nil {
		return Brief{}, err
	}
	return buildBrief(roi, cost), nil
}

func buildBrief(roi ROISummary, cost Status) Brief {
	reasons := briefReasons(roi)
	decision := briefDecision(reasons)
	return Brief{
		PolicyPath:                      roi.PolicyPath,
		PublicSafe:                      true,
		Decision:                        decision,
		Reasons:                         reasons,
		Recommendation:                  briefRecommendation(decision, roi),
		NextSafeAction:                  briefNextSafeAction(decision, roi),
		BudgetState:                     roi.BudgetState,
		TotalUnits:                      roi.TotalUnits,
		WarningUnitThreshold:            cost.WarningUnitThreshold,
		ReviewUnitThreshold:             cost.ReviewUnitThreshold,
		AttributionCoveragePercent:      roi.AttributionCoveragePercent,
		TrackedScopeCount:               roi.TrackedScopeCount,
		ScopeCount:                      roi.ScopeCount,
		SustainabilityPosture:           roi.SustainabilityPosture,
		TrendPosture:                    roi.TrendPosture,
		ReviewGateCount:                 roi.ReviewGateCount,
		AcceptedChangeCount:             roi.AcceptedChangeCount,
		CacheSavingsUnits:               roi.CacheSavingsUnits,
		ValueProxyUnits:                 roi.ValueProxyUnits,
		CostPerAcceptedChange:           roi.CostPerAcceptedChange,
		StorageArchivePattern:           roi.StorageArchivePattern,
		StorageArchiveReady:             roi.StorageArchiveReady,
		NoiseBudgetReady:                roi.NoiseBudgetReady,
		MaxNoiseRatioPercent:            roi.MaxNoiseRatioPercent,
		ArchiveManifestEntryCount:       roi.ArchiveManifestEntryCount,
		ArchiveManifestBudgetBreaches:   roi.ArchiveManifestBudgetBreaches,
		ArchiveManifestInvalidEntries:   roi.ArchiveManifestInvalidEntries,
		ArchiveManifestCompressionRatio: roi.ArchiveManifestCompressionRatio,
		ConfigIsEvidence:                roi.ConfigIsEvidence,
		ForbiddenPublicFieldCount:       roi.ForbiddenPublicFieldCount,
		CheckedAt:                       roi.CheckedAt,
	}
}
