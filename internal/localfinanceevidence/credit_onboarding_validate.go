package localfinanceevidence

import "errors"

func validateCreditOnboarding(value CreditTemplateOnboarding, first, second CreditTemplateImport) error {
	if err := validateCreditPreview(value.Version1Preview, first); err != nil {
		return err
	}
	if err := validateCreditPreview(value.Version2Preview, second); err != nil {
		return err
	}
	if err := validateCreditBatch(value.BatchPreview, value.Version1Preview); err != nil {
		return err
	}
	if value.Version1Preview.PreviewHash == value.Version2Preview.PreviewHash ||
		value.Version1Preview.SourceSHA256 == value.Version2Preview.SourceSHA256 ||
		value.Version1Preview.ProfileSHA256 == value.Version2Preview.ProfileSHA256 ||
		value.Version1Preview.FingerprintSetHash != value.Version2Preview.FingerprintSetHash ||
		!value.AmbiguousProfileRejected || !value.UnsupportedStatementRejected ||
		!value.MismatchedExpectedTotalsBlocked || !value.MismatchedExpectedBalanceBlocked ||
		!value.BalanceConventionValidated || !value.DuplicateBatchContentRejected ||
		!value.DuplicateBatchIdentityRejected || !value.BatchPathTraversalRejected ||
		!value.BatchRootEscapeRejected || !value.SourceMutationRebound {
		return errors.New("Ledger profile onboarding attacks are invalid")
	}
	return nil
}
