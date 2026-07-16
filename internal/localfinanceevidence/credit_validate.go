package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateCreditReport(value CreditReport, month string, ref ProofRef) error {
	if value.SchemaVersion != creditProofSchema || value.ExecutionMode != "loopback_emulation" ||
		!value.LoopbackOnly || value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites ||
		value.Month != month || !hashPattern.MatchString(value.AttachmentContentHash) ||
		value.ReportHash != ref.ReportHash || !requiredCreditChecks(value.Checks) {
		return errors.New("Ledger credit execution proof boundary is invalid")
	}
	if value.OAuth.SchemaVersion != "myhome.ledger-oauth-token-rehearsal/v1" || value.OAuth.ExecutionMode != "loopback_emulation" ||
		!value.OAuth.LoopbackOnly || value.OAuth.CredentialsRead || value.OAuth.ExternalNetwork || value.OAuth.ExternalWrites ||
		!value.OAuth.AuthorizationCodeExchange || !value.OAuth.RefreshTokenExchange || !value.OAuth.OfficialOriginPinned ||
		!value.OAuth.RedirectRejected || !value.OAuth.OversizedResponseRejected || value.OAuth.TokenRequests != 2 ||
		value.OAuth.RedirectRequests != 1 || value.OAuth.OversizedRequests != 1 {
		return errors.New("Ledger OAuth token-boundary proof is invalid")
	}
	if value.FirstGmail.SchemaVersion != "myhome.ledger-gmail-sync/v1" ||
		value.FirstGmail.AttachmentsWritten != 1 || value.FirstGmail.ReceiptsWritten < 1 || value.FirstGmail.Retries < 1 ||
		value.FirstWatch.Files != 1 || value.FirstWatch.Inserted < 1 ||
		value.ReplayGmail.AttachmentsWritten != 0 || value.ReplayGmail.PreviouslyProcessed < 1 ||
		value.ReplayWatch.Files != 0 || value.ArchiveFallbackGmail.AttachmentsWritten != 0 ||
		value.ArchiveFallbackGmail.ReceiptsWritten < 1 || value.ArchiveFallbackWatch.Files != 0 {
		return errors.New("Ledger credit retry or idempotency proof is invalid")
	}
	monthly := value.Monthly
	if monthly.TransactionCount < 1 || !monthly.Reconciled || monthly.CardSpendMinor-monthly.CardRefundMinor != monthly.NetCardSpendMinor || value.Emulator.InjectedFailures < 1 {
		return errors.New("Ledger credit monthly reconciliation proof is invalid")
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Ledger credit execution report hash is invalid")
	}
	return nil
}

func requiredCreditChecks(checks []string) bool {
	for _, required := range []string{"allowed-sender-enforced", "append-only-receipts-validated", "archive-dedup-enforced", "bounded-retry-recovered", "credit-refunds-netted", "idempotent-replay", "loopback-origin-enforced", "oauth-official-origin-pinned", "oauth-redirect-rejected", "oauth-response-bounded", "oauth-token-contract-validated"} {
		if !slices.Contains(checks, required) {
			return false
		}
	}
	return true
}
