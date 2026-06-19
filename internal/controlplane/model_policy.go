package controlplane

const PolicyRelativePath = "generated/control_plane.generated.json"

type Policy struct {
	Context                    string   `json:"context"`
	Version                    string   `json:"version"`
	GeneratedArtifact          string   `json:"generated_artifact"`
	VerifierGeneratedArtifact  string   `json:"verifier_generated_artifact"`
	PrivateManifestLedger      string   `json:"private_manifest_ledger"`
	ManifestRequired           bool     `json:"manifest_required"`
	AppendOnly                 bool     `json:"append_only"`
	PublicStatusRedacted       bool     `json:"public_status_redacted"`
	RawRationalePublicAllowed  bool     `json:"raw_rationale_public_allowed"`
	VerifierSeparationRequired bool     `json:"verifier_separation_required"`
	VerificationCommand        string   `json:"verification_command"`
	VerifierChecks             []string `json:"verifier_checks"`
	MinLeaseSeconds            int      `json:"min_lease_seconds"`
	MaxLeaseSeconds            int      `json:"max_lease_seconds"`
	AllowedDecisionKinds       []string `json:"allowed_decision_kinds"`
	AllowedAuthorityProfiles   []string `json:"allowed_authority_profiles"`
	AllowedLeaseStatuses       []string `json:"allowed_lease_statuses"`
	RequiredFields             []string `json:"required_fields"`
	AllowedEvidencePrefixes    []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields        []string `json:"public_summary_fields"`
	ForbiddenPublicFields      []string `json:"forbidden_public_fields"`
	Commands                   []string `json:"commands"`
}
