package translation

const PolicyRelativePath = "generated/translation.generated.json"

type Policy struct {
	Context                 string   `json:"context"`
	Version                 string   `json:"version"`
	GeneratedArtifact       string   `json:"generated_artifact"`
	PrivateLossLedger       string   `json:"private_loss_ledger"`
	PrivateManifestRoot     string   `json:"private_manifest_root"`
	ManifestRequired        bool     `json:"manifest_required"`
	PublicStatusRedacted    bool     `json:"public_status_redacted"`
	RawLossPublicAllowed    bool     `json:"raw_loss_public_allowed"`
	AllowedContexts         []string `json:"allowed_contexts"`
	RequiredManifestFields  []string `json:"required_manifest_fields"`
	LossLevels              []string `json:"loss_levels"`
	AllowedLossCategories   []string `json:"allowed_loss_categories"`
	ForbiddenLossCategories []string `json:"forbidden_loss_categories"`
	AllowedEvidencePrefixes []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields     []string `json:"public_summary_fields"`
	ForbiddenPublicFields   []string `json:"forbidden_public_fields"`
	Commands                []string `json:"commands"`
}
