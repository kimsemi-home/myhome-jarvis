package commandcenter

type AuthorityReviewDecisionContract struct {
	Context                      string                                 `json:"context"`
	Version                      string                                 `json:"version"`
	PublicSafe                   bool                                   `json:"public_safe"`
	ReviewerPosture              string                                 `json:"reviewer_posture"`
	ReviewOnly                   bool                                   `json:"review_only"`
	CanApplyDecision             bool                                   `json:"can_apply_decision"`
	ReadyCapabilitiesNonBlocking []string                               `json:"ready_capabilities_non_blocking"`
	ContractItems                []AuthorityReviewDecisionContractItem  `json:"contract_items"`
	RequiredEvidenceChecks       []AuthorityReviewContractEvidenceCheck `json:"required_evidence_checks"`
	ForbiddenGrantFlags          AuthorityReviewContractForbiddenGrants `json:"forbidden_grant_flags"`
}

type AuthorityReviewDecisionContractItem struct {
	CapabilityKey               string   `json:"capability_key"`
	DecisionKey                 string   `json:"decision_key"`
	Scope                       string   `json:"scope"`
	RequiredReviewClass         string   `json:"required_review_class"`
	RequiredEvidenceKeys        []string `json:"required_evidence_keys"`
	HumanDecisionRecordRequired bool     `json:"human_decision_record_required"`
	ThisPacketGrantsApproval    bool     `json:"this_packet_grants_approval"`
	AllowsExternalWrites        bool     `json:"allows_external_writes"`
	AllowsRepoCreation          bool     `json:"allows_repo_creation"`
	AllowsWorkflowChanges       bool     `json:"allows_workflow_changes"`
	AllowsSelfApproval          bool     `json:"allows_self_approval"`
}

type AuthorityReviewContractEvidenceCheck struct {
	Key        string `json:"key"`
	State      string `json:"state"`
	Required   bool   `json:"required"`
	PublicSafe bool   `json:"public_safe"`
}

type AuthorityReviewContractForbiddenGrants struct {
	ApprovalGranted        bool `json:"approval_granted"`
	ExternalWritesAllowed  bool `json:"external_writes_allowed"`
	RepoCreationAllowed    bool `json:"repo_creation_allowed"`
	WorkflowChangesAllowed bool `json:"workflow_changes_allowed"`
	SelfApprovalAllowed    bool `json:"self_approval_allowed"`
}
