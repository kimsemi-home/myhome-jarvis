package authority

import (
	"testing"
	"time"
)

func TestNormalizeReviewRecordRejectsApprovalGrant(t *testing.T) {
	policy := testPolicy()
	packet := ReviewRequestPacketFromPlan(ReviewPlan(policy, Assess(policy, clearInputs())))
	evidence := ReviewRequestEvidenceFromPacket(packet)
	queue := ReviewQueueStatusFromPacket(packet, evidence)
	trueValue := true
	falseValue := false

	_, err := normalizeReviewRecordRequest(policy, ReviewRecordRequest{
		RequestID:             packet.RequestID,
		EvidenceRef:           evidence.EvidenceRef,
		QueueItemRef:          queue.QueueItemRef,
		QueueState:            queue.QueueState,
		RequiredReviewClasses: packet.RequiredReviewClasses,
		ApprovalGranted:       &trueValue,
		ExternalWritesAllowed: &falseValue,
		SelfApprovalAllowed:   &falseValue,
	}, packet, evidence, queue, time.Now())
	if err == nil {
		t.Fatal("expected approval grant to be rejected")
	}
}

func TestNormalizeReviewRecordRejectsPrivateLocator(t *testing.T) {
	policy := testPolicy()
	packet := ReviewRequestPacketFromPlan(ReviewPlan(policy, Assess(policy, clearInputs())))
	evidence := ReviewRequestEvidenceFromPacket(packet)
	queue := ReviewQueueStatusFromPacket(packet, evidence)
	falseValue := false

	_, err := normalizeReviewRecordRequest(policy, ReviewRecordRequest{
		RequestID:             packet.RequestID,
		EvidenceRef:           evidence.EvidenceRef,
		QueueItemRef:          queue.QueueItemRef,
		QueueState:            queue.QueueState,
		RequiredReviewClasses: packet.RequiredReviewClasses,
		LinearIssueRef:        "private-linear-locator/KIM-171",
		ApprovalGranted:       &falseValue,
		ExternalWritesAllowed: &falseValue,
		SelfApprovalAllowed:   &falseValue,
	}, packet, evidence, queue, time.Now())
	if err == nil {
		t.Fatal("expected private locator to be rejected")
	}
}
