package cicache

type Status struct {
	GraphPath                  string       `json:"graph_path"`
	WorkflowPath               string       `json:"workflow_path"`
	OK                         bool         `json:"ok"`
	PublicSafe                 bool         `json:"public_safe"`
	CachedUnitCount            int          `json:"cached_unit_count"`
	UncachedUnitCount          int          `json:"uncached_unit_count"`
	GeneratedArtifactCount     int          `json:"generated_artifact_count"`
	GeneratedCoverageRequired  bool         `json:"generated_coverage_required"`
	GeneratedCoverageOK        bool         `json:"generated_coverage_ok"`
	PublicSafetyNonSkippable   bool         `json:"public_safety_non_skippable"`
	InvalidCachedUnitCount     int          `json:"invalid_cached_unit_count"`
	WorkflowContractIssueCount int          `json:"workflow_contract_issue_count"`
	RawEvidencePublicAllowed   bool         `json:"raw_evidence_public_allowed"`
	PrivatePayloadsAllowed     bool         `json:"private_payloads_allowed"`
	CacheHitSkipsVerification  bool         `json:"cache_hit_skips_verification"`
	CacheMissRunsVerification  bool         `json:"cache_miss_runs_verification"`
	PushOnlyCacheSaveRequired  bool         `json:"push_only_cache_save_required"`
	Units                      []UnitStatus `json:"units"`
}

type UnitStatus struct {
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	Kind                        string `json:"kind"`
	CacheKey                    string `json:"cache_key,omitempty"`
	HashInputCount              int    `json:"hash_input_count"`
	GeneratedCoverageCount      int    `json:"generated_coverage_count"`
	CacheRestoreConfigured      bool   `json:"cache_restore_configured"`
	CacheHitSkipsVerification   bool   `json:"cache_hit_skips_verification"`
	CacheMissRunsVerification   bool   `json:"cache_miss_runs_verification"`
	PushOnlyCacheSaveConfigured bool   `json:"push_only_cache_save_configured"`
	GeneratedArtifactCoverageOK bool   `json:"generated_artifact_coverage_ok"`
	PublicSafetyNonSkippable    bool   `json:"public_safety_non_skippable,omitempty"`
	Valid                       bool   `json:"valid"`
}

type graphFile struct {
	GeneratedArtifacts []string    `json:"generated_artifacts"`
	Units              []graphUnit `json:"units"`
}

type graphUnit struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Kind       string   `json:"kind"`
	Cache      string   `json:"cache"`
	HashInputs []string `json:"hash_inputs"`
}
