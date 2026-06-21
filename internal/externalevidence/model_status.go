package externalevidence

type Status struct {
	PolicyPath                     string   `json:"policy_path"`
	SchemaVersion                  string   `json:"schema_version"`
	PublicSafe                     bool     `json:"public_safe"`
	ExternalCollectionAllowed      bool     `json:"external_collection_allowed"`
	CredentialsAllowed             bool     `json:"credentials_allowed"`
	CookiesAllowed                 bool     `json:"cookies_allowed"`
	RawPayloadPublicAllowed        bool     `json:"raw_payload_public_allowed"`
	PrivateRoot                    string   `json:"private_root"`
	ManifestPath                   string   `json:"manifest_path"`
	RawLayerPath                   string   `json:"raw_layer_path"`
	BronzeLayerPath                string   `json:"bronze_layer_path"`
	SilverLayerPath                string   `json:"silver_layer_path"`
	GoldLayerPath                  string   `json:"gold_layer_path"`
	ArchiveSourceKey               string   `json:"archive_source_key"`
	StorageArchiveSourcePath       string   `json:"storage_archive_source_path"`
	SourceCount                    int      `json:"source_count"`
	SourceClasses                  []string `json:"source_classes"`
	PreprocessingRules             []string `json:"preprocessing_rules"`
	ManifestPresent                bool     `json:"manifest_present"`
	ManifestRecordCount            int      `json:"manifest_record_count"`
	LatestCollectedAt              string   `json:"latest_collected_at,omitempty"`
	RepoSplitRecommendation        string   `json:"repo_split_recommendation"`
	FutureRepoCandidate            string   `json:"future_repo_candidate"`
	RepoCreationGate               string   `json:"repo_creation_gate"`
	SplitTriggerCount              int      `json:"split_trigger_count"`
	CurrentRepoResponsibilityCount int      `json:"current_repo_responsibility_count"`
	FutureRepoResponsibilityCount  int      `json:"future_repo_responsibility_count"`
	PublicRepoRules                []string `json:"public_repo_rules"`
	Commands                       []string `json:"commands"`
	CheckedAt                      string   `json:"checked_at"`
}
