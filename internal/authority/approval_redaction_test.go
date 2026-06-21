package authority

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestApprovalDecisionResultRedactsPrivateReviewContext(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	record, err := normalizeApprovalDecisionRequest(
		testPolicy(),
		approvalFixtureRequest(now),
		approvalFixturePacket(now),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(resultForApprovalDecision(record))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "reviewer_boundary") ||
		strings.Contains(string(body), "data/private") ||
		strings.Contains(string(body), "raw_private") {
		t.Fatalf("approval result leaked private context: %s", body)
	}
}

func TestApprovalStatusDoesNotExposeLedgerPath(t *testing.T) {
	status := ApprovalDecisionStatus{PublicSafe: true, LedgerState: "missing"}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "approvals.jsonl") ||
		strings.Contains(string(body), "data/private") {
		t.Fatalf("approval status leaked private path: %s", body)
	}
}
