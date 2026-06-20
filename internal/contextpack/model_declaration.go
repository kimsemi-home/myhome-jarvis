package contextpack

type Declaration struct {
	PackID                       string     `json:"pack_id"`
	ContextPackVersion           string     `json:"context_pack_version"`
	UpstreamCompatibilityVersion string     `json:"upstream_compatibility_version"`
	OntologyVersion              string     `json:"ontology_version"`
	AuthorityContractVersion     string     `json:"authority_contract_version"`
	SecurityContractVersion      string     `json:"security_contract_version"`
	VerificationProfile          string     `json:"verification_profile"`
	SSOTArtifactVersions         []Artifact `json:"ssot_artifact_versions"`
}

type VerifyResult struct {
	DeclarationPath      string   `json:"declaration_path"`
	Valid                bool     `json:"valid"`
	DriftCount           int      `json:"drift_count"`
	MissingFieldCount    int      `json:"missing_field_count"`
	MissingArtifactCount int      `json:"missing_artifact_count"`
	StaleVersionCount    int      `json:"stale_version_count"`
	ForbiddenValueCount  int      `json:"forbidden_value_count"`
	Findings             []string `json:"findings"`
	CheckedAt            string   `json:"checked_at"`
}
