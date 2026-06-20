package commandcenter

type VisionSummary struct {
	PolicyPath         string   `json:"policy_path"`
	Mission            string   `json:"mission"`
	OperatingMode      string   `json:"operating_mode"`
	CapabilityCount    int      `json:"capability_count"`
	GuardrailCount     int      `json:"guardrail_count"`
	PillarKeys         []string `json:"pillar_keys"`
	ReadyPillarCount   int      `json:"ready_pillar_count"`
	GatedPillarCount   int      `json:"gated_pillar_count"`
	BlockedPillarCount int      `json:"blocked_pillar_count"`
	ReadyPillarKeys    []string `json:"ready_pillar_keys"`
	GatedPillarKeys    []string `json:"gated_pillar_keys"`
	BlockedPillarKeys  []string `json:"blocked_pillar_keys"`
}

type PDCASummary struct {
	Ready                bool `json:"ready"`
	CycleCount           int  `json:"cycle_count"`
	ReadyStepCount       int  `json:"ready_step_count"`
	MissingArtifactCount int  `json:"missing_artifact_count"`
}

type EvidenceSummary struct {
	SourceCount              int `json:"source_count"`
	PresentSourceCount       int `json:"present_source_count"`
	NodeCount                int `json:"node_count"`
	EdgeCount                int `json:"edge_count"`
	DanglingEvidenceRefCount int `json:"dangling_evidence_ref_count"`
	OpenLearningCount        int `json:"open_learning_count"`
}

type IncidentSummary struct {
	OpenCount               int `json:"open_count"`
	IncidentDebtCount       int `json:"incident_debt_count"`
	StaleQuarantineCount    int `json:"stale_quarantine_count"`
	MissingEvidenceRefCount int `json:"missing_evidence_ref_count"`
}

type SupervisorSummary struct {
	Recorded       bool   `json:"recorded"`
	StatePath      string `json:"state_path"`
	ProcessRunning bool   `json:"process_running"`
	ProbeOK        bool   `json:"probe_ok"`
	Stale          bool   `json:"stale"`
	Message        string `json:"message"`
}
