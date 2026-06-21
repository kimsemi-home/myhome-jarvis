package authority

import "time"

func approvalScopeSummary(
	record ApprovalDecisionRecord,
	now time.Time,
) ApprovalScopeSummary {
	leaseState := "expired"
	if approvalRecordActive(record, now) {
		leaseState = "active"
	}
	return ApprovalScopeSummary{
		Scope:              record.Scope,
		Target:             record.Target,
		LeaseState:         leaseState,
		ExpiresAt:          record.ExpiresAt,
		GrantFlags:         record.GrantFlags,
		CanUnlockScopeOnly: approvalFlagsMatchScope(record),
	}
}

func approvalRecordActive(record ApprovalDecisionRecord, now time.Time) bool {
	expiresAt, err := time.Parse(time.RFC3339, record.ExpiresAt)
	if err != nil {
		return false
	}
	return record.PublicSafe && approvalRecordPublicSafe(record) &&
		approvalFlagsMatchScope(record) && expiresAt.After(now)
}

func approvalRecordPublicSafe(record ApprovalDecisionRecord) bool {
	return record.PublicSafe &&
		record.DecisionPacketRef == "external_evidence_repo_split_decision" &&
		record.DecisionPacketContext == "ExternalEvidenceRepoSplitDecisionPacket" &&
		record.LeaseState == "active" && record.Scope != "" &&
		record.Target != "" && approvalFlagsMatchScope(record) &&
		!unsafeApprovalText(record.Target) &&
		!unsafeApprovalText(record.ReviewerBoundary) &&
		!record.ReviewerIsRequester &&
		!record.GrantFlags.SelfApprovalAllowed
}
