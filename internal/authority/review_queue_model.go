package authority

type ReviewQueueStatus struct {
	PolicyPath              string `json:"policy_path"`
	RequestID               string `json:"request_id"`
	EvidenceRef             string `json:"evidence_ref"`
	QueueItemRef            string `json:"queue_item_ref"`
	RequestState            string `json:"request_state"`
	EvidenceState           string `json:"evidence_state"`
	QueueState              string `json:"queue_state"`
	QueueReady              bool   `json:"queue_ready"`
	PublicSafe              bool   `json:"public_safe"`
	Redaction               string `json:"redaction"`
	PendingReviewClassCount int    `json:"pending_review_class_count"`
	ApprovalState           string `json:"approval_state"`
	ApprovalGranted         bool   `json:"approval_granted"`
	ExternalWritesAllowed   bool   `json:"external_writes_allowed"`
	SelfApprovalAllowed     bool   `json:"self_approval_allowed"`
	NextSafeAction          string `json:"next_safe_action"`
	CheckedAt               string `json:"checked_at"`
}
