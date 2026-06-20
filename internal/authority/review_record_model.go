package authority

type ReviewRecord struct {
	At                         string   `json:"at"`
	RequestID                  string   `json:"request_id"`
	RequestState               string   `json:"request_state"`
	EvidenceRef                string   `json:"evidence_ref"`
	EvidenceState              string   `json:"evidence_state"`
	QueueItemRef               string   `json:"queue_item_ref"`
	QueueState                 string   `json:"queue_state"`
	LinearIssueRef             string   `json:"linear_issue_ref,omitempty"`
	SourceAction               string   `json:"source_action"`
	NextSafeAction             string   `json:"next_safe_action"`
	ReviewCapacityState        string   `json:"review_capacity_state"`
	RequiredReviewClasses      []string `json:"required_review_classes"`
	RequiredReviewClassCount   int      `json:"required_review_class_count"`
	HighRiskBlockedDecisionCnt int      `json:"high_risk_blocked_decision_count"`
	ReviewRequiredDecisionCnt  int      `json:"review_required_decision_count"`
	ReviewRequiredProfileCnt   int      `json:"review_required_profile_count"`
	ApprovalState              string   `json:"approval_state"`
	ApprovalGranted            bool     `json:"approval_granted"`
	ExternalWritesAllowed      bool     `json:"external_writes_allowed"`
	SelfApprovalAllowed        bool     `json:"self_approval_allowed"`
	PublicSafe                 bool     `json:"public_safe"`
}

type ReviewRecordRequest struct {
	At                    string   `json:"at,omitempty"`
	RequestID             string   `json:"request_id"`
	EvidenceRef           string   `json:"evidence_ref"`
	QueueItemRef          string   `json:"queue_item_ref"`
	QueueState            string   `json:"queue_state"`
	RequiredReviewClasses []string `json:"required_review_classes"`
	LinearIssueRef        string   `json:"linear_issue_ref,omitempty"`
	ApprovalGranted       *bool    `json:"approval_granted"`
	ExternalWritesAllowed *bool    `json:"external_writes_allowed"`
	SelfApprovalAllowed   *bool    `json:"self_approval_allowed"`
}

type ReviewRecordResult struct {
	RequestID                string `json:"request_id"`
	LinearIssueRef           string `json:"linear_issue_ref,omitempty"`
	RequestState             string `json:"request_state"`
	QueueState               string `json:"queue_state"`
	LedgerState              string `json:"ledger_state"`
	ApprovalState            string `json:"approval_state"`
	ApprovalGranted          bool   `json:"approval_granted"`
	ExternalWritesAllowed    bool   `json:"external_writes_allowed"`
	SelfApprovalAllowed      bool   `json:"self_approval_allowed"`
	RequiredReviewClassCount int    `json:"required_review_class_count"`
	RecordedAt               string `json:"recorded_at"`
	PublicSafe               bool   `json:"public_safe"`
}
