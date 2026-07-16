package localfinanceevidence

import "errors"

type ShortsActivationBrowser struct {
	Executable                    string `json:"executable"`
	FakeRunner                    bool   `json:"fake_runner"`
	ActualBrowserLaunched         bool   `json:"actual_browser_launched"`
	CommandsValidated             int    `json:"commands_validated"`
	BoundBeforeAuthorization      bool   `json:"bound_before_authorization"`
	AuthorizationOriginPinned     bool   `json:"authorization_origin_pinned"`
	RedirectBound                 bool   `json:"redirect_bound"`
	SystemBrowserOnly             bool   `json:"system_browser_only"`
	ExternalNetworkPermitRequired bool   `json:"external_network_permit_required"`
	DefaultPermitDenied           bool   `json:"default_permit_denied"`
	StateValidated                bool   `json:"state_validated"`
	WrongStateRejected            bool   `json:"wrong_state_rejected"`
	PKCEExchangePlanBound         bool   `json:"pkce_exchange_plan_bound"`
	TokenExchangePlanOnly         bool   `json:"token_exchange_plan_only"`
	UserDenialHandled             bool   `json:"user_denial_handled"`
	LaunchFailureFallback         bool   `json:"launch_failure_fallback"`
	FallbackURLReported           bool   `json:"fallback_url_reported"`
	RawAuthorizationURLReported   bool   `json:"raw_authorization_url_reported"`
	RawStateReported              bool   `json:"raw_state_reported"`
	RawCodeReported               bool   `json:"raw_code_reported"`
}

func validateShortsActivationBrowser(value ShortsActivationBrowser) error {
	if value.Executable != "/usr/bin/open" || !value.FakeRunner || value.ActualBrowserLaunched ||
		value.CommandsValidated != 4 || !value.BoundBeforeAuthorization || !value.AuthorizationOriginPinned ||
		!value.RedirectBound || !value.SystemBrowserOnly || !value.ExternalNetworkPermitRequired ||
		!value.DefaultPermitDenied || !value.StateValidated || !value.WrongStateRejected ||
		!value.PKCEExchangePlanBound || !value.TokenExchangePlanOnly || !value.UserDenialHandled ||
		!value.LaunchFailureFallback || value.FallbackURLReported || value.RawAuthorizationURLReported ||
		value.RawStateReported || value.RawCodeReported {
		return errors.New("Shorts activation browser proof is invalid")
	}
	return nil
}
