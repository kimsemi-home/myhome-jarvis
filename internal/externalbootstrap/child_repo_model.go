package externalbootstrap

type ChildRepoStatus struct {
	Context                  string             `json:"context"`
	Version                  string             `json:"version"`
	PublicSafe               bool               `json:"public_safe"`
	CandidateRepo            string             `json:"candidate_repo"`
	CheckoutState            string             `json:"checkout_state"`
	EvidenceState            string             `json:"evidence_state"`
	Valid                    bool               `json:"valid"`
	ContextPackValid         bool               `json:"context_pack_valid"`
	HashCacheValid           bool               `json:"hash_cache_valid"`
	PublicSafetyOK           bool               `json:"public_safety_ok"`
	PrivateDataAbsent        bool               `json:"private_data_absent"`
	MissingFileCount         int                `json:"missing_file_count"`
	DriftCount               int                `json:"drift_count"`
	InvalidHashCacheCount    int                `json:"invalid_hash_cache_count"`
	PublicSafetyFindingCount int                `json:"public_safety_finding_count"`
	RequiredHashCacheKeys    []string           `json:"required_hash_cache_keys"`
	ObservedHashCacheKeys    []string           `json:"observed_hash_cache_keys"`
	ContextHandoff           ContextHandoff     `json:"context_handoff"`
	Findings                 []ChildRepoFinding `json:"findings,omitempty"`
	RawDetailsPublicAllowed  bool               `json:"raw_details_public_allowed"`
	NextSafeAction           string             `json:"next_safe_action"`
	CheckedAt                string             `json:"checked_at"`
}

type ChildRepoFinding struct {
	Path    string `json:"path"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type childContextPack struct {
	PackID                       string `json:"pack_id"`
	UpstreamRepo                 string `json:"upstream_repo"`
	Context                      string `json:"context"`
	ContextPackVersion           string `json:"context_pack_version"`
	UpstreamCompatibilityVersion string `json:"upstream_compatibility_version"`
	OntologyVersion              string `json:"ontology_version"`
	AuthorityContractVersion     string `json:"authority_contract_version"`
	SecurityContractVersion      string `json:"security_contract_version"`
	VerificationProfile          string `json:"verification_profile"`
	ExportedArtifactCount        int    `json:"exported_artifact_count"`
	PrivateLakeStaysPrivate      bool   `json:"private_lake_stays_private"`
	RawPayloadPublicAllowed      bool   `json:"raw_payload_public_allowed"`
	ExternalWritesAllowed        bool   `json:"external_writes_allowed"`
	RepoCreationScope            string `json:"repo_creation_scope"`
	CandidateRepo                string `json:"candidate_repo"`
}

type childHashCacheDocument struct {
	Context                   string           `json:"context"`
	Version                   string           `json:"version"`
	CandidateRepo             string           `json:"candidate_repo"`
	SourcePolicy              string           `json:"source_policy"`
	GeneratedContractVerified bool             `json:"generated_contract_verified"`
	HashCacheInputs           []HashCacheInput `json:"hash_cache_inputs"`
}
