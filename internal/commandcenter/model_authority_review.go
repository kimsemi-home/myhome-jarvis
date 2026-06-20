package commandcenter

type AuthorityReviewSummary struct {
	PublicSafe                      bool   `json:"public_safe"`
	ReviewRequestable               bool   `json:"review_requestable"`
	ReviewCapacityState             string `json:"review_capacity_state"`
	NextSafeAction                  string `json:"next_safe_action"`
	HighRiskBlockedDecisionCount    int    `json:"high_risk_blocked_decision_count"`
	RequiredReviewClassCount        int    `json:"required_review_class_count"`
	PublicRepoReviewProfileCount    int    `json:"public_repo_review_profile_count"`
	WorkflowReviewProfileCount      int    `json:"workflow_review_profile_count"`
	SelfApprovalBlockedProfileCount int    `json:"self_approval_blocked_profile_count"`
}
