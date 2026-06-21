package externalevidence

type RepoSplitDecisionEvidenceCheck struct {
	Key        string `json:"key"`
	State      string `json:"state"`
	Required   bool   `json:"required"`
	PublicSafe bool   `json:"public_safe"`
}

type RepoSplitDecisionForbiddenGrants struct {
	ApprovalGranted        bool `json:"approval_granted"`
	ExternalWritesAllowed  bool `json:"external_writes_allowed"`
	RepoCreationAllowed    bool `json:"repo_creation_allowed"`
	WorkflowChangesAllowed bool `json:"workflow_changes_allowed"`
	SelfApprovalAllowed    bool `json:"self_approval_allowed"`
}
