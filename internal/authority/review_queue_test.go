package authority

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReviewQueueStatusIsPendingWithoutApproval(t *testing.T) {
	packet := ReviewRequestPacketFromPlan(ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs())))
	status := ReviewQueueStatusFromPacket(packet, ReviewRequestEvidenceFromPacket(packet))

	if !status.QueueReady || status.QueueState != "pending_human_review" {
		t.Fatalf("queue status = %#v", status)
	}
	if status.PendingReviewClassCount != len(packet.RequiredReviewClasses) {
		t.Fatalf("pending review class count = %#v", status)
	}
	if status.ApprovalGranted || status.ExternalWritesAllowed || status.SelfApprovalAllowed {
		t.Fatalf("queue status granted authority = %#v", status)
	}
}

func TestReviewQueueStatusRedactsPrivateFields(t *testing.T) {
	packet := ReviewRequestPacketFromPlan(ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs())))
	body, err := json.Marshal(ReviewQueueStatusFromPacket(packet, ReviewRequestEvidenceFromPacket(packet)))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_rationale", "raw_evidence", "reviewer_identity",
		"linear_url", "local_absolute_path", "token", "credential",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("review queue leaked %q in %s", forbidden, body)
		}
	}
}
