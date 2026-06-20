package codexsustainability

type ProposalRecordRequest struct {
	At                    string   `json:"at,omitempty"`
	ProposalID            string   `json:"proposal_id"`
	CostPerAcceptedChange int64    `json:"cost_per_accepted_change"`
	MedianCycleMinutes    int64    `json:"median_cycle_minutes"`
	CacheSavingsUnits     int64    `json:"cache_savings_units"`
	DefectReworkRate      float64  `json:"defect_rework_rate"`
	MonetizationRef       string   `json:"monetization_ref"`
	EvidenceRefs          []string `json:"evidence_refs"`
}

type ProposalRecordResult struct {
	ProposalID            string `json:"proposal_id"`
	CostPerAcceptedChange int64  `json:"cost_per_accepted_change"`
	MedianCycleMinutes    int64  `json:"median_cycle_minutes"`
	CacheSavingsUnits     int64  `json:"cache_savings_units"`
	ReviewState           string `json:"review_state"`
	ReviewGateCount       int    `json:"review_gate_count"`
	EvidenceRefCount      int    `json:"evidence_ref_count"`
	RecordedAt            string `json:"recorded_at"`
}
