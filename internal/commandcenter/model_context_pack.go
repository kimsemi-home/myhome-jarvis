package commandcenter

type ContextPackSummary struct {
	PackID                        string `json:"pack_id"`
	Version                       string `json:"version"`
	UpstreamCompatibilityVersion  string `json:"upstream_compatibility_version"`
	OntologyVersion               string `json:"ontology_version"`
	PublicSafe                    bool   `json:"public_safe"`
	SplitCriteriaCount            int    `json:"split_criteria_count"`
	ExportedArtifactCount         int    `json:"exported_artifact_count"`
	AuthorityContractVersion      string `json:"authority_contract_version"`
	SecurityContractVersion       string `json:"security_contract_version"`
	VerificationProfile           string `json:"verification_profile"`
	VerificationRequiredUnitCount int    `json:"verification_required_unit_count"`
}
