package commandcenter

type Status struct {
	Context          string           `json:"context"`
	Version          string           `json:"version"`
	PublicSafe       bool             `json:"public_safe"`
	Redaction        string           `json:"redaction"`
	Vision           VisionSummary    `json:"vision"`
	PDCA             PDCASummary      `json:"pdca"`
	Evidence         EvidenceSummary  `json:"evidence"`
	Incidents        IncidentSummary  `json:"incidents"`
	Authority        AuthoritySummary `json:"authority"`
	Review           ReviewSummary    `json:"review"`
	Cost             CostSummary      `json:"cost"`
	BlockedGateCount int              `json:"blocked_gate_count"`
	BlockedGates     []GateSummary    `json:"blocked_gates"`
	NextSafeAction   string           `json:"next_safe_action"`
	CompactState     string           `json:"compact_state"`
	CheckedAt        string           `json:"checked_at"`
}

type GateSummary struct {
	Key    string `json:"key"`
	Label  string `json:"label"`
	State  string `json:"state"`
	Reason string `json:"reason"`
	Count  int    `json:"count"`
}
