package commandcenter

type AuthoritySummary struct {
	Outcome                       string `json:"outcome"`
	ActiveRule                    string `json:"active_rule"`
	BlockedDecisionCount          int    `json:"blocked_decision_count"`
	AuthorityDebtCount            int    `json:"authority_debt_count"`
	PublicRepoMode                bool   `json:"public_repo_mode"`
	PublicSafetyOK                bool   `json:"public_safety_ok"`
	SelfAuthorityAllowed          bool   `json:"self_authority_allowed"`
	SelfApprovalBlockedProfiles   int    `json:"self_approval_blocked_profile_count"`
	ReviewRequiredProfileCount    int    `json:"review_required_profile_count"`
	PublicSafetyGatedProfileCount int    `json:"public_safety_gated_profile_count"`
}

type ReviewSummary struct {
	CapacityState     string `json:"capacity_state"`
	ActiveRule        string `json:"active_rule"`
	OpenCount         int    `json:"open_count"`
	HighRiskOpenCount int    `json:"high_risk_open_count"`
	ReviewDebtCount   int    `json:"review_debt_count"`
}

type CostSummary struct {
	BudgetState          string `json:"budget_state"`
	TotalUnits           int64  `json:"total_units"`
	WarningUnitThreshold int64  `json:"warning_unit_threshold"`
	ReviewUnitThreshold  int64  `json:"review_unit_threshold"`
	ReviewRequiredCount  int    `json:"review_required_count"`
	MissingEvidenceCount int    `json:"missing_evidence_count"`
}

type MonetizationSummary struct {
	ExperimentCount           int `json:"experiment_count"`
	DecisionCount             int `json:"decision_count"`
	ReviewRequiredCount       int `json:"review_required_count"`
	MonetizationDebtCount     int `json:"monetization_debt_count"`
	MissingEvidenceCount      int `json:"missing_evidence_count"`
	MissingCostEstimateCount  int `json:"missing_cost_estimate_count"`
	ExpectedValueUnknownCount int `json:"expected_value_unknown_count"`
}
