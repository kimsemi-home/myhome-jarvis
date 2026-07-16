package localfinanceevidence

type ShortsActivationReport struct {
	SchemaVersion              string                   `json:"schema_version"`
	ExecutionMode              string                   `json:"execution_mode"`
	LoopbackOnly               bool                     `json:"loopback_only"`
	CredentialsRead            bool                     `json:"credentials_read"`
	ExternalNetwork            bool                     `json:"external_network"`
	ExternalWrites             bool                     `json:"external_writes"`
	RuntimeEntrypointsInactive bool                     `json:"runtime_entrypoints_inactive"`
	Callback                   ShortsActivationCallback `json:"callback_receiver"`
	Keychain                   ShortsActivationKeychain `json:"keychain_adapter"`
	Checks                     []string                 `json:"checks"`
	ReportHash                 string                   `json:"report_hash"`
}

type ShortsActivationCallback struct {
	Network                  string `json:"network"`
	Host                     string `json:"host"`
	PortStrategy             string `json:"port_strategy"`
	Path                     string `json:"path"`
	InvalidAttemptsRejected  int    `json:"invalid_attempts_rejected"`
	ValidCallbackConsumed    bool   `json:"valid_callback_consumed"`
	DuplicateCallbackDenied  bool   `json:"duplicate_callback_denied"`
	CanceledReceiverClosed   bool   `json:"canceled_receiver_closed"`
	RawStateReported         bool   `json:"raw_state_reported"`
	RawCodeReported          bool   `json:"raw_code_reported"`
	ExternalNetworkRequested bool   `json:"external_network_requested"`
}

type ShortsActivationKeychain struct {
	Executable                  string `json:"executable"`
	FakeRunner                  bool   `json:"fake_runner"`
	ActualKeychainExecuted      bool   `json:"actual_keychain_executed"`
	CommandsValidated           int    `json:"commands_validated"`
	InteractiveProvisioning     bool   `json:"interactive_provisioning"`
	SecretValueInArguments      bool   `json:"secret_value_in_arguments"`
	DefaultPermitDenied         bool   `json:"default_permit_denied"`
	ExpiredPermitDenied         bool   `json:"expired_permit_denied"`
	UnlistedReferenceDenied     bool   `json:"unlisted_reference_denied"`
	ExternalNetworkPermitDenied bool   `json:"external_network_permit_denied"`
	ReadinessPlanBound          bool   `json:"readiness_plan_bound"`
	RawCredentialReported       bool   `json:"raw_credential_reported"`
	ReturnedMaterialZeroed      bool   `json:"returned_material_zeroed"`
}
