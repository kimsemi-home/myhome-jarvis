package authority

import (
	"fmt"
	"strings"
	"time"
)

func normalizeReviewRecordRequest(
	policy Policy,
	request ReviewRecordRequest,
	packet ReviewRequestPacket,
	evidence ReviewRequestEvidenceStatus,
	queue ReviewQueueStatus,
	now time.Time,
) (ReviewRecord, error) {
	if !privateJSONL(policy.PrivateReviewRequestLedger) {
		return ReviewRecord{}, fmt.Errorf("authority review ledger must stay private jsonl")
	}
	recordedAt, err := normalizeReviewRecordTime(request.At, now)
	if err != nil {
		return ReviewRecord{}, err
	}
	if err := validateExplicitNonApproval(request); err != nil {
		return ReviewRecord{}, err
	}
	requestID := strings.TrimSpace(request.RequestID)
	evidenceRef := strings.TrimSpace(request.EvidenceRef)
	queueItemRef := strings.TrimSpace(request.QueueItemRef)
	queueState := normalizeToken(request.QueueState)
	classes := normalizeList(request.RequiredReviewClasses)
	linearIssueRef, err := normalizeLinearIssueRef(request.LinearIssueRef)
	if err != nil {
		return ReviewRecord{}, err
	}
	if err := validateReviewRecordRefs(requestID, evidenceRef, queueItemRef, packet, evidence, queue); err != nil {
		return ReviewRecord{}, err
	}
	if queueState != queue.QueueState || queueState != "pending_human_review" {
		return ReviewRecord{}, fmt.Errorf("authority review queue state must be pending_human_review")
	}
	if !sameStringSet(classes, normalizeList(packet.RequiredReviewClasses)) {
		return ReviewRecord{}, fmt.Errorf("authority review classes do not match current packet")
	}
	if err := validateReviewRecordPublicBoundary(packet, evidence, queue); err != nil {
		return ReviewRecord{}, err
	}
	return ReviewRecord{
		At:                         recordedAt,
		RequestID:                  requestID,
		RequestState:               packet.RequestState,
		EvidenceRef:                evidence.EvidenceRef,
		EvidenceState:              evidence.EvidenceState,
		QueueItemRef:               queue.QueueItemRef,
		QueueState:                 queue.QueueState,
		LinearIssueRef:             linearIssueRef,
		SourceAction:               packet.SourceAction,
		NextSafeAction:             queue.NextSafeAction,
		ReviewCapacityState:        packet.ReviewCapacityState,
		RequiredReviewClasses:      classes,
		RequiredReviewClassCount:   len(classes),
		HighRiskBlockedDecisionCnt: packet.HighRiskBlockedDecisionCount,
		ReviewRequiredDecisionCnt:  packet.ReviewRequiredDecisionCount,
		ReviewRequiredProfileCnt:   packet.ReviewRequiredProfileCount,
		ApprovalState:              queue.ApprovalState,
		ApprovalGranted:            false,
		ExternalWritesAllowed:      false,
		SelfApprovalAllowed:        false,
		PublicSafe:                 true,
	}, nil
}
