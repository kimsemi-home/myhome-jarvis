package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateCreditPreview(value CreditImportPreview, imported CreditTemplateImport) error {
	if err := validateCreditPreviewContract(value); err != nil {
		return err
	}
	if value.SourceSHA256 != imported.FileSHA256 || value.ProfileSHA256 != imported.ProfileSHA256 ||
		value.FingerprintSetHash != imported.FingerprintSetHash || value.TemplateID != imported.TemplateID ||
		value.TemplateVersion != imported.TemplateVersion || value.TransactionCount != imported.Read ||
		value.DebitMinor != imported.NormalizedDebitMinor || value.CreditMinor != imported.NormalizedCreditMinor {
		return errors.New("Ledger import preview is not bound to the import")
	}
	return nil
}

func validateCreditPreviewContract(value CreditImportPreview) error {
	if value.SchemaVersion != "myhome.ledger-import-preview/v2" || value.ExecutionMode != "preview_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || value.SourceExtension != ".csv" ||
		value.CandidateCount != 1 || !hashPattern.MatchString(value.SourceSHA256) ||
		!hashPattern.MatchString(value.ProfileSHA256) || !hashPattern.MatchString(value.FingerprintSetHash) ||
		value.AccountKind != "credit_card" || value.BalanceConvention != "liability" ||
		value.TransactionCount < 1 || value.DebitMinor < 0 || value.CreditMinor < 0 ||
		value.NetCardSpendMinor != value.DebitMinor-value.CreditMinor || value.SuggestionCount < 0 ||
		!value.Expected.Provided || value.Expected.TransactionCount != value.TransactionCount ||
		value.Expected.DebitMinor != value.DebitMinor || value.Expected.CreditMinor != value.CreditMinor ||
		!value.ExpectedBalance.Provided || !value.Reconciled || !value.BalanceReconciled || !value.ReadyToImport ||
		value.RawRowsReported || !slices.Equal(value.Checks, requiredCreditPreviewChecks()) ||
		!expectedCreditPreview(value) || !validCreditBalance(value) {
		return errors.New("Ledger import preview contract is invalid")
	}
	copy := value
	copy.PreviewHash = ""
	body, err := json.Marshal(copy)
	if err != nil || !hashPattern.MatchString(value.PreviewHash) || value.PreviewHash != digest(string(body)) {
		return errors.New("Ledger import preview hash is invalid")
	}
	return nil
}

func requiredCreditPreviewChecks() []string {
	return []string{
		"account-kind-balance-convention", "balance-reconciliation-explicit", "bounded-input",
		"candidate-unique", "profile-hash-bound", "raw-rows-not-reported", "reconciliation-explicit", "source-hash-bound",
	}
}
