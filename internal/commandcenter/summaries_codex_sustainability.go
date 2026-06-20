package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

func summarizeCodexSustainability(status codexsustainability.Status) CodexSustainabilitySummary {
	return CodexSustainabilitySummary{
		PublicSafe:                            true,
		SustainabilityPosture:                 status.SustainabilityPosture,
		TrendPosture:                          status.TrendPosture,
		EvidenceFreshness:                     status.EvidenceFreshness,
		ReviewGateCount:                       status.ReviewGateCount,
		RecordCount:                           status.RecordCount,
		TrendBaselineCount:                    status.TrendBaselineCount,
		EstimatedCostUnits:                    status.EstimatedCostUnits,
		AcceptedChangeCount:                   status.AcceptedChangeCount,
		CostPerAcceptedChange:                 status.CostPerAcceptedChange,
		MedianCycleMinutes:                    status.MedianCycleMinutes,
		CacheHitCount:                         status.CacheHitCount,
		CacheMissCount:                        status.CacheMissCount,
		CacheSavingsUnits:                     status.CacheSavingsUnits,
		ValidationFailureCount:                status.ValidationFailureCount,
		HumanReviewDebtCount:                  status.HumanReviewDebtCount,
		ReworkCount:                           status.ReworkCount,
		OptimizationClaimWithoutEvidenceCount: status.OptimizationClaimWithoutEvidenceCount,
		LatestTrendBaselineVersion:            status.LatestTrendBaselineVersion,
	}
}
