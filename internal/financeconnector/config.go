package financeconnector

import (
	"fmt"
	"strings"
)

const (
	ModeFixture = "fixture"
	ModeMyData  = "mydata"
)

// RuntimeConfig carries only a 1Password reference. The resolved value must
// be injected into the server process by the deployment environment; Flutter
// and public status responses never receive it.
type RuntimeConfig struct {
	Mode              string `json:"mode"`
	Provider          string `json:"provider"`
	OnePasswordRef    string `json:"onepassword_ref,omitempty"`
	LiveModeRequested bool   `json:"live_mode_requested"`
	ExternalCallsLive bool   `json:"external_calls_live"`
}

func RuntimeConfigFromEnv(getenv func(string) string) (RuntimeConfig, error) {
	mode := strings.TrimSpace(getenv("MYHOME_FINANCE_MODE"))
	if mode == "" {
		mode = ModeFixture
	}
	provider := strings.TrimSpace(getenv("MYHOME_MYDATA_PROVIDER"))
	if provider == "" {
		provider = ProviderTossMyData
	}
	ref := strings.TrimSpace(getenv("MYHOME_FINANCE_1PASSWORD_REF"))
	config := RuntimeConfig{
		Mode:              mode,
		Provider:          provider,
		OnePasswordRef:    ref,
		LiveModeRequested: mode == ModeMyData,
		ExternalCallsLive: false,
	}
	if mode != ModeFixture && mode != ModeMyData {
		return RuntimeConfig{}, fmt.Errorf("unsupported finance mode %q", mode)
	}
	if mode == ModeMyData && !strings.HasPrefix(ref, "op://") {
		return RuntimeConfig{}, fmt.Errorf("mydata mode requires a 1Password op:// reference")
	}
	return config, nil
}
