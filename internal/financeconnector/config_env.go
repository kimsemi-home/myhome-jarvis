package financeconnector

import (
	"fmt"
	"strings"
)

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
	connections, err := parseConnections(getenv("MYHOME_FINANCE_CONNECTIONS_JSON"))
	if err != nil {
		return RuntimeConfig{}, err
	}
	if len(connections) == 0 && ref != "" {
		connections = []ConnectionConfig{{
			Owner: "household", Credential: CredentialBinding{OnePasswordRef: ref},
			OrgCode:       strings.TrimSpace(getenv("MYHOME_MYDATA_ORG_CODE")),
			AccountNumber: strings.TrimSpace(getenv("MYHOME_MYDATA_ACCOUNT_NUM")),
		}}
	}
	config := RuntimeConfig{
		Mode: mode, Provider: provider,
		BaseURL: strings.TrimRight(strings.TrimSpace(getenv("MYHOME_MYDATA_BASE_URL")), "/"),
		AuthBaseURL: strings.TrimRight(
			strings.TrimSpace(getenv("MYHOME_MYDATA_AUTH_BASE_URL")), "/",
		),
		APIType:  strings.TrimSpace(getenv("MYHOME_MYDATA_API_TYPE")),
		FromDate: strings.TrimSpace(getenv("MYHOME_MYDATA_FROM_DATE")),
		ToDate:   strings.TrimSpace(getenv("MYHOME_MYDATA_TO_DATE")),
		IncludeCards: !strings.EqualFold(
			strings.TrimSpace(getenv("MYHOME_MYDATA_INCLUDE_CARDS")), "false",
		),
		Connections: connections, LiveModeRequested: mode == ModeMyData,
		ExternalCallsLive: mode == ModeMyData && strings.EqualFold(
			strings.TrimSpace(getenv("MYHOME_FINANCE_ALLOW_EXTERNAL")), "true",
		),
	}
	if config.APIType == "" {
		config.APIType = "2"
	}
	if config.AuthBaseURL == "" {
		config.AuthBaseURL = "https://mydata.cert.toss.im"
	}
	return validateRuntimeConfig(config)
}

func validateRuntimeConfig(config RuntimeConfig) (RuntimeConfig, error) {
	if config.Mode != ModeFixture && config.Mode != ModeMyData {
		return RuntimeConfig{}, fmt.Errorf("unsupported finance mode %q", config.Mode)
	}
	if config.Mode == ModeMyData && (len(config.Connections) == 0 || !allConnectionsHaveRefs(config.Connections)) {
		return RuntimeConfig{}, fmt.Errorf("mydata mode requires connection 1Password op:// references")
	}
	if config.Mode == ModeMyData && !config.ExternalCallsLive {
		return RuntimeConfig{}, fmt.Errorf("mydata mode requires MYHOME_FINANCE_ALLOW_EXTERNAL=true")
	}
	return config, nil
}
