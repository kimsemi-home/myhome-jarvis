package authority

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReviewRequestPacketHasStablePublicID(t *testing.T) {
	plan := ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs()))
	first := ReviewRequestPacketFromPlan(plan)
	plan.CheckedAt = "2099-01-01T00:00:00Z"
	second := ReviewRequestPacketFromPlan(plan)

	if first.RequestID != second.RequestID || !strings.HasPrefix(first.RequestID, "authority-review-") {
		t.Fatalf("request ids = %q %q", first.RequestID, second.RequestID)
	}
	if first.RequestState != "ready" || first.NextHandling == "" {
		t.Fatalf("request packet = %#v", first)
	}
	if first.ApprovalGranted || first.ExternalWritesAllowed || first.SelfApprovalAllowed {
		t.Fatalf("request packet granted authority = %#v", first)
	}
}

func TestReviewRequestPacketRedactsPrivateFields(t *testing.T) {
	packet := ReviewRequestPacketFromPlan(ReviewPlan(testPolicy(), Assess(testPolicy(), clearInputs())))
	body, err := json.Marshal(packet)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_rationale", "raw_evidence", "reviewer_identity",
		"linear_url", "local_absolute_path", "token", "credential",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("review request leaked %q in %s", forbidden, body)
		}
	}
}
