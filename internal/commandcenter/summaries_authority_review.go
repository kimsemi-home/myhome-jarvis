package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func summarizeAuthorityReview(status authority.ReviewPlanStatus) AuthorityReviewSummary {
	packet := authority.ReviewRequestPacketFromPlan(status)
	evidence := authority.ReviewRequestEvidenceFromPacket(packet)
	return AuthorityReviewSummary{
		RequestID:                       packet.RequestID,
		RequestState:                    packet.RequestState,
		EvidenceRef:                     evidence.EvidenceRef,
		EvidenceState:                   evidence.EvidenceState,
		EvidenceReady:                   evidence.EvidenceReady,
		PublicSafe:                      status.PublicSafe,
		ReviewRequestable:               status.ReviewRequestable,
		ReviewCapacityState:             status.ReviewCapacityState,
		NextSafeAction:                  status.NextSafeAction,
		HighRiskBlockedDecisionCount:    status.HighRiskBlockedDecisionCount,
		RequiredReviewClassCount:        len(status.RequiredReviewClasses),
		PublicRepoReviewProfileCount:    status.PublicRepoReviewProfileCount,
		WorkflowReviewProfileCount:      status.WorkflowReviewProfileCount,
		SelfApprovalBlockedProfileCount: status.SelfApprovalBlockedProfileCount,
	}
}
