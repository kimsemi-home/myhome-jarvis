package authority

import "time"

func ReviewPlanForRoot(root string) (ReviewPlanStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ReviewPlanStatus{}, err
	}
	status, err := StatusForRoot(root)
	if err != nil {
		return ReviewPlanStatus{}, err
	}
	plan := ReviewPlan(policy, status)
	packet := ReviewRequestPacketFromPlan(plan)
	ledger, err := ReviewRecordLedgerForRoot(root, policy, packet.RequestID)
	if err != nil {
		return ReviewPlanStatus{}, err
	}
	applyReviewRecordLedger(&plan, ledger)
	return plan, nil
}

func ReviewPlan(policy Policy, status Status) ReviewPlanStatus {
	plan := ReviewPlanStatus{
		PolicyPath:                      PolicyRelativePath,
		Outcome:                         status.Outcome,
		ActiveRule:                      status.ActiveRule,
		PublicSafe:                      reviewPlanPublicSafe(policy),
		Redaction:                       "review-classes-only",
		ReviewCapacityState:             status.HumanReviewCapacityState,
		BlockedDecisionCount:            status.BlockedDecisionCount,
		ReviewRequiredProfileCount:      status.ReviewRequiredProfileCount,
		SelfApprovalBlockedProfileCount: status.SelfApprovalBlockedProfileCount,
		NextSafeAction:                  "none",
		CheckedAt:                       time.Now().UTC().Format(time.RFC3339),
	}
	applyReviewDecisionCounts(policy, status, &plan)
	applyReviewProfilePlans(policy, &plan)
	plan.RequiredReviewClasses = normalizeList(plan.RequiredReviewClasses)
	plan.ReviewRequestable = reviewRequestable(policy, status, plan)
	if plan.ReviewRequestable {
		plan.NextSafeAction = "request_authority_review"
	} else if plan.BlockedDecisionCount > 0 {
		plan.NextSafeAction = "resolve_authority_debt"
	}
	return plan
}

func applyReviewRecordLedger(plan *ReviewPlanStatus, ledger ReviewRecordLedgerSummary) {
	plan.ReviewRequestRecorded = ledger.Recorded
	plan.ReviewRequestRecordCount = ledger.RecordCount
	plan.ReviewRequestInvalidRecordCount = ledger.InvalidRecordCount
	plan.ReviewRequestLedgerState = ledger.LedgerState
	plan.ReviewRequestApprovalState = ledger.ApprovalState
	plan.ReviewRequestLastRecordedAt = ledger.LastRecordedAt
	if plan.ReviewRequestable && ledger.Recorded {
		plan.NextSafeAction = "await_human_authority_review"
	}
	applyReviewRequestFreshness(plan, time.Now().UTC())
}

func reviewPlanPublicSafe(policy Policy) bool {
	return policy.PublicStatusRedacted &&
		!policy.SelfAuthorityAllowed &&
		!policy.ReasoningTierGrantsApproval &&
		policy.PublicRepoHighRiskBlocked
}
