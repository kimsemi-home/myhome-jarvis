package pdca

type Status struct {
	PolicyPath           string         `json:"policy_path"`
	LedgerPath           string         `json:"ledger_path"`
	Exists               bool           `json:"exists"`
	CycleCount           int            `json:"cycle_count"`
	OpenCount            int            `json:"open_count"`
	ClosedCount          int            `json:"closed_count"`
	InvalidCycleCount    int            `json:"invalid_cycle_count"`
	StepCount            int            `json:"step_count"`
	ReadyStepCount       int            `json:"ready_step_count"`
	MissingArtifactCount int            `json:"missing_artifact_count"`
	EvidenceSourceCount  int            `json:"evidence_source_count"`
	Ready                bool           `json:"ready"`
	ByStatus             map[string]int `json:"by_status"`
	LastObservedAt       string         `json:"last_observed_at,omitempty"`
	CheckedAt            string         `json:"checked_at"`
}
