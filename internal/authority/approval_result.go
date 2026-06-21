package authority

func resultForApprovalDecision(
	record ApprovalDecisionRecord,
) ApprovalDecisionResult {
	return ApprovalDecisionResult{
		DecisionPacketRef:  record.DecisionPacketRef,
		Scope:              record.Scope,
		Target:             record.Target,
		LeaseState:         record.LeaseState,
		ExpiresAt:          record.ExpiresAt,
		GrantFlags:         record.GrantFlags,
		LedgerState:        "recorded_private",
		RecordedAt:         record.At,
		PublicSafe:         record.PublicSafe,
		CanUnlockScopeOnly: approvalFlagsMatchScope(record),
	}
}
