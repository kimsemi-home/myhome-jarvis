package authority

import (
	"testing"
	"time"
)

func TestNormalizeReviewRecordRequiresPendingNonApproval(t *testing.T) {
	policy := testPolicy()
	packet := ReviewRequestPacketFromPlan(ReviewPlan(policy, Assess(policy, clearInputs())))
	evidence := ReviewRequestEvidenceFromPacket(packet)
	queue := ReviewQueueStatusFromPacket(packet, evidence)
	falseValue := false

	record, err := normalizeReviewRecordRequest(policy, ReviewRecordRequest{
		At:                    "2026-06-20T14:00:00+09:00",
		RequestID:             packet.RequestID,
		EvidenceRef:           evidence.EvidenceRef,
		QueueItemRef:          queue.QueueItemRef,
		QueueState:            queue.QueueState,
		RequiredReviewClasses: packet.RequiredReviewClasses,
		LinearIssueRef:        "kim-171",
		ApprovalGranted:       &falseValue,
		ExternalWritesAllowed: &falseValue,
		SelfApprovalAllowed:   &falseValue,
	}, packet, evidence, queue, time.Date(2026, 6, 20, 5, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if record.LinearIssueRef != "KIM-171" ||
		record.QueueState != "pending_human_review" ||
		record.ApprovalGranted ||
		record.ExternalWritesAllowed ||
		record.SelfApprovalAllowed {
		t.Fatalf("review record = %#v", record)
	}
	if record.At != "2026-06-20T05:00:00Z" ||
		record.RequiredReviewClassCount != len(packet.RequiredReviewClasses) {
		t.Fatalf("review record summary = %#v", record)
	}
}
