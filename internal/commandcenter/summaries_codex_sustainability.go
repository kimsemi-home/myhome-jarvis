package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

func summarizeCodexSustainability(status codexsustainability.Status) CodexSustainabilitySummary {
	return CodexSustainabilitySummary{
		SustainabilityPosture:                 status.SustainabilityPosture,
		TrendPosture:                          status.TrendPosture,
		EvidenceFreshness:                     status.EvidenceFreshness,
		ReviewGateCount:                       status.ReviewGateCount,
		EstimatedCostUnits:                    status.EstimatedCostUnits,
		AcceptedChangeCount:                   status.AcceptedChangeCount,
		CostPerAcceptedChange:                 status.CostPerAcceptedChange,
		MedianCycleMinutes:                    status.MedianCycleMinutes,
		CacheHitCount:                         status.CacheHitCount,
		CacheMissCount:                        status.CacheMissCount,
		ValidationFailureCount:                status.ValidationFailureCount,
		HumanReviewDebtCount:                  status.HumanReviewDebtCount,
		OptimizationClaimWithoutEvidenceCount: status.OptimizationClaimWithoutEvidenceCount,
	}
}
