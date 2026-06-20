package authority

type ReviewRequestPacket struct {
	PolicyPath                        string   `json:"policy_path"`
	RequestID                         string   `json:"request_id"`
	RequestState                      string   `json:"request_state"`
	PublicSafe                        bool     `json:"public_safe"`
	Redaction                         string   `json:"redaction"`
	SourceAction                      string   `json:"source_action"`
	NextHandling                      string   `json:"next_handling"`
	ApprovalGranted                   bool     `json:"approval_granted"`
	ExternalWritesAllowed             bool     `json:"external_writes_allowed"`
	SelfApprovalAllowed               bool     `json:"self_approval_allowed"`
	ReviewRequestable                 bool     `json:"review_requestable"`
	ReviewCapacityState               string   `json:"review_capacity_state"`
	HighRiskBlockedDecisionCount      int      `json:"high_risk_blocked_decision_count"`
	ReviewRequiredDecisionCount       int      `json:"review_required_decision_count"`
	ReviewRequiredProfileCount        int      `json:"review_required_profile_count"`
	PublicRepoReviewProfileCount      int      `json:"public_repo_review_profile_count"`
	WorkflowReviewProfileCount        int      `json:"workflow_review_profile_count"`
	SelfApprovalBlockedProfileCount   int      `json:"self_approval_blocked_profile_count"`
	ExternalWritesAllowedProfileCount int      `json:"external_writes_allowed_profile_count"`
	RequiredReviewClasses             []string `json:"required_review_classes"`
	IncludedEvidenceFields            []string `json:"included_evidence_fields"`
	CheckedAt                         string   `json:"checked_at"`
}
