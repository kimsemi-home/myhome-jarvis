package localfinancereadiness

import "slices"

func validPortfolioKISBoundary(value Plan) bool {
	if len(value.OAuthScopes) != 0 {
		return false
	}
	for _, host := range []string{"data-dbg.krx.co.kr", "openapi.koreainvestment.com:9443", "opendart.fss.or.kr"} {
		if !slices.Contains(value.OfficialHosts, host) {
			return false
		}
	}
	for _, check := range []string{"kis-exact-official-origin", "kis-order-path-rejected", "kis-redirect-disabled", "kis-token-contract-validated", "kis-token-response-bounded"} {
		if !slices.Contains(value.Checks, check) {
			return false
		}
	}
	return true
}
