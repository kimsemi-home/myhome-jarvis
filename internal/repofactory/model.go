package repofactory

type Policy struct {
	Context                          string         `json:"context"`
	Version                          string         `json:"version"`
	GeneratedArtifact                string         `json:"generated_artifact"`
	TargetOwner                      string         `json:"target_owner"`
	PublicRepoDefault                bool           `json:"public_repo_default"`
	CodexProjectRequired             bool           `json:"codex_project_required"`
	RepoCreationAllowedWithoutReview bool           `json:"repo_creation_allowed_without_review"`
	PublicSafetyEvidenceRequired     bool           `json:"public_safety_evidence_required"`
	AuthorityReviewRequired          bool           `json:"authority_review_required"`
	PrivateAssetsPublicAllowed       bool           `json:"private_assets_public_allowed"`
	LocalPathsPublicAllowed          bool           `json:"local_paths_public_allowed"`
	TemplateFiles                    []TemplateFile `json:"template_files"`
	CreationGates                    []CreationGate `json:"creation_gates"`
	BootstrapChecklist               []string       `json:"bootstrap_checklist"`
	AllowedPublicPathPrefixes        []string       `json:"allowed_public_path_prefixes"`
	ForbiddenPublicFragments         []string       `json:"forbidden_public_fragments"`
	PublicSummaryFields              []string       `json:"public_summary_fields"`
	Commands                         []string       `json:"commands"`
}

type TemplateFile struct {
	Role           string `json:"role"`
	Path           string `json:"path"`
	SourceArtifact string `json:"source_artifact"`
	Purpose        string `json:"purpose"`
}

type CreationGate struct {
	Key                string `json:"key"`
	Required           bool   `json:"required"`
	BlocksRepoCreation bool   `json:"blocks_repo_creation"`
	Evidence           string `json:"evidence"`
}

type Status struct {
	PolicyPath                     string   `json:"policy_path"`
	TemplateFileCount              int      `json:"template_file_count"`
	CreationGateCount              int      `json:"creation_gate_count"`
	BootstrapCheckCount            int      `json:"bootstrap_check_count"`
	AuthorityReviewRequired        bool     `json:"authority_review_required"`
	PublicSafetyEvidenceRequired   bool     `json:"public_safety_evidence_required"`
	CodexProjectRequired           bool     `json:"codex_project_required"`
	CreationAllowedWithoutReview   bool     `json:"creation_allowed_without_review"`
	PublicSafe                     bool     `json:"public_safe"`
	MissingTemplateRoleCount       int      `json:"missing_template_role_count"`
	MissingCreationGateCount       int      `json:"missing_creation_gate_count"`
	ForbiddenTemplateValueCount    int      `json:"forbidden_template_value_count"`
	TemplateRoles                  []string `json:"template_roles"`
	CreationGateKeys               []string `json:"creation_gate_keys"`
	BootstrapChecklistReady        bool     `json:"bootstrap_checklist_ready"`
	GeneratedCIPresent             bool     `json:"generated_ci_present"`
	SecurityScanPresent            bool     `json:"security_scan_present"`
	PrivateDataPolicyPresent       bool     `json:"private_data_policy_present"`
	BootstrapChecklistPresent      bool     `json:"bootstrap_checklist_present"`
	CodexProjectTemplatePresent    bool     `json:"codex_project_template_present"`
	RepoCreationBlockedUntilReview bool     `json:"repo_creation_blocked_until_review"`
	CheckedAt                      string   `json:"checked_at"`
}
