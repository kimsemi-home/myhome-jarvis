package connectors

const generatedConnectorPath = "generated/connectors.generated.json"

type Status struct {
	FixtureOnly             bool        `json:"fixture_only"`
	RealCredentialsAllowed  bool        `json:"real_credentials_allowed"`
	ExternalAPICallsAllowed bool        `json:"external_api_calls_allowed"`
	ConnectorCount          int         `json:"connector_count"`
	PlannedCount            int         `json:"planned_count"`
	FixtureModeCount        int         `json:"fixture_mode_count"`
	ReadOnlyOperationCount  int         `json:"read_only_operation_count"`
	ForbiddenOperationCount int         `json:"forbidden_operation_count"`
	GeneratedPath           string      `json:"generated_path"`
	Connectors              []Connector `json:"connectors"`
	Message                 string      `json:"message"`
	CheckedAt               string      `json:"checked_at"`
}

type Connector struct {
	Key                 string   `json:"key"`
	Label               string   `json:"label"`
	Category            string   `json:"category"`
	Status              string   `json:"status"`
	FixtureMode         bool     `json:"fixture_mode"`
	DataClasses         []string `json:"data_classes"`
	AllowedOperations   []string `json:"allowed_operations"`
	ForbiddenOperations []string `json:"forbidden_operations"`
	NextStep            string   `json:"next_step"`
}

type generatedPolicy struct {
	FixtureOnly             bool        `json:"fixture_only"`
	RealCredentialsAllowed  bool        `json:"real_credentials_allowed"`
	ExternalAPICallsAllowed bool        `json:"external_api_calls_allowed"`
	Connectors              []Connector `json:"connectors"`
}
