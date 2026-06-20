package codexsustainability

import "time"

type Status struct {
	PolicyPath                            string `json:"policy_path"`
	LedgerPath                            string `json:"ledger_path"`
	Exists                                bool   `json:"exists"`
	RecordCount                           int    `json:"record_count"`
	InvalidRecordCount                    int    `json:"invalid_record_count"`
	MissingEvidenceCount                  int    `json:"missing_evidence_count"`
	TrendBaselineCount                    int    `json:"trend_baseline_count"`
	FeatureProposalCount                  int    `json:"feature_proposal_count"`
	OptimizationClaimWithoutEvidenceCount int    `json:"optimization_claim_without_evidence_count"`
	EstimatedCostUnits                    int64  `json:"estimated_cost_units"`
	AcceptedChangeCount                   int64  `json:"accepted_change_count"`
	CostPerAcceptedChange                 int64  `json:"cost_per_accepted_change"`
	MedianCycleMinutes                    int64  `json:"median_cycle_minutes"`
	CacheHitCount                         int64  `json:"cache_hit_count"`
	CacheMissCount                        int64  `json:"cache_miss_count"`
	CacheSavingsUnits                     int64  `json:"cache_savings_units"`
	ValidationFailureCount                int64  `json:"validation_failure_count"`
	HumanReviewDebtCount                  int64  `json:"human_review_debt_count"`
	ReworkCount                           int64  `json:"rework_count"`
	TrendPosture                          string `json:"trend_posture"`
	SustainabilityPosture                 string `json:"sustainability_posture"`
	EvidenceFreshness                     string `json:"evidence_freshness"`
	ReviewGateCount                       int    `json:"review_gate_count"`
	LatestTrendBaselineVersion            string `json:"latest_trend_baseline_version,omitempty"`
	LastObservedAt                        string `json:"last_observed_at,omitempty"`
	CheckedAt                             string `json:"checked_at"`
	cycleMinutes                          []int64
	latestTrendAt                         string
	trendBaselineCycleMinutes             int64
	maxProposalCostPerAcceptedChange      int64
}

func newStatus(policy Policy, checkedAt time.Time) Status {
	return Status{PolicyPath: PolicyRelativePath, LedgerPath: policy.PrivateEvidenceLedger,
		TrendPosture: "missing", SustainabilityPosture: "blocked",
		EvidenceFreshness: "missing", CheckedAt: checkedAt.Format(time.RFC3339)}
}
