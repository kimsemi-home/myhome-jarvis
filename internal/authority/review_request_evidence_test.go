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
	if status.NextSafeAction != "request_authority_review" {
		t.Fatalf("next safe action = %q", status.NextSafeAction)
	}
}

func TestReviewRequestEvidenceFollowsRecordedReviewPlan(t *testing.T) {
	plan := ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs()))
	if plan.NextSafeAction != "request_authority_review" {
		t.Fatalf("pre-record plan next action = %q", plan.NextSafeAction)
	}
	applyReviewRecordLedger(&plan, ReviewRecordLedgerSummary{
		Recorded:       true,
		LedgerState:    "recorded_pending_review",
		ApprovalState:  "not_approved",
		LastRecordedAt: "2026-06-20T05:00:00Z",
	})

	status := ReviewRequestEvidenceFromPlan(plan)
	if status.NextSafeAction != "await_human_authority_review" {
		t.Fatalf("recorded evidence next action = %q", status.NextSafeAction)
	}
	if status.ApprovalGranted || status.ExternalWritesAllowed ||
		status.SelfApprovalAllowed {
		t.Fatalf("recorded evidence granted authority = %#v", status)
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
