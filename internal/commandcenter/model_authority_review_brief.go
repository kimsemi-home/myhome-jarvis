package commandcenter

type AuthorityReviewBrief struct {
	Context                      string                      `json:"context"`
	Version                      string                      `json:"version"`
	PublicSafe                   bool                        `json:"public_safe"`
	Redaction                    string                      `json:"redaction"`
	PolicyPath                   string                      `json:"policy_path"`
	RequestID                    string                      `json:"request_id"`
	RequestState                 string                      `json:"request_state"`
	EvidenceRef                  string                      `json:"evidence_ref"`
	EvidenceReady                bool                        `json:"evidence_ready"`
	QueueState                   string                      `json:"queue_state"`
	QueueReady                   bool                        `json:"queue_ready"`
	ReviewRequestRecorded        bool                        `json:"review_request_recorded"`
	ReviewRequestLedgerState     string                      `json:"review_request_ledger_state"`
	ReviewRequestAgeHours        int                         `json:"review_request_age_hours"`
	ReviewRequestStaleAfterHours int                         `json:"review_request_stale_after_hours"`
	ReviewRequestStale           bool                        `json:"review_request_stale"`
	ReviewEscalationAction       string                      `json:"review_request_escalation_action"`
	RequiredReviewClasses        []string                    `json:"required_review_classes"`
	RequiredReviewClassCount     int                         `json:"required_review_class_count"`
	GatedCapabilityKeys          []string                    `json:"gated_capability_keys"`
	BlockedGateKeys              []string                    `json:"blocked_gate_keys"`
	WorkItemRef                  string                      `json:"work_item_ref"`
	WorkItemState                string                      `json:"work_item_state"`
	DecisionKey                  string                      `json:"decision_key"`
	AuthorityRef                 string                      `json:"authority_ref"`
	ApprovalBoundary             AuthorityReviewBoundary     `json:"approval_boundary"`
	RepoFactoryGate              AuthorityReviewRepoFactory  `json:"repo_factory_gate"`
	RepoFactoryPreflight         RepoFactoryPreflightSummary `json:"repo_factory_preflight"`
	LocalRuntime                 LocalRuntimeSummary         `json:"local_runtime"`
	MergeEvidence                MergeEvidenceSummary        `json:"merge_evidence"`
	CodexSustainability          CodexSustainabilitySummary  `json:"codex_sustainability"`
	VisionGoalComplete           bool                        `json:"vision_goal_complete"`
	VisionNextSafeAction         string                      `json:"vision_next_safe_action"`
	NextSafeAction               string                      `json:"next_safe_action"`
	CheckedAt                    string                      `json:"checked_at"`
}

type AuthorityReviewBoundary struct {
	ApprovalState         string `json:"approval_state"`
	ApprovalGranted       bool   `json:"approval_granted"`
	ExternalWritesAllowed bool   `json:"external_writes_allowed"`
	SelfApprovalAllowed   bool   `json:"self_approval_allowed"`
	ReviewOnly            bool   `json:"review_only"`
}

type AuthorityReviewRepoFactory struct {
	PublicSafe                     bool `json:"public_safe"`
	AuthorityReviewRequired        bool `json:"authority_review_required"`
	PublicSafetyEvidenceRequired   bool `json:"public_safety_evidence_required"`
	RepoCreationBlockedUntilReview bool `json:"repo_creation_blocked_until_review"`
	MissingCreationGateCount       int  `json:"missing_creation_gate_count"`
	ForbiddenTemplateValueCount    int  `json:"forbidden_template_value_count"`
}
