package authority

type ReviewPlanStatus struct {
	PolicyPath                        string              `json:"policy_path"`
	Outcome                           string              `json:"outcome"`
	ActiveRule                        string              `json:"active_rule"`
	PublicSafe                        bool                `json:"public_safe"`
	Redaction                         string              `json:"redaction"`
	ReviewRequestable                 bool                `json:"review_requestable"`
	ReviewCapacityState               string              `json:"review_capacity_state"`
	NextSafeAction                    string              `json:"next_safe_action"`
	BlockedDecisionCount              int                 `json:"blocked_decision_count"`
	HighRiskBlockedDecisionCount      int                 `json:"high_risk_blocked_decision_count"`
	ReviewRequiredDecisionCount       int                 `json:"review_required_decision_count"`
	ReviewRequiredProfileCount        int                 `json:"review_required_profile_count"`
	PublicSafetyReviewProfileCount    int                 `json:"public_safety_review_profile_count"`
	PublicRepoReviewProfileCount      int                 `json:"public_repo_review_profile_count"`
	WorkflowReviewProfileCount        int                 `json:"workflow_review_profile_count"`
	VerifierSeparationRequiredCount   int                 `json:"verifier_separation_required_count"`
	SelfApprovalBlockedProfileCount   int                 `json:"self_approval_blocked_profile_count"`
	ExternalWritesAllowedProfileCount int                 `json:"external_writes_allowed_profile_count"`
	RequiredReviewClasses             []string            `json:"required_review_classes"`
	Profiles                          []ReviewProfilePlan `json:"profiles"`
	CheckedAt                         string              `json:"checked_at"`
}

type ReviewProfilePlan struct {
	ProfileKey                   string `json:"profile_key"`
	AuthorityProfile             string `json:"authority_profile"`
	ReviewClass                  string `json:"review_class"`
	RequiresHumanReview          bool   `json:"requires_human_review"`
	PublicSafetyGateRequired     bool   `json:"public_safety_gate_required"`
	PublicRepoChangeGateRequired bool   `json:"public_repo_change_gate_required"`
	WorkflowChangeGateRequired   bool   `json:"workflow_change_gate_required"`
	VerifierSeparationRequired   bool   `json:"verifier_separation_required"`
	SelfApprovalAllowed          bool   `json:"self_approval_allowed"`
}
