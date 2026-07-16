package localfinanceevidence

import (
	"errors"
	"slices"
)

func validateCreditTemplateParts(value CreditTemplateReport) error {
	history, reconciliation, guards := value.History, value.Reconciliation, value.Guards
	if history.EntryCount != 6 || history.Version1Entries != 3 || history.Version2Entries != 3 ||
		!hashPattern.MatchString(history.HistoryHash) || !history.ExistingCategoryPreserved || history.RawRowsReported {
		return errors.New("Ledger category suggestion history is invalid")
	}
	if reconciliation.TransactionCount != 3 || reconciliation.CardSpendMinor != 20900 ||
		reconciliation.CardRefundMinor != 2200 || reconciliation.NetCardSpendMinor != 18700 ||
		!reconciliation.BothVersionsMatch || !reconciliation.MonthlyMatch {
		return errors.New("Ledger template reconciliation is invalid")
	}
	if !guards.StableIdentityConflictRejected || !guards.MissingSourceIDRejected ||
		!guards.TemplateVersionMutationRejected || guards.PartialWritesAfterRejection {
		return errors.New("Ledger template guard attacks are invalid")
	}
	if !slices.Equal(value.Checks, requiredCreditTemplateChecks()) {
		return errors.New("Ledger template checks are invalid")
	}
	return nil
}

func requiredCreditTemplateChecks() []string {
	return []string{
		"ambiguous-profile-denied", "append-only-category-history", "classification-suggestion-only", "cross-version-source-id-dedup",
		"deterministic-fingerprint-set", "immutable-template-version-registry", "profile-hashes-distinct",
		"profile-auto-detected", "preview-hash-sealed", "raw-rows-not-reported", "reconciliation-gates-import",
		"source-hashes-bound", "source-mutation-rebound", "stable-identity-conflict-denied",
		"structured-export-reconciled", "template-versions-explicit", "unsupported-statement-denied",
	}
}
