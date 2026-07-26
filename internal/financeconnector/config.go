package financeconnector

const (
	ModeFixture = "fixture"
	ModeMyData  = "mydata"
)

// RuntimeConfig carries only secret references; resolved values stay server-side.
type RuntimeConfig struct {
	Mode              string             `json:"mode"`
	Provider          string             `json:"provider"`
	BaseURL           string             `json:"base_url,omitempty"`
	AuthBaseURL       string             `json:"auth_base_url,omitempty"`
	AuthTokenBaseURL  string             `json:"auth_token_base_url,omitempty"`
	APIType           string             `json:"api_type,omitempty"`
	FromDate          string             `json:"from_date,omitempty"`
	ToDate            string             `json:"to_date,omitempty"`
	IncludeCards      bool               `json:"include_cards"`
	Connections       []ConnectionConfig `json:"-"`
	LiveModeRequested bool               `json:"live_mode_requested"`
	ExternalCallsLive bool               `json:"external_calls_live"`
}

type ConnectionConfig struct {
	Owner         string            `json:"owner"`
	Credential    CredentialBinding `json:"-"`
	OrgCode       string            `json:"-"`
	AccountNumber string            `json:"-"`
	CardIDs       []string          `json:"-"`
}

// CredentialBinding keeps provider authentication and data access on one
// server-side reference. A connection cannot select them independently.
type CredentialBinding struct {
	OnePasswordRef string `json:"-"`
}
