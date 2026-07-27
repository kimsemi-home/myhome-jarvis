package financeconnector

import (
	"context"
	"strings"
)

type PreflightConnection struct {
	Owner       string `json:"owner"`
	BundleReady bool   `json:"bundle_ready"`
}

type PreflightReport struct {
	Mode                   string                `json:"mode"`
	Provider               string                `json:"provider"`
	ExternalCallsEnabled   bool                  `json:"external_calls_enabled"`
	ConsentReady           bool                  `json:"consent_ready"`
	OnePasswordSession     bool                  `json:"onepassword_session"`
	DataBaseURLConfigured  bool                  `json:"data_base_url_configured"`
	AuthBaseURLConfigured  bool                  `json:"auth_base_url_configured"`
	AuthTokenURLConfigured bool                  `json:"auth_token_url_configured"`
	ConnectionCount        int                   `json:"connection_count"`
	Connections            []PreflightConnection `json:"connections"`
	FailureReasons         []string              `json:"failure_reasons,omitempty"`
	Ready                  bool                  `json:"ready"`
}

func Preflight(ctx context.Context, root string, config RuntimeConfig) (PreflightReport, error) {
	report := PreflightReport{
		Mode: config.Mode, Provider: config.Provider,
		ExternalCallsEnabled:   config.ExternalCallsLive,
		DataBaseURLConfigured:  strings.TrimSpace(config.BaseURL) != "",
		AuthBaseURLConfigured:  strings.TrimSpace(config.AuthBaseURL) != "",
		AuthTokenURLConfigured: strings.TrimSpace(config.AuthTokenBaseURL) != "",
		ConnectionCount:        len(config.Connections),
		OnePasswordSession:     len(config.Connections) > 0,
	}
	for _, connection := range config.Connections {
		report.Connections = append(report.Connections, PreflightConnection{Owner: connection.Owner})
	}
	if config.Mode != ModeMyData {
		report.FailureReasons = append(report.FailureReasons, "mydata_mode_required")
	}
	if !report.DataBaseURLConfigured {
		report.FailureReasons = append(report.FailureReasons, "data_base_url_missing")
	}
	if err := RequireActiveConsent(root); err == nil {
		report.ConsentReady = true
	} else {
		report.FailureReasons = append(report.FailureReasons, "consent_not_ready")
	}
	for index, connection := range config.Connections {
		credential, err := ResolveOnePasswordCLI(ctx, connection.Credential.OnePasswordRef, config.Provider)
		if err != nil {
			report.OnePasswordSession = false
			report.FailureReasons = append(report.FailureReasons, "credential_bundle_unavailable")
			continue
		}
		if err := validatePreflightCredential(config, credential); err != nil {
			report.FailureReasons = append(report.FailureReasons, "credential_bundle_invalid")
			continue
		}
		report.Connections[index].BundleReady = true
	}
	report.Ready = report.Mode == ModeMyData && report.ExternalCallsEnabled &&
		report.ConsentReady && report.OnePasswordSession && report.DataBaseURLConfigured &&
		report.AuthBaseURLConfigured && report.AuthTokenURLConfigured &&
		len(report.Connections) > 0 && allPreflightBundlesReady(report.Connections)
	if !report.Ready {
		return report, errFinancePreflightFailed
	}
	return report, nil
}
