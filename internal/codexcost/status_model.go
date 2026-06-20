package codexcost

import "time"

type Status struct {
	PolicyPath           string           `json:"policy_path"`
	LedgerPath           string           `json:"ledger_path"`
	Exists               bool             `json:"exists"`
	RecordCount          int              `json:"record_count"`
	InvalidRecordCount   int              `json:"invalid_record_count"`
	ReviewRequiredCount  int              `json:"review_required_count"`
	MissingEvidenceCount int              `json:"missing_evidence_count"`
	TotalUnits           int64            `json:"total_units"`
	WarningUnitThreshold int64            `json:"warning_unit_threshold"`
	ReviewUnitThreshold  int64            `json:"review_unit_threshold"`
	BudgetState          string           `json:"budget_state"`
	ByUnitKind           map[string]int64 `json:"by_unit_kind"`
	ByScope              map[string]int64 `json:"by_scope"`
	ByStatus             map[string]int   `json:"by_status"`
	LastObservedAt       string           `json:"last_observed_at,omitempty"`
	CheckedAt            string           `json:"checked_at"`
}

func newStatus(policy Policy, checkedAt time.Time) Status {
	return Status{
		PolicyPath:           PolicyRelativePath,
		LedgerPath:           policy.PrivateUsageLedger,
		WarningUnitThreshold: policy.WarningUnitThreshold,
		ReviewUnitThreshold:  policy.ReviewUnitThreshold,
		BudgetState:          "ok",
		ByUnitKind:           map[string]int64{},
		ByScope:              map[string]int64{},
		ByStatus:             map[string]int{},
		CheckedAt:            checkedAt.Format(time.RFC3339),
	}
}
