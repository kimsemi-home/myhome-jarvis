package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validatePortfolioReport(value PortfolioReport, month string, ref ProofRef) error {
	if value.SchemaVersion != portfolioProofSchema || value.ExecutionMode != "loopback_emulation" ||
		!value.LoopbackOnly || value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || value.FinancialActions ||
		value.Month != month || value.ReportHash != ref.ReportHash || value.KIS.Method != "GET" ||
		value.KIS.Path != "/uapi/domestic-stock/v1/trading/inquire-balance" || value.KIS.Transaction != "TTTC8434R" ||
		value.KIS.Retries != 1 || value.KIS.HoldingCount < 1 || !requiredPortfolioChecks(value.Checks) {
		return errors.New("Portfolio execution proof boundary is invalid")
	}
	token := value.Token
	if token.SchemaVersion != "myhome.portfolio-kis-token-rehearsal/v1" || token.ExecutionMode != "loopback_emulation" ||
		!token.LoopbackOnly || token.CredentialsRead || token.ExternalNetwork || token.ExternalWrites ||
		!token.ClientCredentialsExchange || !token.TokenContractValidated || !token.OfficialTokenEndpointAllowed ||
		!token.OfficialOriginPinned || !token.OrderPathRejected || !token.RedirectRejected || !token.OversizedResponseRejected {
		return errors.New("Portfolio KIS token proof boundary is invalid")
	}
	if value.Store.FirstSnapshotCount != 1 || value.Store.ReplaySnapshotCount != 1 ||
		!value.Ledger.FirstInserted || value.Ledger.ReplayInserted || !value.Ledger.FingerprintMatched || !value.Ledger.AggregateOnly {
		return errors.New("Portfolio execution proof idempotency is invalid")
	}
	metrics := value.Emulator
	if metrics.TokenRequests != 2 || metrics.BalanceRequests != 2 || metrics.LedgerRequests != 2 ||
		metrics.InjectedFailures != 1 || metrics.OrderRequests != 0 || metrics.ForbiddenRequests != 0 ||
		metrics.RedirectRequests != 1 || metrics.OversizedRequests != 1 {
		return errors.New("Portfolio execution proof metrics are invalid")
	}
	monthly := value.Monthly
	if !monthly.Reconciled || monthly.CashMinor+monthly.SecuritiesValueMinor != monthly.TotalValueMinor ||
		monthly.HoldingValueMinor != monthly.SecuritiesValueMinor {
		return errors.New("Portfolio execution proof reconciliation is invalid")
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Portfolio execution report hash is invalid")
	}
	return nil
}

func requiredPortfolioChecks(checks []string) bool {
	required := []string{"aggregate-only-ledger-boundary", "bounded-retry-recovered", "exact-ipv4-loopback-origin", "financial-actions-disabled", "ledger-replay-idempotent", "order-requests-zero", "kis-order-path-rejected", "kis-token-contract-validated", "kis-token-origin-pinned", "kis-token-redirect-rejected", "kis-token-response-bounded", "readonly-balance-contract", "snapshot-reconciliation-preserved", "sqlite-replay-idempotent"}
	if len(checks) != len(required) {
		return false
	}
	for _, check := range required {
		if !slices.Contains(checks, check) {
			return false
		}
	}
	return true
}
