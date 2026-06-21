package externalevidence

type Policy struct {
	SchemaVersion                    string              `json:"schema_version"`
	PublicSafe                       bool                `json:"public_safe"`
	ExternalNetworkCollectionAllowed bool                `json:"external_network_collection_allowed"`
	CredentialsAllowed               bool                `json:"credentials_allowed"`
	CookiesAllowed                   bool                `json:"cookies_allowed"`
	RawPayloadPublicAllowed          bool                `json:"raw_payload_public_allowed"`
	PrivateRoot                      string              `json:"private_root"`
	ManifestPath                     string              `json:"manifest_path"`
	RawLayerPath                     string              `json:"raw_layer_path"`
	BronzeLayerPath                  string              `json:"bronze_layer_path"`
	SilverLayerPath                  string              `json:"silver_layer_path"`
	GoldLayerPath                    string              `json:"gold_layer_path"`
	ArchiveSourceKey                 string              `json:"archive_source_key"`
	StorageArchiveSourcePath         string              `json:"storage_archive_source_path"`
	CollectionMaxBytes               int64               `json:"collection_max_bytes"`
	SourceClasses                    []string            `json:"source_classes"`
	SourceDescriptors                []SourceDescriptor  `json:"source_descriptors"`
	PreprocessingRules               []string            `json:"preprocessing_rules"`
	RepoSplitAssessment              RepoSplitAssessment `json:"repo_split_assessment"`
	Commands                         []string            `json:"commands"`
}

type SourceDescriptor struct {
	Key            string `json:"key"`
	Class          string `json:"class"`
	Method         string `json:"method"`
	URL            string `json:"url"`
	FreshnessHours int    `json:"freshness_hours"`
	Preprocess     string `json:"preprocess"`
}

type RepoSplitAssessment struct {
	Recommendation              string   `json:"recommendation"`
	FutureRepoCandidate         string   `json:"future_repo_candidate"`
	CreationGate                string   `json:"creation_gate"`
	CurrentRepoResponsibilities []string `json:"current_repo_responsibilities"`
	FutureRepoResponsibilities  []string `json:"future_repo_responsibilities"`
	SplitTriggers               []string `json:"split_triggers"`
	PublicRepoRules             []string `json:"public_repo_rules"`
}
