package authority

import (
	"time"
)

func ApprovalDecisionStatusForRoot(root string) (ApprovalDecisionStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ApprovalDecisionStatus{}, err
	}
	return approvalDecisionStatusForRoot(root, policy, time.Now().UTC())
}

func approvalDecisionStatusForRoot(
	root string,
	policy Policy,
	now time.Time,
) (ApprovalDecisionStatus, error) {
	records, invalid, missing, err := readApprovalDecisionLedger(root, policy)
	if err != nil {
		return ApprovalDecisionStatus{}, err
	}
	status := ApprovalDecisionStatus{PolicyPath: PolicyRelativePath,
		PublicSafe: true, InvalidRecordCount: invalid,
		NextSafeAction: "record_human_approval_decision",
		CheckedAt:      now.Format(time.RFC3339)}
	if missing {
		status.LedgerState = "missing"
		return status, nil
	}
	status.LedgerState = "present"
	applyApprovalRecords(&status, records, now)
	if status.ActiveApprovalCount > 0 {
		status.NextSafeAction = "use_matching_scoped_approval_only"
	}
	return status, nil
}
