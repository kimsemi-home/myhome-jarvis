package localfinancereadiness

import "slices"

func validRevenueOAuthBoundary(value Plan) bool {
	if !slices.Equal(value.OAuthScopes, []string{"https://www.googleapis.com/auth/youtube.readonly", "https://www.googleapis.com/auth/yt-analytics-monetary.readonly"}) {
		return false
	}
	for _, host := range []string{"accounts.google.com", "oauth2.googleapis.com", "www.googleapis.com", "youtubeanalytics.googleapis.com"} {
		if !slices.Contains(value.OfficialHosts, host) {
			return false
		}
	}
	for _, check := range []string{"oauth-callback-exact", "oauth-official-origin-pinned", "oauth-redirect-disabled", "oauth-response-bounded"} {
		if !slices.Contains(value.Checks, check) {
			return false
		}
	}
	return true
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
