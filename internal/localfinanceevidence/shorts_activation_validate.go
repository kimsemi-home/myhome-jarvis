package localfinanceevidence

import (
	"encoding/json"
	"errors"
)

func validateShortsActivationReport(value ShortsActivationReport, ref ProofRef) error {
	if value.SchemaVersion != shortsActivationProofSchema || value.ExecutionMode != "loopback_fake_runner" ||
		!value.LoopbackOnly || value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites ||
		!value.RuntimeEntrypointsInactive || value.ReportHash != ref.ReportHash {
		return errors.New("Shorts activation proof boundary is invalid")
	}
	if err := validateShortsActivationCallback(value.Callback); err != nil {
		return err
	}
	if err := validateShortsActivationBrowser(value.Browser); err != nil {
		return err
	}
	if err := validateShortsActivationKeychain(value.Keychain); err != nil {
		return err
	}
	if !slicesEqual(value.Checks, requiredShortsActivationChecks()) {
		return errors.New("Shorts activation proof checks are invalid")
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Shorts activation report hash is invalid")
	}
	return nil
}

func requiredShortsActivationChecks() []string {
	return []string{
		"authorization-created-after-bind", "browser-command-exact", "browser-fake-runner-only",
		"browser-launch-fallback-bounded", "browser-user-denial-no-exchange",
		"callback-cancellation-closes-listener", "callback-exact-host-path-query",
		"callback-invalid-attempts-non-consuming", "callback-random-ipv4-port", "callback-single-consumption",
		"credentials-not-read", "external-network-disabled", "external-writes-disabled",
		"keychain-actual-execution-disabled", "keychain-argv-secret-free", "keychain-default-deny",
		"keychain-fake-runner-only", "keychain-permit-expiry-enforced", "keychain-readiness-plan-bound",
		"keychain-reference-scope-enforced", "oauth-pkce-exchange-plan-bound", "oauth-state-mismatch-rejected",
		"raw-authorization-material-not-reported", "runtime-entrypoints-inactive", "system-browser-not-embedded",
		"token-exchange-plan-only",
	}
}
