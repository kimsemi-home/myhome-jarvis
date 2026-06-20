package commandcenter

type RepoFactoryPreflightSummary struct {
	PublicSafe                     bool                       `json:"public_safe"`
	CreationDecision               string                     `json:"creation_decision"`
	CreationAllowed                bool                       `json:"creation_allowed"`
	RepoCreationBlockedUntilReview bool                       `json:"repo_creation_blocked_until_review"`
	SelfApprovalAllowed            bool                       `json:"self_approval_allowed"`
	TemplateReadyCount             int                        `json:"template_ready_count"`
	TemplateFileCount              int                        `json:"template_file_count"`
	GateReadyCount                 int                        `json:"gate_ready_count"`
	CreationGateCount              int                        `json:"creation_gate_count"`
	BlockingGateCount              int                        `json:"blocking_gate_count"`
	MissingEvidenceKeys            []string                   `json:"missing_evidence_keys"`
	NextSafeAction                 string                     `json:"next_safe_action"`
	ContextPackEvidence            ContextPackEvidenceSummary `json:"context_pack_evidence"`
}

type ContextPackEvidenceSummary struct {
	DeclarationPath              string `json:"declaration_path"`
	Valid                        bool   `json:"valid"`
	EvidenceState                string `json:"evidence_state"`
	DriftCount                   int    `json:"drift_count"`
	MissingFieldCount            int    `json:"missing_field_count"`
	MissingArtifactCount         int    `json:"missing_artifact_count"`
	StaleVersionCount            int    `json:"stale_version_count"`
	ForbiddenValueCount          int    `json:"forbidden_value_count"`
	PackID                       string `json:"pack_id"`
	ContextPackVersion           string `json:"context_pack_version"`
	UpstreamCompatibilityVersion string `json:"upstream_compatibility_version"`
	OntologyVersion              string `json:"ontology_version"`
	AuthorityContractVersion     string `json:"authority_contract_version"`
	SecurityContractVersion      string `json:"security_contract_version"`
	VerificationProfile          string `json:"verification_profile"`
	ExportedArtifactCount        int    `json:"exported_artifact_count"`
	RawDetailsPublicAllowed      bool   `json:"raw_details_public_allowed"`
}
