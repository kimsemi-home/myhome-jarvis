package shortsfactory

type GateRequest struct {
	SchemaVersion      string           `json:"schema_version"`
	JobID              string           `json:"job_id"`
	LogicalChannelSlot string           `json:"logical_channel_slot"`
	State              string           `json:"state"`
	ClaimEvidence      ClaimEvidence    `json:"claim_evidence"`
	Originality        Originality      `json:"originality"`
	Rights             Rights           `json:"rights"`
	SyntheticContent   SyntheticContent `json:"synthetic_content"`
	Privacy            Privacy          `json:"privacy"`
	YouTubeConsent     YouTubeConsent   `json:"youtube_consent"`
	InputHashes        InputHashes      `json:"input_hashes"`
}

type ClaimEvidence struct {
	IndependentSources       int  `json:"independent_sources"`
	PrimarySources           int  `json:"primary_sources"`
	FreshSources             int  `json:"fresh_sources"`
	UnresolvedContradictions int  `json:"unresolved_contradictions"`
	ContradictionReviewed    bool `json:"contradiction_reviewed"`
	UncertaintyDisclosed     bool `json:"uncertainty_disclosed"`
}

type Originality struct {
	OriginalScript        bool `json:"original_script"`
	OriginalAnalysis      bool `json:"original_analysis"`
	CrossChannelDuplicate bool `json:"cross_channel_duplicate"`
	TemplateOnly          bool `json:"template_only"`
}

type Rights struct {
	AllAssetsCleared bool `json:"all_assets_cleared"`
	EvidenceCount    int  `json:"evidence_count"`
}

type SyntheticContent struct {
	RealisticOrMeaningful bool `json:"realistic_or_meaningful"`
	DisclosureRequired    bool `json:"disclosure_required"`
	DisclosurePlanned     bool `json:"disclosure_planned"`
}
