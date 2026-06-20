package commandcenter

type AuthorityReviewDecisionPacket struct {
	Context                      string                          `json:"context"`
	Version                      string                          `json:"version"`
	PublicSafe                   bool                            `json:"public_safe"`
	Redaction                    string                          `json:"redaction"`
	PolicyPath                   string                          `json:"policy_path"`
	RequestID                    string                          `json:"request_id"`
	RequestState                 string                          `json:"request_state"`
	QueueState                   string                          `json:"queue_state"`
	EvidenceReady                bool                            `json:"evidence_ready"`
	ReviewRequestRecorded        bool                            `json:"review_request_recorded"`
	ReviewRequestLedgerState     string                          `json:"review_request_ledger_state"`
	ReviewRequestAgeHours        int                             `json:"review_request_age_hours"`
	ReviewRequestStaleAfterHours int                             `json:"review_request_stale_after_hours"`
	ReviewRequestStale           bool                            `json:"review_request_stale"`
	ReviewEscalationAction       string                          `json:"review_request_escalation_action"`
	RequiredReviewClasses        []string                        `json:"required_review_classes"`
	RequiredReviewClassCount     int                             `json:"required_review_class_count"`
	GatedCapabilityKeys          []string                        `json:"gated_capability_keys"`
	BlockedGateKeys              []string                        `json:"blocked_gate_keys"`
	ApprovalBoundary             AuthorityReviewBoundary         `json:"approval_boundary"`
	RepoFactoryGate              AuthorityReviewRepoFactory      `json:"repo_factory_gate"`
	RepoFactoryPreflight         RepoFactoryPreflightSummary     `json:"repo_factory_preflight"`
	PublicSafetyPosture          AuthorityReviewPublicSafety     `json:"public_safety_posture"`
	StorageEvidence              StorageArchiveSummary           `json:"storage_evidence"`
	LocalRuntime                 LocalRuntimeSummary             `json:"local_runtime"`
	MergeEvidence                MergeEvidenceSummary            `json:"merge_evidence"`
	CodexSustainability          CodexSustainabilitySummary      `json:"codex_sustainability"`
	DecisionPacketState          string                          `json:"decision_packet_state"`
	CanApplyDecision             bool                            `json:"can_apply_decision"`
	DecisionOptions              []AuthorityReviewDecisionOption `json:"decision_options"`
	NextSafeAction               string                          `json:"next_safe_action"`
	CheckedAt                    string                          `json:"checked_at"`
}

type AuthorityReviewPublicSafety struct {
	PublicRepoMode                  bool `json:"public_repo_mode"`
	PublicSafetyOK                  bool `json:"public_safety_ok"`
	HighRiskBlockedDecisionCount    int  `json:"high_risk_blocked_decision_count"`
	PublicRepoReviewProfileCount    int  `json:"public_repo_review_profile_count"`
	WorkflowReviewProfileCount      int  `json:"workflow_review_profile_count"`
	SelfApprovalBlockedProfileCount int  `json:"self_approval_blocked_profile_count"`
}

type AuthorityReviewDecisionOption struct {
	Key                           string `json:"key"`
	Label                         string `json:"label"`
	Effect                        string `json:"effect"`
	RequiresHuman                 bool   `json:"requires_human"`
	RequiresSeparateRecordCommand bool   `json:"requires_separate_record_command"`
	ThisPacketGrantsApproval      bool   `json:"this_packet_grants_approval"`
	AllowsExternalWrites          bool   `json:"allows_external_writes"`
	AllowsRepoCreation            bool   `json:"allows_repo_creation"`
	AllowsWorkflowChanges         bool   `json:"allows_workflow_changes"`
	AllowsSelfApproval            bool   `json:"allows_self_approval"`
}
