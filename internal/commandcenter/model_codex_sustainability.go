package commandcenter

type CodexSustainabilitySummary struct {
	SustainabilityPosture                 string `json:"sustainability_posture"`
	TrendPosture                          string `json:"trend_posture"`
	EvidenceFreshness                     string `json:"evidence_freshness"`
	ReviewGateCount                       int    `json:"review_gate_count"`
	EstimatedCostUnits                    int64  `json:"estimated_cost_units"`
	AcceptedChangeCount                   int64  `json:"accepted_change_count"`
	CostPerAcceptedChange                 int64  `json:"cost_per_accepted_change"`
	MedianCycleMinutes                    int64  `json:"median_cycle_minutes"`
	CacheHitCount                         int64  `json:"cache_hit_count"`
	CacheMissCount                        int64  `json:"cache_miss_count"`
	ValidationFailureCount                int64  `json:"validation_failure_count"`
	HumanReviewDebtCount                  int64  `json:"human_review_debt_count"`
	OptimizationClaimWithoutEvidenceCount int    `json:"optimization_claim_without_evidence_count"`
}
