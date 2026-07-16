package localfinancereadiness

import (
	"errors"
	"slices"
	"strings"
)

var expectedSchedules = map[string]Schedule{
	"ledger":    {Day: 2, Hour: 7, Minute: 0, Timezone: "Asia/Seoul"},
	"portfolio": {Day: 3, Hour: 7, Minute: 20, Timezone: "Asia/Seoul"},
	"revenue":   {Day: 5, Hour: 7, Minute: 40, Timezone: "Asia/Seoul"},
}

func validatePlan(value Plan, ref Ref) error {
	expected, ok := expectedSchedules[value.Component]
	args := value.Launchd.ProgramArguments
	if !ok || value.Component != ref.Component || value.SchemaVersion != PlanSchema || value.ExecutionMode != "plan_only" ||
		value.CredentialsRead || value.ExternalNetworkEnabled || value.ExternalWritesEnabled || value.InstallAllowed ||
		value.Schedule != expected || !strings.HasPrefix(value.PublicConfigSource, "config/") ||
		strings.Contains(value.PublicConfigSource, "private") || len(value.KeychainHandles) == 0 || len(value.OfficialHosts) == 0 ||
		len(args) != 3 || args[1] != "collect" || args[2] != "monthly" || !slices.Contains(value.Checks, "private-config-ignored") {
		return errors.New("plan safety boundary is invalid")
	}
	for _, scope := range value.OAuthScopes {
		if !strings.HasSuffix(scope, ".readonly") {
			return errors.New("write-capable OAuth scope is present")
		}
	}
	if value.Component == "ledger" && !validLedgerOAuthBoundary(value) {
		return errors.New("Ledger OAuth readiness boundary is incomplete")
	}
	if !hashPattern.MatchString(value.TemplateHash) || value.PlanHash != ref.PlanHash || value.PlanHash != planHash(value) {
		return errors.New("plan hash is invalid")
	}
	return nil
}

func validLedgerOAuthBoundary(value Plan) bool {
	if !slices.Equal(value.OAuthScopes, []string{"https://www.googleapis.com/auth/gmail.readonly"}) {
		return false
	}
	for _, host := range []string{"accounts.google.com", "gmail.googleapis.com", "oauth2.googleapis.com"} {
		if !slices.Contains(value.OfficialHosts, host) {
			return false
		}
	}
	for _, check := range []string{"oauth-exact-loopback-callback", "oauth-redirects-disabled", "oauth-response-bounded", "oauth-token-origin-pinned"} {
		if !slices.Contains(value.Checks, check) {
			return false
		}
	}
	return true
}
