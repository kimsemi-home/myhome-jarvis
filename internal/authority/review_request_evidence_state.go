package authority

func reviewRequestEvidenceReady(packet ReviewRequestPacket, missing int) bool {
	return packet.PublicSafe &&
		packet.RequestID != "" &&
		packet.RequestState == "ready" &&
		!packet.ApprovalGranted &&
		!packet.ExternalWritesAllowed &&
		!packet.SelfApprovalAllowed &&
		missing == 0
}

func reviewRequestEvidenceState(ready bool) string {
	if ready {
		return "ready_to_attach"
	}
	return "blocked"
}

func missingReviewEvidenceFieldCount(packet ReviewRequestPacket) int {
	missing := 0
	if packet.PolicyPath == "" {
		missing++
	}
	if len(packet.RequiredReviewClasses) == 0 {
		missing++
	}
	if packet.HighRiskBlockedDecisionCount == 0 {
		missing++
	}
	if len(packet.IncludedEvidenceFields) == 0 {
		missing++
	}
	return missing
}
