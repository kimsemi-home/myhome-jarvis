package localfinanceevidence

import (
	"errors"
	"slices"
)

const creditBatchApplyProofSchema = "myhome.ledger-credit-batch-apply-rehearsal/v1"

const approvedManifestHash = "f378379d8cf2aadf06f3a9e984796d634e0e27760b71372d3a8e96afced90efa"
const approvedPreviewSetHash = "708cc7e3903d230ab6205762bf5103de193e490b6d22ee4f8bf1540f8955deac"
const approvedBatchHash = "d555bdd7ddcedb83901e9953c71f363d917d5279c46f583d0194fa1a624e0050"

func validateCreditApplyPlan(value CreditBatchApplyPlan) error {
	checks := []string{"all-statements-ready", "approval-challenge-hash-bound", "batch-hash-bound", "manifest-hash-bound", "preview-set-hash-bound"}
	if value.SchemaVersion != "myhome.ledger-import-batch-apply-plan/v1" || value.ExecutionMode != "approval_plan_only" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || value.ManifestSHA256 != approvedManifestHash ||
		value.PreviewSetHash != approvedPreviewSetHash || value.BatchHash != approvedBatchHash || value.StatementCount != 2 ||
		value.ApprovalChallengeSHA256 != creditApprovalChallenge(value) || !slices.Equal(value.Checks, checks) ||
		value.PlanHash != creditApplyPlanHash(value) {
		return errors.New("Ledger batch apply plan is invalid")
	}
	return nil
}

func validateCreditApproval(value CreditBatchApproval, plan CreditBatchApplyPlan) error {
	if value.SchemaVersion != "myhome.ledger-import-batch-approval/v1" || value.Decision != "approve" ||
		value.PlanHash != plan.PlanHash || value.ApprovalChallengeSHA256 != plan.ApprovalChallengeSHA256 ||
		value.ManifestSHA256 != plan.ManifestSHA256 || value.PreviewSetHash != plan.PreviewSetHash ||
		value.BatchHash != plan.BatchHash || value.ApprovalHash != creditApprovalHash(value) {
		return errors.New("Ledger batch approval is invalid")
	}
	return nil
}
