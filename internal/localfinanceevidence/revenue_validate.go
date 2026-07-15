package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateRevenueReport(value RevenueReport, month string, ref ProofRef) error {
	if value.SchemaVersion != revenueProofSchema || value.Component != "revenue" || value.ExecutionMode != "ipv4_loopback" ||
		value.Month != month || value.Currency != "KRW" || value.RevenueStatus != "estimated" || !value.LoopbackOnly ||
		value.CredentialsRead || value.OAuthPerformed || value.ExternalNetworkEnabled || value.ExternalWritesEnabled ||
		value.ChannelWritesEnabled || value.ServiceInstalled || !value.ChannelIdentityBound || value.ReportHash != ref.ReportHash {
		return errors.New("Revenue execution proof boundary is invalid")
	}
	if value.DailyRows != 2 || value.VideoRows != 2 || value.PersistedDailyRows != 2 || value.PersistedVideoRows != 2 ||
		value.CostRows != 2 || value.CostReplayDuplicates != 2 || value.GrossMinor != 8300 || value.CostMinor != 2000 ||
		value.NetMinor != 6300 || value.GrossMinor-value.CostMinor != value.NetMinor || value.Views != 6800 ||
		value.WatchMinutes != 2900.5 || value.MonetizedPlaybacks != 1440 {
		return errors.New("Revenue execution proof reconciliation is invalid")
	}
	if value.ChannelRequests != 1 || value.DailyReportRequests != 2 || value.VideoReportRequests != 1 ||
		value.LedgerRequests != 2 || value.InjectedFailures != 1 || value.ForbiddenRequests != 0 ||
		value.ChannelWriteRequests != 0 || value.LedgerStoredEvents != 1 || !value.FirstLedgerInserted ||
		value.ReplayLedgerInserted || !hashPattern.MatchString(value.LedgerFingerprint) || !hashPattern.MatchString(value.EvidenceHash) {
		return errors.New("Revenue execution proof request accounting is invalid")
	}
	if !requiredRevenueChecks(value.Checks) {
		return errors.New("Revenue execution proof checks are invalid")
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Revenue execution report hash is invalid")
	}
	return nil
}

func requiredRevenueChecks(checks []string) bool {
	required := []string{"aggregate-only-ledger-event", "bound-channel-identity", "bounded-retry-after-503", "cost-attribution-idempotent", "exact-ipv4-loopback", "ledger-fingerprint-replay", "monetary-readonly-query", "monthly-replacement-idempotent", "proxy-and-redirect-disabled", "youtube-channel-write-absent"}
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
