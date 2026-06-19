package domain

type StoragePolicy struct {
	FixtureFormat  string   `json:"fixture_format"`
	LakeLayers     []string `json:"lake_layers"`
	Datasets       []string `json:"datasets"`
	LongTermFormat string   `json:"long_term_format"`
	Compression    string   `json:"compression"`
	PrivateRoot    string   `json:"private_root"`
}

type RecommendationsSummary struct {
	Count int                  `json:"count"`
	Items []RecommendationItem `json:"items"`
}

type RecommendationItem struct {
	Kind                       string `json:"kind"`
	Title                      string `json:"title"`
	Rationale                  string `json:"rationale"`
	Score                      int    `json:"score"`
	Currency                   string `json:"currency"`
	EstimatedMonthlyMinorUnits int64  `json:"estimated_monthly_minor_units"`
	EvidenceCount              int    `json:"evidence_count"`
}

type HouseholdSummary struct {
	Scopes []HouseholdScopeSummary `json:"scopes"`
}

type HouseholdScopeSummary struct {
	Scope                   string `json:"scope"`
	Label                   string `json:"label"`
	Currency                string `json:"currency"`
	FinanceRecords          int    `json:"finance_records"`
	FinanceCreditMinorUnits int64  `json:"finance_credit_minor_units"`
	FinanceDebitMinorUnits  int64  `json:"finance_debit_minor_units"`
	FinanceNetMinorUnits    int64  `json:"finance_net_minor_units"`
	PurchaseRecords         int    `json:"purchase_records"`
	PurchaseSpendMinorUnits int64  `json:"purchase_spend_minor_units"`
}
