package localfinanceevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRehashedBatchApplyHierarchyForDifferentBatchFails(t *testing.T) {
	report := loadCreditApplyProof(t)
	report.Plan.BatchHash = strings.Repeat("a", 64)
	report.Plan.ApprovalChallengeSHA256 = creditApprovalChallenge(report.Plan)
	report.Plan.PlanHash = creditApplyPlanHash(report.Plan)
	report.Approval.PlanHash = report.Plan.PlanHash
	report.Approval.ApprovalChallengeSHA256 = report.Plan.ApprovalChallengeSHA256
	report.Approval.BatchHash = report.Plan.BatchHash
	report.Approval.ApprovalHash = creditApprovalHash(report.Approval)
	for _, apply := range []*CreditBatchApplyReport{&report.FirstApply, &report.Replay} {
		apply.PlanHash = report.Plan.PlanHash
		apply.ApprovalHash = report.Approval.ApprovalHash
		apply.BatchHash = report.Plan.BatchHash
		apply.ReportHash = creditApplyReportHash(*apply)
	}
	report.ReportHash = creditApplyRehearsalHash(report)
	if validateCreditBatchApplyRehearsal(report, creditApplyRef(report)) == nil {
		t.Fatal("accepted fully rehashed approval hierarchy for an unapproved batch")
	}
}

func TestRehashedBatchApplyPartialRollbackFails(t *testing.T) {
	report := loadCreditApplyProof(t)
	report.RollbackState.Transactions = 1
	report.ReportHash = creditApplyRehearsalHash(report)
	if validateCreditBatchApplyRehearsal(report, creditApplyRef(report)) == nil {
		t.Fatal("accepted rehashed batch proof with a partial transaction write")
	}
}

func loadCreditApplyProof(t *testing.T) CreditBatchApplyRehearsal {
	t.Helper()
	body, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "local_finance", "proofs", "ledger-credit-batch-apply.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report CreditBatchApplyRehearsal
	if err := json.Unmarshal(body, &report); err != nil {
		t.Fatal(err)
	}
	return report
}

func creditApplyRef(report CreditBatchApplyRehearsal) ProofRef {
	return ProofRef{Component: "ledger-batch-apply", Capability: "credit-batch-apply-rehearsal",
		ProofSchema: creditBatchApplyProofSchema, ReportHash: report.ReportHash}
}
