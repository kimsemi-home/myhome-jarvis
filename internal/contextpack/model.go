package contextpack

type Policy struct {
	Context                        string              `json:"context"`
	Version                        string              `json:"version"`
	GeneratedArtifact              string              `json:"generated_artifact"`
	PackID                         string              `json:"pack_id"`
	UpstreamCompatibilityVersion   string              `json:"upstream_compatibility_version"`
	OntologyVersion                string              `json:"ontology_version"`
	DeclarationPath                string              `json:"declaration_path"`
	PublicStatusRedacted           bool                `json:"public_status_redacted"`
	RawPrivateContextPublicAllowed bool                `json:"raw_private_context_public_allowed"`
	MissionSource                  string              `json:"mission_source"`
	BoundedContextSource           string              `json:"bounded_context_source"`
	SplitCriteria                  []SplitCriterion    `json:"split_criteria"`
	ExportedArtifacts              []Artifact          `json:"exported_artifacts"`
	AuthorityContract              AuthorityContract   `json:"authority_contract"`
	SecurityContract               SecurityContract    `json:"security_contract"`
	VerificationProfile            VerificationProfile `json:"verification_profile"`
	RequiredDeclarationFields      []string            `json:"required_declaration_fields"`
	ForbiddenPublicFields          []string            `json:"forbidden_public_fields"`
	Commands                       []string            `json:"commands"`
}

type SplitCriterion struct {
	Key     string `json:"key"`
	Meaning string `json:"meaning"`
}

type Artifact struct {
	Role    string `json:"role"`
	Path    string `json:"path"`
	Version string `json:"version"`
}
