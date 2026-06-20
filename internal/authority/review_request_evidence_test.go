package authority

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReviewRequestEvidenceIsReadyWithoutApproval(t *testing.T) {
	packet := ReviewRequestPacketFromPlan(ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs())))
	status := ReviewRequestEvidenceFromPacket(packet)

	if !status.EvidenceReady || status.EvidenceState != "ready_to_attach" {
		t.Fatalf("evidence status = %#v", status)
	}
	if status.EvidenceRef != "authority_review_request:"+packet.RequestID {
		t.Fatalf("evidence ref = %q", status.EvidenceRef)
	}
	if status.ApprovalState != "not_approved" || status.ApprovalGranted {
		t.Fatalf("approval state = %#v", status)
	}
}

func TestReviewRequestEvidenceRedactsPrivateFields(t *testing.T) {
	packet := ReviewRequestPacketFromPlan(ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs())))
	body, err := json.Marshal(ReviewRequestEvidenceFromPacket(packet))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_rationale", "raw_evidence", "reviewer_identity",
		"linear_url", "local_absolute_path", "token", "credential",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("review evidence leaked %q in %s", forbidden, body)
		}
	}
}
