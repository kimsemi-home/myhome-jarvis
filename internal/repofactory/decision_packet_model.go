package repofactory

type DecisionPacket struct {
	Context                        string               `json:"context"`
	Version                        string               `json:"version"`
	PolicyPath                     string               `json:"policy_path"`
	PublicSafe                     bool                 `json:"public_safe"`
	CreationDecision               string               `json:"creation_decision"`
	CreationAllowed                bool                 `json:"creation_allowed"`
	RepoCreationBlockedUntilReview bool                 `json:"repo_creation_blocked_until_review"`
	SelfApprovalAllowed            bool                 `json:"self_approval_allowed"`
	HumanReviewRequired            bool                 `json:"human_review_required"`
	PublicSafetyEvidenceRequired   bool                 `json:"public_safety_evidence_required"`
	CodexProjectRequired           bool                 `json:"codex_project_required"`
	TemplateReadyCount             int                  `json:"template_ready_count"`
	TemplateFileCount              int                  `json:"template_file_count"`
	GateReadyCount                 int                  `json:"gate_ready_count"`
	CreationGateCount              int                  `json:"creation_gate_count"`
	BlockingGateCount              int                  `json:"blocking_gate_count"`
	MissingEvidenceKeys            []string             `json:"missing_evidence_keys"`
	NextSafeAction                 string               `json:"next_safe_action"`
	PublicSafetyEvidence           PublicSafetyEvidence `json:"public_safety_evidence"`
	ContextPackEvidence            ContextPackEvidence  `json:"context_pack_evidence"`
	TemplateEvidence               []TemplateEvidence   `json:"template_evidence"`
	CreationGateEvidence           []GateEvidence       `json:"creation_gate_evidence"`
	CheckedAt                      string               `json:"checked_at"`
}

type TemplateEvidence struct {
	Role           string `json:"role"`
	PublicPath     string `json:"public_path"`
	SourceArtifact string `json:"source_artifact"`
	State          string `json:"state"`
}

type GateEvidence struct {
	Key                string `json:"key"`
	Required           bool   `json:"required"`
	BlocksRepoCreation bool   `json:"blocks_repo_creation"`
	EvidenceKind       string `json:"evidence_kind"`
	State              string `json:"state"`
}

type PublicSafetyEvidence struct {
	OK                      bool     `json:"ok"`
	CurrentOK               bool     `json:"current_ok"`
	HistoryOK               bool     `json:"history_ok"`
	CurrentFindingCount     int      `json:"current_finding_count"`
	HistoryFindingCount     int      `json:"history_finding_count"`
	EvidenceState           string   `json:"evidence_state"`
	ValidationCommands      []string `json:"validation_commands"`
	RawDetailsPublicAllowed bool     `json:"raw_details_public_allowed"`
}

type ContextPackEvidence struct {
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
