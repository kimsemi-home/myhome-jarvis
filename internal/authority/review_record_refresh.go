package authority

import "time"

func RefreshReviewRequest(root string, linearIssueRef string) (ReviewRecordResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	packet, err := ReviewRequestPacketForRoot(root)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	evidence := ReviewRequestEvidenceFromPacket(packet)
	queue := ReviewQueueStatusFromPacket(packet, evidence)
	falseValue := false
	request := ReviewRecordRequest{
		RequestID:             packet.RequestID,
		EvidenceRef:           evidence.EvidenceRef,
		QueueItemRef:          queue.QueueItemRef,
		QueueState:            queue.QueueState,
		RequiredReviewClasses: packet.RequiredReviewClasses,
		LinearIssueRef:        linearIssueRef,
		ApprovalGranted:       &falseValue,
		ExternalWritesAllowed: &falseValue,
		SelfApprovalAllowed:   &falseValue,
	}
	record, err := normalizeReviewRecordRequest(
		policy, request, packet, evidence, queue, time.Now().UTC(),
	)
	if err != nil {
		return ReviewRecordResult{}, err
	}
	if err := appendReviewRecord(root, policy, record); err != nil {
		return ReviewRecordResult{}, err
	}
	return resultForReviewRecord(record), nil
}
