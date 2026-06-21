package externalevidence

type CollectReport struct {
	PolicyPath              string          `json:"policy_path"`
	PublicSafe              bool            `json:"public_safe"`
	RawPayloadPublicAllowed bool            `json:"raw_payload_public_allowed"`
	CollectionRunState      string          `json:"collection_run_state"`
	SourceCount             int             `json:"source_count"`
	CollectedCount          int             `json:"collected_count"`
	CachedCount             int             `json:"cached_count"`
	FailedCount             int             `json:"failed_count"`
	ManifestPath            string          `json:"manifest_path"`
	RawLayerPath            string          `json:"raw_layer_path"`
	BronzeLayerPath         string          `json:"bronze_layer_path"`
	SilverLayerPath         string          `json:"silver_layer_path"`
	GoldLayerPath           string          `json:"gold_layer_path"`
	ArchiveSourceKey        string          `json:"archive_source_key"`
	RepoSplitRecommendation string          `json:"repo_split_recommendation"`
	RepoCreationGate        string          `json:"repo_creation_gate"`
	Results                 []CollectResult `json:"results"`
	CheckedAt               string          `json:"checked_at"`
}

type CollectResult struct {
	SourceKey      string `json:"source_key"`
	SourceClass    string `json:"source_class"`
	State          string `json:"state"`
	EvidenceRef    string `json:"evidence_ref,omitempty"`
	RawSHA256      string `json:"raw_sha256,omitempty"`
	PayloadBytes   int    `json:"payload_bytes,omitempty"`
	HTTPStatus     int    `json:"http_status,omitempty"`
	ErrorCategory  string `json:"error_category,omitempty"`
	FreshnessHours int    `json:"freshness_hours"`
}
