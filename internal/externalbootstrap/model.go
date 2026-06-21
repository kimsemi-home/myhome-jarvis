package externalbootstrap

type Packet struct {
	Context                  string           `json:"context"`
	Version                  string           `json:"version"`
	PublicSafe               bool             `json:"public_safe"`
	CandidateRepo            string           `json:"candidate_repo"`
	CreationDecision         string           `json:"creation_decision"`
	CreationAllowed          bool             `json:"creation_allowed"`
	CreationBlockedReason    string           `json:"creation_blocked_reason,omitempty"`
	RequiredApprovalScope    string           `json:"required_approval_scope"`
	ApprovalLedgerState      string           `json:"approval_ledger_state"`
	ApprovalLeaseState       string           `json:"approval_lease_state,omitempty"`
	ApprovalLeaseExpiresAt   string           `json:"approval_lease_expires_at,omitempty"`
	ApprovalUnlocksScopeOnly bool             `json:"approval_unlocks_scope_only"`
	RepoSplitDecisionState   string           `json:"repo_split_decision_state"`
	RepoSplitCheckedAt       string           `json:"repo_split_checked_at"`
	RepoFactoryDecision      string           `json:"repo_factory_decision"`
	ContextHandoff           ContextHandoff   `json:"context_handoff"`
	SkeletonFiles            []SkeletonFile   `json:"skeleton_files"`
	HashCacheInputs          []HashCacheInput `json:"hash_cache_inputs"`
	PrivateLakeStaysPrivate  bool             `json:"private_lake_stays_private"`
	RawPayloadPublicAllowed  bool             `json:"raw_payload_public_allowed"`
	ExternalWritesAllowed    bool             `json:"external_writes_allowed"`
	NextSafeAction           string           `json:"next_safe_action"`
	CheckedAt                string           `json:"checked_at"`
}

type ContextHandoff struct {
	PackID                       string `json:"pack_id"`
	ContextPackVersion           string `json:"context_pack_version"`
	UpstreamCompatibilityVersion string `json:"upstream_compatibility_version"`
	OntologyVersion              string `json:"ontology_version"`
	AuthorityContractVersion     string `json:"authority_contract_version"`
	SecurityContractVersion      string `json:"security_contract_version"`
	VerificationProfile          string `json:"verification_profile"`
	ExportedArtifactCount        int    `json:"exported_artifact_count"`
}
