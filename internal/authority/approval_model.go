package authority

type ApprovalDecisionRecord struct {
	At                      string             `json:"at"`
	DecisionPacketRef       string             `json:"decision_packet_ref"`
	DecisionPacketContext   string             `json:"decision_packet_context"`
	DecisionPacketCheckedAt string             `json:"decision_packet_checked_at"`
	Scope                   string             `json:"scope"`
	Target                  string             `json:"target"`
	ReviewerBoundary        string             `json:"reviewer_boundary"`
	ReviewerIsRequester     bool               `json:"reviewer_is_requester"`
	ExpiresAt               string             `json:"expires_at"`
	LeaseState              string             `json:"lease_state"`
	GrantFlags              ApprovalGrantFlags `json:"grant_flags"`
	PublicSafe              bool               `json:"public_safe"`
}

type ApprovalDecisionRequest struct {
	At                      string `json:"at,omitempty"`
	DecisionPacketRef       string `json:"decision_packet_ref"`
	DecisionPacketContext   string `json:"decision_packet_context"`
	DecisionPacketCheckedAt string `json:"decision_packet_checked_at"`
	Scope                   string `json:"scope"`
	Target                  string `json:"target"`
	ReviewerBoundary        string `json:"reviewer_boundary"`
	ReviewerIsRequester     *bool  `json:"reviewer_is_requester"`
	ExpiresAt               string `json:"expires_at"`
	ApprovalGranted         *bool  `json:"approval_granted"`
	RepoCreationAllowed     *bool  `json:"repo_creation_allowed"`
	WorkflowChangesAllowed  *bool  `json:"workflow_changes_allowed"`
	ExternalWritesAllowed   *bool  `json:"external_writes_allowed"`
	SelfApprovalAllowed     *bool  `json:"self_approval_allowed"`
}
