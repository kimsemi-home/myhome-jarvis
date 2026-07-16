package localfinanceevidence

type ShortsReport struct {
	SchemaVersion         string         `json:"schema_version"`
	ExecutionMode         string         `json:"execution_mode"`
	LoopbackOnly          bool           `json:"loopback_only"`
	CredentialsRead       bool           `json:"credentials_read"`
	ExternalNetwork       bool           `json:"external_network"`
	ExternalWrites        bool           `json:"external_writes"`
	PlanHash              string         `json:"plan_hash"`
	ContentHash           string         `json:"content_hash"`
	OAuth                 ShortsOAuth    `json:"oauth_token_boundary"`
	Channel               ShortsChannel  `json:"channel_identity_boundary"`
	SessionStateReplay    bool           `json:"session_state_replay"`
	SessionLocationPinned bool           `json:"session_location_pinned"`
	NotificationsDisabled bool           `json:"notifications_disabled"`
	First                 ShortsReceipt  `json:"first_execution"`
	Replay                ShortsReceipt  `json:"idempotent_replay"`
	Emulator              ShortsEmulator `json:"emulator_metrics"`
	Checks                []string       `json:"checks"`
	ReportHash            string         `json:"report_hash"`
}

type ShortsOAuth struct {
	SchemaVersion             string `json:"schema_version"`
	AuthorizationCodeExchange bool   `json:"authorization_code_exchange"`
	RefreshTokenExchange      bool   `json:"refresh_token_exchange"`
	ScopeContractValidated    bool   `json:"scope_contract_validated"`
	OfficialOriginPinned      bool   `json:"official_origin_pinned"`
	RedirectRejected          bool   `json:"redirect_rejected"`
	OversizedResponseRejected bool   `json:"oversized_response_rejected"`
}

type ShortsChannel struct {
	SchemaVersion        string `json:"schema_version"`
	Method               string `json:"method"`
	Path                 string `json:"path"`
	Query                string `json:"query"`
	ExactlyOneChannel    bool   `json:"exactly_one_channel"`
	BindingHashMatched   bool   `json:"binding_hash_matched"`
	RawIdentityPersisted bool   `json:"raw_identity_persisted"`
	OfficialOriginPinned bool   `json:"official_origin_pinned"`
}
