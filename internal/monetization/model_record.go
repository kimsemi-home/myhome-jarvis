package monetization

type Record struct {
	At                         string   `json:"at"`
	ExperimentID               string   `json:"experiment_id"`
	HypothesisKey              string   `json:"hypothesis_key"`
	State                      string   `json:"state"`
	DecisionKind               string   `json:"decision_kind"`
	ReviewStatus               string   `json:"review_status"`
	ExpectedValueBand          string   `json:"expected_value_band"`
	CostEstimateUnits          int64    `json:"cost_estimate_units"`
	CostUnitKind               string   `json:"cost_unit_kind"`
	EvidenceRefs               []string `json:"evidence_refs"`
	AuthorityProfile           string   `json:"authority_profile,omitempty"`
	PrivateRevenueNotes        string   `json:"private_revenue_notes,omitempty"`
	PrivateCounterparty        string   `json:"private_counterparty,omitempty"`
	RawRevenueAmountMinorUnits int64    `json:"raw_revenue_amount_minor_units,omitempty"`
}
