package authority

import "time"

func ReviewRequestEvidenceForRoot(root string) (ReviewRequestEvidenceStatus, error) {
	plan, err := ReviewPlanForRoot(root)
	if err != nil {
		return ReviewRequestEvidenceStatus{}, err
	}
	return ReviewRequestEvidenceFromPlan(plan), nil
}

func ReviewRequestEvidenceFromPlan(plan ReviewPlanStatus) ReviewRequestEvidenceStatus {
	packet := ReviewRequestPacketFromPlan(plan)
	return reviewRequestEvidenceFromPacket(packet, plan.NextSafeAction)
}

func ReviewRequestEvidenceFromPacket(packet ReviewRequestPacket) ReviewRequestEvidenceStatus {
	return reviewRequestEvidenceFromPacket(packet, packet.SourceAction)
}

func reviewRequestEvidenceFromPacket(
	packet ReviewRequestPacket,
	nextSafeAction string,
) ReviewRequestEvidenceStatus {
	missing := missingReviewEvidenceFieldCount(packet)
	ready := reviewRequestEvidenceReady(packet, missing)
	if nextSafeAction == "" {
		nextSafeAction = packet.SourceAction
	}
	return ReviewRequestEvidenceStatus{
		PolicyPath:              packet.PolicyPath,
		RequestID:               packet.RequestID,
		RequestState:            packet.RequestState,
		EvidenceRef:             "authority_review_request:" + packet.RequestID,
		EvidenceState:           reviewRequestEvidenceState(ready),
		EvidenceReady:           ready,
		PublicSafe:              packet.PublicSafe,
		Redaction:               "request-evidence-public-status",
		ApprovalState:           "not_approved",
		ApprovalGranted:         false,
		ExternalWritesAllowed:   false,
		SelfApprovalAllowed:     false,
		RequiredEvidenceFields:  normalizeList(packet.IncludedEvidenceFields),
		MissingEvidenceFieldCnt: missing,
		NextSafeAction:          nextSafeAction,
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}
