package commandcenter

type Status struct {
	Context              string                      `json:"context"`
	Version              string                      `json:"version"`
	PublicSafe           bool                        `json:"public_safe"`
	Redaction            string                      `json:"redaction"`
	Vision               VisionSummary               `json:"vision"`
	PDCA                 PDCASummary                 `json:"pdca"`
	Evidence             EvidenceSummary             `json:"evidence"`
	EvidenceIntegrity    EvidenceIntegritySummary    `json:"evidence_integrity"`
	Incidents            IncidentSummary             `json:"incidents"`
	Authority            AuthoritySummary            `json:"authority"`
	AuthorityReview      AuthorityReviewSummary      `json:"authority_review"`
	Review               ReviewSummary               `json:"review"`
	FinanceConsent       FinanceConsentSummary       `json:"finance_consent"`
	Cost                 CostSummary                 `json:"cost"`
	CodexCostBrief       CodexCostBriefSummary       `json:"codex_cost_brief"`
	CodexCostScaling     CodexCostScalingSummary     `json:"codex_cost_scaling"`
	CodexSustainability  CodexSustainabilitySummary  `json:"codex_sustainability"`
	StorageArchive       StorageArchiveSummary       `json:"storage_archive"`
	Supervisor           SupervisorSummary           `json:"supervisor"`
	ContextPack          ContextPackSummary          `json:"context_pack"`
	MediaReadiness       MediaReadinessSummary       `json:"media_readiness"`
	MergeEvidence        MergeEvidenceSummary        `json:"merge_evidence"`
	Monetization         MonetizationSummary         `json:"monetization"`
	RepoFactory          RepoFactorySummary          `json:"repo_factory"`
	RepoFactoryPreflight RepoFactoryPreflightSummary `json:"repo_factory_preflight"`
	WorkItem             WorkItemSummary             `json:"work_item"`
	BlockedGateCount     int                         `json:"blocked_gate_count"`
	BlockedGates         []GateSummary               `json:"blocked_gates"`
	NextSafeAction       string                      `json:"next_safe_action"`
	CompactState         string                      `json:"compact_state"`
	CheckedAt            string                      `json:"checked_at"`
}

type GateSummary struct {
	Key    string `json:"key"`
	Label  string `json:"label"`
	State  string `json:"state"`
	Reason string `json:"reason"`
	Count  int    `json:"count"`
}
