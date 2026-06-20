package codexcost

type AttributionStatus struct {
	LedgerPath           string           `json:"ledger_path"`
	Exists               bool             `json:"exists"`
	RecordCount          int              `json:"record_count"`
	InvalidRecordCount   int              `json:"invalid_record_count"`
	MissingEvidenceCount int              `json:"missing_evidence_count"`
	TotalUnits           int64            `json:"total_units"`
	EntryUnits           int64            `json:"entry_units"`
	CoverageUnits        int64            `json:"coverage_units"`
	ByScope              map[string]int64 `json:"by_scope"`
	SubjectCountByScope  map[string]int   `json:"subject_count_by_scope"`
	CostRefCount         int              `json:"cost_ref_count"`
	CheckedAt            string           `json:"checked_at"`
	subjectsByScope      map[string]map[string]bool
	costRefUnits         map[string]int64
}
