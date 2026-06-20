package commandcenter

type RepoFactoryPreflightSummary struct {
	PublicSafe                     bool     `json:"public_safe"`
	CreationDecision               string   `json:"creation_decision"`
	CreationAllowed                bool     `json:"creation_allowed"`
	RepoCreationBlockedUntilReview bool     `json:"repo_creation_blocked_until_review"`
	SelfApprovalAllowed            bool     `json:"self_approval_allowed"`
	TemplateReadyCount             int      `json:"template_ready_count"`
	TemplateFileCount              int      `json:"template_file_count"`
	GateReadyCount                 int      `json:"gate_ready_count"`
	CreationGateCount              int      `json:"creation_gate_count"`
	BlockingGateCount              int      `json:"blocking_gate_count"`
	MissingEvidenceKeys            []string `json:"missing_evidence_keys"`
	NextSafeAction                 string   `json:"next_safe_action"`
}
