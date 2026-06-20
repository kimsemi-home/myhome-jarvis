package authority

import "time"

func ReviewQueueStatusForRoot(root string) (ReviewQueueStatus, error) {
	packet, err := ReviewRequestPacketForRoot(root)
	if err != nil {
		return ReviewQueueStatus{}, err
	}
	evidence := ReviewRequestEvidenceFromPacket(packet)
	return ReviewQueueStatusFromPacket(packet, evidence), nil
}

func ReviewQueueStatusFromPacket(
	packet ReviewRequestPacket,
	evidence ReviewRequestEvidenceStatus,
) ReviewQueueStatus {
	ready := reviewQueueReady(packet, evidence)
	return ReviewQueueStatus{
		PolicyPath:              packet.PolicyPath,
		RequestID:               packet.RequestID,
		EvidenceRef:             evidence.EvidenceRef,
		QueueItemRef:            "authority_review_queue:" + packet.RequestID,
		RequestState:            packet.RequestState,
		EvidenceState:           evidence.EvidenceState,
		QueueState:              reviewQueueState(ready),
		QueueReady:              ready,
		PublicSafe:              packet.PublicSafe && evidence.PublicSafe,
		Redaction:               "review-queue-public-status",
		PendingReviewClassCount: len(packet.RequiredReviewClasses),
		ApprovalState:           "not_approved",
		ApprovalGranted:         false,
		ExternalWritesAllowed:   false,
		SelfApprovalAllowed:     false,
		NextSafeAction:          "await_human_authority_review",
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}
