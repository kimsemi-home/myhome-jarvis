package commandcenter

type LocalRuntimeSummary struct {
	PublicSafe              bool   `json:"public_safe"`
	EvidenceRef             string `json:"evidence_ref"`
	State                   string `json:"state"`
	Recorded                bool   `json:"recorded"`
	ProcessRunning          bool   `json:"process_running"`
	ProbeOK                 bool   `json:"probe_ok"`
	Stale                   bool   `json:"stale"`
	HealthDebtCount         int    `json:"health_debt_count"`
	Message                 string `json:"message"`
	NextSafeAction          string `json:"next_safe_action"`
	RawRuntimePublicAllowed bool   `json:"raw_runtime_public_allowed"`
}
