package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateCreditOnboarding(value CreditTemplateOnboarding, first, second CreditTemplateImport) error {
	if err := validateCreditPreview(value.Version1Preview, first); err != nil {
		return err
	}
	if err := validateCreditPreview(value.Version2Preview, second); err != nil {
		return err
	}
	if value.Version1Preview.PreviewHash == value.Version2Preview.PreviewHash ||
		value.Version1Preview.SourceSHA256 == value.Version2Preview.SourceSHA256 ||
		value.Version1Preview.ProfileSHA256 == value.Version2Preview.ProfileSHA256 ||
		value.Version1Preview.FingerprintSetHash != value.Version2Preview.FingerprintSetHash ||
		!value.AmbiguousProfileRejected || !value.UnsupportedStatementRejected ||
		!value.MismatchedExpectedTotalsBlocked || !value.SourceMutationRebound {
		return errors.New("Ledger profile onboarding attacks are invalid")
	}
	return nil
}

func validateCreditPreview(value CreditImportPreview, imported CreditTemplateImport) error {
	if value.SchemaVersion != "myhome.ledger-import-preview/v1" || value.ExecutionMode != "preview_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || value.SourceExtension != ".csv" ||
		value.CandidateCount != 1 || value.TemplateID != "synthetic-card-export" || value.IssuerKey != "synthetic-card" ||
		value.TransactionCount != 3 || value.DebitMinor != 20900 || value.CreditMinor != 2200 ||
		value.NetCardSpendMinor != 18700 || value.SuggestionCount != 3 || !value.Expected.Provided ||
		value.Expected.TransactionCount != value.TransactionCount || value.Expected.DebitMinor != value.DebitMinor ||
		value.Expected.CreditMinor != value.CreditMinor || !value.Reconciled || !value.ReadyToImport || value.RawRowsReported ||
		!slices.Equal(value.Checks, requiredCreditPreviewChecks()) {
		return errors.New("Ledger import preview contract is invalid")
	}
	if value.SourceSHA256 != imported.FileSHA256 || value.ProfileSHA256 != imported.ProfileSHA256 ||
		value.FingerprintSetHash != imported.FingerprintSetHash || value.TemplateID != imported.TemplateID ||
		value.TemplateVersion != imported.TemplateVersion || value.TransactionCount != imported.Read ||
		value.DebitMinor != imported.NormalizedDebitMinor || value.CreditMinor != imported.NormalizedCreditMinor {
		return errors.New("Ledger import preview is not bound to the import")
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
		"bounded-input", "candidate-unique", "profile-hash-bound", "raw-rows-not-reported",
		"reconciliation-explicit", "source-hash-bound",
	}
}
