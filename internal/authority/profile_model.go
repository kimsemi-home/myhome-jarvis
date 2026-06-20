package authority

type AssistantProfile struct {
	Key                          string   `json:"key"`
	AuthorityProfile             string   `json:"authority_profile"`
	DataSensitivity              string   `json:"data_sensitivity"`
	RequiresHumanReview          bool     `json:"requires_human_review"`
	PublicSafetyGateRequired     bool     `json:"public_safety_gate_required"`
	PublicRepoChangeGateRequired bool     `json:"public_repo_change_gate_required"`
	WorkflowChangeGateRequired   bool     `json:"workflow_change_gate_required"`
	ExternalWritesAllowed        bool     `json:"external_writes_allowed"`
	VerifierSeparationRequired   bool     `json:"verifier_separation_required"`
	SelfApprovalAllowed          bool     `json:"self_approval_allowed"`
	RequiredEvidence             []string `json:"required_evidence"`
	AllowedDecisions             []string `json:"allowed_decisions"`
}
