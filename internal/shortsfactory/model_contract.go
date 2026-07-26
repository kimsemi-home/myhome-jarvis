package shortsfactory

type Contract struct {
	SchemaVersion                     string   `json:"schema_version"`
	ContractID                        string   `json:"contract_id"`
	PublicSafe                        bool     `json:"public_safe"`
	LogicalChannelSlots               int      `json:"logical_channel_slots"`
	ApprovalAuthority                 string   `json:"approval_authority"`
	CriteriaVersion                   string   `json:"criteria_version"`
	MinimumIndependentSourcesPerClaim int      `json:"minimum_independent_sources_per_claim"`
	MinimumPrimarySourcesPerClaim     int      `json:"minimum_primary_sources_per_claim"`
	MaximumAPIDataRevalidationDays    int      `json:"maximum_api_data_revalidation_days"`
	DefaultUploadVisibility           string   `json:"default_upload_visibility"`
	ExternalWriteDefault              string   `json:"external_write_default"`
	YouTubeConsentReceiptRequired     bool     `json:"youtube_action_consent_receipt_required"`
	RequiredCriteria                  []string `json:"required_criteria"`
	ReleasedOpenLoopSteps             []string `json:"released_open_loop_steps"`
	LinkedPublicContracts             []string `json:"linked_public_contracts"`
	PrivateRuntimeRepository          string   `json:"private_runtime_repository"`
}
