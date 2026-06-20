package authority

import "time"

func ReviewRequestEvidenceForRoot(root string) (ReviewRequestEvidenceStatus, error) {
	packet, err := ReviewRequestPacketForRoot(root)
	if err != nil {
		return ReviewRequestEvidenceStatus{}, err
	}
	return ReviewRequestEvidenceFromPacket(packet), nil
}

func ReviewRequestEvidenceFromPacket(packet ReviewRequestPacket) ReviewRequestEvidenceStatus {
	missing := missingReviewEvidenceFieldCount(packet)
	ready := reviewRequestEvidenceReady(packet, missing)
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
		NextSafeAction:          packet.SourceAction,
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}
