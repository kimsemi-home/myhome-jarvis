package localfinancereadiness

const (
	ManifestSchema = "myhome.local-finance-readiness-manifest/v1"
	PlanSchema     = "myhome.connection-readiness-plan/v1"
)

type Ref struct {
	Component      string `json:"component"`
	Path           string `json:"path"`
	ArtifactSHA256 string `json:"artifact_sha256"`
	PlanHash       string `json:"plan_hash"`
}

type Stage struct {
	Position  int      `json:"position"`
	Component string   `json:"component"`
	Day       int      `json:"day"`
	Hour      int      `json:"hour"`
	Minute    int      `json:"minute"`
	DependsOn []string `json:"depends_on"`
	Action    string   `json:"action"`
}

type Manifest struct {
	SchemaVersion          string  `json:"schema_version"`
	ExecutionMode          string  `json:"execution_mode"`
	CredentialsRead        bool    `json:"credentials_read"`
	ExternalNetworkEnabled bool    `json:"external_network_enabled"`
	ExternalWritesEnabled  bool    `json:"external_writes_enabled"`
	InstallAllowed         bool    `json:"install_allowed"`
	Timezone               string  `json:"timezone"`
	Plans                  []Ref   `json:"plans"`
	Stages                 []Stage `json:"stages"`
	AggregateHash          string  `json:"aggregate_hash"`
}

type Plan struct {
	SchemaVersion          string   `json:"schema_version"`
	Component              string   `json:"component"`
	ExecutionMode          string   `json:"execution_mode"`
	CredentialsRead        bool     `json:"credentials_read"`
	ExternalNetworkEnabled bool     `json:"external_network_enabled"`
	ExternalWritesEnabled  bool     `json:"external_writes_enabled"`
	InstallAllowed         bool     `json:"install_allowed"`
	PublicConfigSource     string   `json:"public_config_source"`
	KeychainHandles        []Handle `json:"keychain_handles"`
	OAuthScopes            []string `json:"oauth_scopes"`
	OfficialHosts          []string `json:"official_hosts"`
	LocalDependencies      []string `json:"local_dependencies"`
	EnabledCollectors      []string `json:"enabled_collectors"`
	DeferredCollectors     []string `json:"deferred_collectors"`
	Schedule               Schedule `json:"monthly_schedule"`
	Launchd                Launchd  `json:"launchd"`
	Checks                 []string `json:"checks"`
	TemplateHash           string   `json:"template_hash"`
	PlanHash               string   `json:"plan_hash"`
}
