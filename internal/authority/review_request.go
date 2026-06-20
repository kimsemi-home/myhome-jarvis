package authority

import "time"

func ReviewRequestPacketForRoot(root string) (ReviewRequestPacket, error) {
	plan, err := ReviewPlanForRoot(root)
	if err != nil {
		return ReviewRequestPacket{}, err
	}
	return ReviewRequestPacketFromPlan(plan), nil
}

func ReviewRequestPacketFromPlan(plan ReviewPlanStatus) ReviewRequestPacket {
	packet := ReviewRequestPacket{
		PolicyPath:                        plan.PolicyPath,
		RequestState:                      reviewRequestState(plan),
		PublicSafe:                        plan.PublicSafe,
		Redaction:                         "review-request-public-packet",
		SourceAction:                      "request_authority_review",
		NextHandling:                      reviewRequestNextHandling(plan),
		ApprovalGranted:                   false,
		ExternalWritesAllowed:             false,
		SelfApprovalAllowed:               false,
		ReviewRequestable:                 plan.ReviewRequestable,
		ReviewCapacityState:               plan.ReviewCapacityState,
		HighRiskBlockedDecisionCount:      plan.HighRiskBlockedDecisionCount,
		ReviewRequiredDecisionCount:       plan.ReviewRequiredDecisionCount,
		ReviewRequiredProfileCount:        plan.ReviewRequiredProfileCount,
		PublicRepoReviewProfileCount:      plan.PublicRepoReviewProfileCount,
		WorkflowReviewProfileCount:        plan.WorkflowReviewProfileCount,
		SelfApprovalBlockedProfileCount:   plan.SelfApprovalBlockedProfileCount,
		ExternalWritesAllowedProfileCount: plan.ExternalWritesAllowedProfileCount,
		RequiredReviewClasses:             normalizeList(plan.RequiredReviewClasses),
		IncludedEvidenceFields:            reviewRequestEvidenceFields(),
		CheckedAt:                         time.Now().UTC().Format(time.RFC3339),
	}
	packet.RequestID = reviewRequestID(packet)
	return packet
}
