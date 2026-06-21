package authority

type ApprovalGrantFlags struct {
	ApprovalGranted        bool `json:"approval_granted"`
	RepoCreationAllowed    bool `json:"repo_creation_allowed"`
	WorkflowChangesAllowed bool `json:"workflow_changes_allowed"`
	ExternalWritesAllowed  bool `json:"external_writes_allowed"`
	SelfApprovalAllowed    bool `json:"self_approval_allowed"`
}

type ApprovalDecisionResult struct {
	DecisionPacketRef  string             `json:"decision_packet_ref"`
	Scope              string             `json:"scope"`
	Target             string             `json:"target"`
	LeaseState         string             `json:"lease_state"`
	ExpiresAt          string             `json:"expires_at"`
	GrantFlags         ApprovalGrantFlags `json:"grant_flags"`
	LedgerState        string             `json:"ledger_state"`
	RecordedAt         string             `json:"recorded_at"`
	PublicSafe         bool               `json:"public_safe"`
	CanUnlockScopeOnly bool               `json:"can_unlock_scope_only"`
}
