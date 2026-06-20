package monetization

type Status struct {
	PolicyPath                string         `json:"policy_path"`
	LedgerPath                string         `json:"ledger_path"`
	Exists                    bool           `json:"exists"`
	ExperimentCount           int            `json:"experiment_count"`
	DecisionCount             int            `json:"decision_count"`
	InvalidRecordCount        int            `json:"invalid_record_count"`
	ReviewRequiredCount       int            `json:"review_required_count"`
	MissingEvidenceCount      int            `json:"missing_evidence_count"`
	MissingCostEstimateCount  int            `json:"missing_cost_estimate_count"`
	ExpectedValueUnknownCount int            `json:"expected_value_unknown_count"`
	MonetizationDebtCount     int            `json:"monetization_debt_count"`
	ByState                   map[string]int `json:"by_state"`
	ByDecisionKind            map[string]int `json:"by_decision_kind"`
	ByReviewStatus            map[string]int `json:"by_review_status"`
	ByExpectedValueBand       map[string]int `json:"by_expected_value_band"`
	LastObservedAt            string         `json:"last_observed_at,omitempty"`
	CheckedAt                 string         `json:"checked_at"`
}
