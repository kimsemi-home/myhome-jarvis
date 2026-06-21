package externalevidence

type RepoSplitDecisionPacket struct {
	Context                         string                           `json:"context"`
	SchemaVersion                   string                           `json:"schema_version"`
	PublicSafe                      bool                             `json:"public_safe"`
	DecisionState                   string                           `json:"decision_state"`
	RecommendedOption               string                           `json:"recommended_option"`
	FutureRepoCandidate             string                           `json:"future_repo_candidate"`
	RepoCreationGate                string                           `json:"repo_creation_gate"`
	AuthorityDecisionRecordRequired bool                             `json:"authority_decision_record_required"`
	CanCreateRepo                   bool                             `json:"can_create_repo"`
	PrivateLakeStaysPrivate         bool                             `json:"private_lake_stays_private"`
	RawPayloadPublicAllowed         bool                             `json:"raw_payload_public_allowed"`
	ExternalWritesAllowed           bool                             `json:"external_writes_allowed"`
	ContextPackVersion              string                           `json:"context_pack_version"`
	OntologyVersion                 string                           `json:"ontology_version"`
	Options                         []RepoSplitDecisionOption        `json:"options"`
	EvidenceChecks                  []RepoSplitDecisionEvidenceCheck `json:"evidence_checks"`
	ForbiddenGrantFlags             RepoSplitDecisionForbiddenGrants `json:"forbidden_grant_flags"`
	NextSafeAction                  string                           `json:"next_safe_action"`
	CheckedAt                       string                           `json:"checked_at"`
}

type RepoSplitDecisionOption struct {
	Key                      string `json:"key"`
	Label                    string `json:"label"`
	Summary                  string `json:"summary"`
	PrivacyRisk              string `json:"privacy_risk"`
	MaintenanceBurden        string `json:"maintenance_burden"`
	SSOTHandoff              string `json:"ssot_handoff"`
	ContextPackHandoff       string `json:"context_pack_handoff"`
	GitHubActionsCost        string `json:"github_actions_cost"`
	ArchiveCacheBehavior     string `json:"archive_cache_behavior"`
	OntologyVersionDiscovery string `json:"ontology_version_discovery"`
	RepoCreationAllowed      bool   `json:"repo_creation_allowed"`
	RawPayloadPublicAllowed  bool   `json:"raw_payload_public_allowed"`
	HumanApprovalRequired    bool   `json:"human_approval_required"`
}
