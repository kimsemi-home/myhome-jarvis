package commandcenter

func authorityReviewCodexSustainabilityFixture() CodexSustainabilitySummary {
	return CodexSustainabilitySummary{
		PublicSafe:                 true,
		SustainabilityPosture:      "sustainable",
		TrendPosture:               "on_trend",
		EvidenceFreshness:          "fresh",
		RecordCount:                69,
		TrendBaselineCount:         24,
		AcceptedChangeCount:        0,
		CostPerAcceptedChange:      100000,
		MedianCycleMinutes:         2,
		CacheHitCount:              63,
		CacheMissCount:             7,
		CacheSavingsUnits:          410352,
		ValidationFailureCount:     0,
		HumanReviewDebtCount:       0,
		ReworkCount:                0,
		ReviewGateCount:            0,
		LatestTrendBaselineVersion: "quality-run-20260620T174011Z",
	}
}
