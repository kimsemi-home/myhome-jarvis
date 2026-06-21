package authority

type ApprovalDecisionStatus struct {
	PolicyPath                string                 `json:"policy_path"`
	PublicSafe                bool                   `json:"public_safe"`
	LedgerState               string                 `json:"ledger_state"`
	ActiveApprovalCount       int                    `json:"active_approval_count"`
	ExpiredApprovalCount      int                    `json:"expired_approval_count"`
	InvalidRecordCount        int                    `json:"invalid_record_count"`
	LatestScope               string                 `json:"latest_scope,omitempty"`
	LatestTarget              string                 `json:"latest_target,omitempty"`
	LatestLeaseState          string                 `json:"latest_lease_state,omitempty"`
	CanCreateRepo             bool                   `json:"can_create_repo"`
	CanChangeWorkflow         bool                   `json:"can_change_workflow"`
	CanWriteExternal          bool                   `json:"can_write_external"`
	UnrelatedAuthorityGranted bool                   `json:"unrelated_authority_granted"`
	ScopeSummaries            []ApprovalScopeSummary `json:"scope_summaries"`
	NextSafeAction            string                 `json:"next_safe_action"`
	CheckedAt                 string                 `json:"checked_at"`
}

type ApprovalScopeSummary struct {
	Scope              string             `json:"scope"`
	Target             string             `json:"target"`
	LeaseState         string             `json:"lease_state"`
	ExpiresAt          string             `json:"expires_at"`
	GrantFlags         ApprovalGrantFlags `json:"grant_flags"`
	CanUnlockScopeOnly bool               `json:"can_unlock_scope_only"`
}
