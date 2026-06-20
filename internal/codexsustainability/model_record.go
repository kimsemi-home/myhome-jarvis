package codexsustainability

type Record struct {
	At                       string   `json:"at"`
	RecordKind               string   `json:"record_kind"`
	Metric                   string   `json:"metric,omitempty"`
	Amount                   int64    `json:"amount,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs"`
	TrendBaselineVersion     string   `json:"trend_baseline_version,omitempty"`
	TrendMeasuredAt          string   `json:"trend_measured_at,omitempty"`
	ProposalID               string   `json:"proposal_id,omitempty"`
	CostPerAcceptedChange    int64    `json:"cost_per_accepted_change,omitempty"`
	MedianCycleMinutes       int64    `json:"median_cycle_minutes,omitempty"`
	CacheSavingsUnits        int64    `json:"cache_savings_units,omitempty"`
	DefectReworkRate         float64  `json:"defect_rework_rate,omitempty"`
	MonetizationRef          string   `json:"monetization_ref,omitempty"`
	RawPrompt                string   `json:"raw_prompt,omitempty"`
	RawTranscript            string   `json:"raw_transcript,omitempty"`
	PrivateNotes             string   `json:"private_notes,omitempty"`
	LocalAbsolutePath        string   `json:"local_absolute_path,omitempty"`
	PrivateFinanceData       string   `json:"private_finance_data,omitempty"`
	UnpublishedRevenueDetail string   `json:"unpublished_revenue_detail,omitempty"`
}
