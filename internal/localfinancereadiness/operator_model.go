package localfinancereadiness

type OperatorSchedule struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type OperatorStage struct {
	Component        string `json:"component"`
	BinaryPath       string `json:"binary_path"`
	WorkingDirectory string `json:"working_directory"`
	DueDay           int    `json:"due_day"`
	BinarySHA256     string `json:"binary_sha256"`
}

type OperatorLaunchd struct {
	TemplatePath     string   `json:"template_path"`
	Label            string   `json:"label"`
	ProgramArguments []string `json:"program_arguments"`
	TemplateHash     string   `json:"template_hash"`
}

type OperatorPlan struct {
	SchemaVersion          string           `json:"schema_version"`
	Component              string           `json:"component"`
	ExecutionMode          string           `json:"execution_mode"`
	CredentialsRead        bool             `json:"credentials_read"`
	ExternalNetworkEnabled bool             `json:"external_network_enabled"`
	ExternalWritesEnabled  bool             `json:"external_writes_enabled"`
	InstallAllowed         bool             `json:"install_allowed"`
	PublicConfigSource     string           `json:"public_config_source"`
	KeychainService        string           `json:"keychain_service"`
	APITokenAccount        string           `json:"api_token_account"`
	Timezone               string           `json:"timezone"`
	DailySchedule          OperatorSchedule `json:"daily_schedule"`
	Stages                 []OperatorStage  `json:"stages"`
	Launchd                OperatorLaunchd  `json:"launchd"`
	Checks                 []string         `json:"checks"`
	PlanHash               string           `json:"plan_hash"`
}
