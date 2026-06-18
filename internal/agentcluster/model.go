package agentcluster

type Status struct {
	Context                       string    `json:"context"`
	Version                       string    `json:"version"`
	PublicSafe                    bool      `json:"public_safe"`
	RawTranscriptStorageAllowed   bool      `json:"raw_transcript_storage_allowed"`
	PrivateDataInEvidenceAllowed  bool      `json:"private_data_in_evidence_allowed"`
	ExternalAgentExecutionAllowed bool      `json:"external_agent_execution_allowed"`
	SelfApprovalAllowed           bool      `json:"self_approval_allowed"`
	ConfidenceSelfReportAllowed   bool      `json:"confidence_self_report_allowed"`
	AuthorityGateRequired         bool      `json:"authority_gate_required"`
	EvidenceStageCount            int       `json:"evidence_stage_count"`
	RoleCount                     int       `json:"role_count"`
	SidecarCount                  int       `json:"sidecar_count"`
	DebtTypeCount                 int       `json:"debt_type_count"`
	FailureConditionCount         int       `json:"failure_condition_count"`
	GeneratedPath                 string    `json:"generated_path"`
	EvidenceFlow                  []string  `json:"evidence_flow"`
	IncidentLifecycle             []string  `json:"incident_lifecycle"`
	Roles                         []Role    `json:"roles"`
	Sidecars                      []Sidecar `json:"sidecars"`
	Signals                       []Signal  `json:"signals"`
	Commands                      []string  `json:"commands"`
	Message                       string    `json:"message"`
	CheckedAt                     string    `json:"checked_at"`
}

type Role struct {
	Key           string   `json:"key"`
	Label         string   `json:"label"`
	ReasoningTier string   `json:"reasoning_tier"`
	Authority     string   `json:"authority"`
	MustProduce   []string `json:"must_produce"`
	MustNot       []string `json:"must_not"`
}

type Sidecar struct {
	Key    string   `json:"key"`
	Label  string   `json:"label"`
	Checks []string `json:"checks"`
}

type Signal struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Status   string `json:"status"`
	Evidence string `json:"evidence"`
}
