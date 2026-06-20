package financeconsent

type Policy struct {
	Context                       string   `json:"context"`
	Version                       string   `json:"version"`
	GeneratedArtifact             string   `json:"generated_artifact"`
	PrivateConsentLedger          string   `json:"private_consent_ledger"`
	AppendOnly                    bool     `json:"append_only"`
	PublicStatusRedacted          bool     `json:"public_status_redacted"`
	FinanceMode                   string   `json:"finance_mode"`
	ReadOnly                      bool     `json:"read_only"`
	ReviewOnly                    bool     `json:"review_only"`
	FixtureOnlyUntilConsent       bool     `json:"fixture_only_until_consent"`
	ExternalWritesAllowed         bool     `json:"external_writes_allowed"`
	TransferActionsAllowed        bool     `json:"transfer_actions_allowed"`
	PaymentActionsAllowed         bool     `json:"payment_actions_allowed"`
	TradeActionsAllowed           bool     `json:"trade_actions_allowed"`
	CardActionsAllowed            bool     `json:"card_actions_allowed"`
	RealConnectorRequiresConsent  bool     `json:"real_connector_requires_active_consent"`
	SpouseScopeRequiresConsent    bool     `json:"spouse_scope_requires_active_consent"`
	HouseholdScopeRequiresConsent bool     `json:"household_scope_requires_active_consent"`
	ConsentKinds                  []string `json:"consent_kinds"`
	ConsentStatuses               []string `json:"consent_statuses"`
	ReviewStatuses                []string `json:"review_statuses"`
	AuthorityProfiles             []string `json:"authority_profiles"`
	RequiredFields                []string `json:"required_fields"`
	AllowedEvidencePrefixes       []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields           []string `json:"public_summary_fields"`
	ForbiddenPublicFields         []string `json:"forbidden_public_fields"`
	Commands                      []string `json:"commands"`
}
