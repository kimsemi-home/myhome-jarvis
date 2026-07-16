package localfinanceevidence

import (
	"errors"
	"slices"
)

func validateCreditBatchApplyRehearsal(value CreditBatchApplyRehearsal, ref ProofRef) error {
	if validateCreditApplyPlan(value.Plan) != nil || validateCreditApproval(value.Approval, value.Plan) != nil ||
		validateCreditApplyReport(value.FirstApply, value.Plan, value.Approval, false) != nil ||
		validateCreditApplyReport(value.Replay, value.Plan, value.Approval, true) != nil {
		return errors.New("Ledger batch apply nested proof is invalid")
	}
	state := value.RollbackState
	checks := []string{"approval-denial-before-write", "approval-staleness-before-write", "atomic-mid-batch-rollback",
		"ephemeral-database-only", "first-apply-five-inserted", "replay-five-deduplicated"}
	if value.SchemaVersion != creditBatchApplyProofSchema || value.ExecutionMode != "fixture_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || !value.EphemeralLocalDatabaseWrites ||
		value.PersistentDatabaseWrites || !value.StaleApprovalRejected || !value.DeniedApprovalRejected ||
		!value.MidBatchFailureRejected || state.Accounts != 0 || state.Transactions != 0 || state.ImportRuns != 0 ||
		state.Templates != 0 || state.CategorySuggestions != 0 || !slices.Equal(value.Checks, checks) ||
		value.ReportHash != ref.ReportHash || value.ReportHash != creditApplyRehearsalHash(value) {
		return errors.New("Ledger batch apply rehearsal is invalid")
	}
	return nil
}
