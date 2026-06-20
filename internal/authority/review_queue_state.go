package authority

func reviewQueueReady(
	packet ReviewRequestPacket,
	evidence ReviewRequestEvidenceStatus,
) bool {
	return evidence.EvidenceReady &&
		packet.RequestID != "" &&
		packet.RequestState == "ready" &&
		len(packet.RequiredReviewClasses) > 0 &&
		!packet.ApprovalGranted &&
		!packet.ExternalWritesAllowed &&
		!packet.SelfApprovalAllowed
}

func reviewQueueState(ready bool) string {
	if ready {
		return "pending_human_review"
	}
	return "blocked"
}
