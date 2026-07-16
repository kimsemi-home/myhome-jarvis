package localfinanceevidence

import (
	"errors"
	"slices"
)

func validateCreditApplyReport(value CreditBatchApplyReport, plan CreditBatchApplyPlan, approval CreditBatchApproval, replay bool) error {
	inserted, duplicates, suggestions := 5, 0, 3
	if replay {
		inserted, duplicates, suggestions = 0, 5, 0
	}
	checks := []string{"approval-exact-match", "atomic-multi-file-transaction", "current-bytes-renormalized",
		"filename-hash-only", "manifest-reloaded", "preview-fields-recomputed", "raw-rows-not-reported"}
	if value.SchemaVersion != "myhome.ledger-import-batch-apply-report/v1" || value.ExecutionMode != "approved_local_batch_apply" ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || !value.LocalDatabaseWrites ||
		value.PlanHash != plan.PlanHash || value.ApprovalHash != approval.ApprovalHash ||
		value.ManifestSHA256 != plan.ManifestSHA256 || value.PreviewSetHash != plan.PreviewSetHash || value.BatchHash != plan.BatchHash ||
		value.StatementCount != 2 || value.TransactionCount != 5 || value.RowsInserted != inserted || value.Duplicates != duplicates ||
		value.SuggestionsRecorded != suggestions || !value.AllOrNothing || value.RawFileNamesReported || value.RawRowsReported ||
		!slices.Equal(value.Checks, checks) || validateCreditApplyStatements(value.Statements, replay) != nil ||
		value.ApplySetHash != creditApplySetHash(value.Statements) || value.ReportHash != creditApplyReportHash(value) {
		return errors.New("Ledger batch apply report is invalid")
	}
	return nil
}

func validateCreditApplyStatements(values []CreditBatchApplyStatement, replay bool) error {
	names := []string{"d8852d20dfb991da62854549f8ebe761d1cba30a5fc62dc658b18684f3432079", "e817e6bee3b048c5ce3abde65a8a1576681ebd6c86e306382145274f15874df6"}
	previews := []string{"8c3488195ef60cd72498d8c569b61870398e25ee0b4513206488fd51b49b2eef", "e5a89a5320372d3050c2219f359972f2d5840b07d0b63728a3d8ba53217e96f0"}
	reads, firstSuggestions := []int{2, 3}, []int{0, 3}
	if len(values) != 2 {
		return errors.New("Ledger batch apply statement count is invalid")
	}
	for index, value := range values {
		inserted, duplicates, suggestions := reads[index], 0, firstSuggestions[index]
		if replay {
			inserted, duplicates, suggestions = 0, reads[index], 0
		}
		if value.SourceNameSHA256 != names[index] || value.PreviewHash != previews[index] || value.RowsRead != reads[index] ||
			value.RowsInserted != inserted || value.Duplicates != duplicates || value.Suggestions != suggestions {
			return errors.New("Ledger batch apply statement result is invalid")
		}
	}
	return nil
}
