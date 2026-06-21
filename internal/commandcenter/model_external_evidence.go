package commandcenter

type ExternalEvidenceSummary struct {
	PublicSafe                  bool   `json:"public_safe"`
	EvidenceReady               bool   `json:"evidence_ready"`
	ExternalCollectionAllowed   bool   `json:"external_collection_allowed"`
	ForbiddenCollectionDisabled bool   `json:"forbidden_collection_disabled"`
	RawPayloadPublicAllowed     bool   `json:"raw_payload_public_allowed"`
	SourceCount                 int    `json:"source_count"`
	SourceClassCount            int    `json:"source_class_count"`
	ManifestPresent             bool   `json:"manifest_present"`
	ManifestRecordCount         int    `json:"manifest_record_count"`
	ArchiveSourceKey            string `json:"archive_source_key"`
	RepoSplitRecommendation     string `json:"repo_split_recommendation"`
	RepoCreationGate            string `json:"repo_creation_gate"`
	SplitTriggerCount           int    `json:"split_trigger_count"`
}
