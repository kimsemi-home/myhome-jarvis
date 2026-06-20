package commandcenter

type WorkItemSummary struct {
	WorkItemRef            string   `json:"work_item_ref"`
	WorkItemState          string   `json:"work_item_state"`
	IntentKey              string   `json:"intent_key"`
	CapabilityKeys         []string `json:"capability_keys"`
	DecisionKey            string   `json:"decision_key"`
	EvidenceRef            string   `json:"evidence_ref"`
	AuthorityRef           string   `json:"authority_ref"`
	GuardrailKeys          []string `json:"guardrail_keys"`
	SourceAction           string   `json:"source_action"`
	BlockedGateKeys        []string `json:"blocked_gate_keys"`
	QueueState             string   `json:"queue_state"`
	ReviewClassCount       int      `json:"review_class_count"`
	ReviewRequestAgeHours  int      `json:"review_request_age_hours"`
	ReviewStaleAfterHours  int      `json:"review_request_stale_after_hours"`
	ReviewRequestStale     bool     `json:"review_request_stale"`
	ReviewEscalationAction string   `json:"review_request_escalation_action"`
	MergeEligibilityHint   string   `json:"merge_eligibility_hint"`
	PublicSafe             bool     `json:"public_safe"`
	Redaction              string   `json:"redaction"`
	ReviewOnly             bool     `json:"review_only"`
	ApprovalState          string   `json:"approval_state"`
	ApprovalGranted        bool     `json:"approval_granted"`
	ExternalWritesAllowed  bool     `json:"external_writes_allowed"`
	SelfApprovalAllowed    bool     `json:"self_approval_allowed"`
	NextSafeAction         string   `json:"next_safe_action"`
}

type WorkItemStatus struct {
	Context string `json:"context"`
	Version string `json:"version"`
	WorkItemSummary
	CheckedAt string `json:"checked_at"`
}
