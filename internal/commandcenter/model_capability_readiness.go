package commandcenter

type CapabilityReadinessSummary struct {
	PublicSafe             bool                            `json:"public_safe"`
	CapabilityCount        int                             `json:"capability_count"`
	ReadyCapabilityCount   int                             `json:"ready_capability_count"`
	GatedCapabilityCount   int                             `json:"gated_capability_count"`
	BlockedCapabilityCount int                             `json:"blocked_capability_count"`
	ReadyCapabilityKeys    []string                        `json:"ready_capability_keys"`
	GatedCapabilityKeys    []string                        `json:"gated_capability_keys"`
	BlockedCapabilityKeys  []string                        `json:"blocked_capability_keys"`
	Media                  CapabilityMediaReadiness        `json:"media"`
	FinanceConsent         CapabilityFinanceReadiness      `json:"finance_consent"`
	Monetization           CapabilityMonetizationReadiness `json:"monetization"`
	CodexCost              CapabilityCodexCostReadiness    `json:"codex_cost"`
}

type CapabilityMediaReadiness struct {
	State                   string `json:"state"`
	PublicSafe              bool   `json:"public_safe"`
	PlaybackReady           bool   `json:"playback_ready"`
	PlaybackAvailableCount  int    `json:"playback_available_count"`
	DegradedCount           int    `json:"degraded_count"`
	MaxPlanningLatencyMS    int64  `json:"max_planning_latency_ms"`
	TargetPlanningLatencyMS int64  `json:"target_planning_latency_ms"`
	LocalLauncherAvailable  bool   `json:"local_launcher_available"`
}

type CapabilityFinanceReadiness struct {
	State                       string `json:"state"`
	ReadinessState              string `json:"readiness_state"`
	FinanceMode                 string `json:"finance_mode"`
	ActiveConsentCount          int    `json:"active_consent_count"`
	MissingRequiredConsentCount int    `json:"missing_required_consent_count"`
	ReviewRequiredCount         int    `json:"review_required_count"`
	MissingEvidenceCount        int    `json:"missing_evidence_count"`
	ForbiddenActionEnabledCount int    `json:"forbidden_action_enabled_count"`
	ConsentDebtCount            int    `json:"consent_debt_count"`
}

type CapabilityMonetizationReadiness struct {
	State                     string `json:"state"`
	ExperimentCount           int    `json:"experiment_count"`
	DecisionCount             int    `json:"decision_count"`
	ReviewRequiredCount       int    `json:"review_required_count"`
	MissingEvidenceCount      int    `json:"missing_evidence_count"`
	MissingCostEstimateCount  int    `json:"missing_cost_estimate_count"`
	ExpectedValueUnknownCount int    `json:"expected_value_unknown_count"`
	MonetizationDebtCount     int    `json:"monetization_debt_count"`
}

type CapabilityCodexCostReadiness struct {
	State                      string `json:"state"`
	PublicSafe                 bool   `json:"public_safe"`
	ScalingPublicSafe          bool   `json:"scaling_public_safe"`
	BudgetState                string `json:"budget_state"`
	TotalUnits                 int64  `json:"total_units"`
	ReviewRequiredCount        int    `json:"review_required_count"`
	MissingEvidenceCount       int    `json:"missing_evidence_count"`
	BriefDecision              string `json:"brief_decision"`
	BriefNextSafeAction        string `json:"brief_next_safe_action"`
	CanApplyExpansion          bool   `json:"can_apply_expansion"`
	ReviewGateCount            int    `json:"review_gate_count"`
	GrantingScalingOptionCount int    `json:"granting_scaling_option_count"`
}
