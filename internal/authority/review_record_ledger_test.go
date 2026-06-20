package authority

import (
	"testing"
	"time"
)

func TestReviewRecordLedgerSummarizesCurrentRequest(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	record := pendingReviewRecord("authority-review-123456abcdef")
	if err := appendReviewRecord(root, policy, record); err != nil {
		t.Fatal(err)
	}

	summary, err := ReviewRecordLedgerForRoot(root, policy, record.RequestID)
	if err != nil {
		t.Fatal(err)
	}
	if !summary.Recorded ||
		summary.LedgerState != "recorded_pending_review" ||
		summary.RecordCount != 1 ||
		summary.ApprovalState != "not_approved" {
		t.Fatalf("ledger summary = %#v", summary)
	}
}

func TestReviewRecordLedgerMissingKeepsRequestable(t *testing.T) {
	summary, err := ReviewRecordLedgerForRoot(
		t.TempDir(),
		testPolicy(),
		"authority-review-123456abcdef",
	)
	if err != nil {
		t.Fatal(err)
	}
	if summary.Recorded || summary.LedgerState != "missing" {
		t.Fatalf("ledger summary = %#v", summary)
	}
}

func TestApplyReviewRecordLedgerAwaitsHumanReview(t *testing.T) {
	policy := testPolicy()
	plan := ReviewPlan(policy, Assess(policy, clearInputs()))
	applyReviewRecordLedger(&plan, ReviewRecordLedgerSummary{
		Recorded:      true,
		RecordCount:   1,
		LedgerState:   "recorded_pending_review",
		ApprovalState: "not_approved",
	})
	if plan.NextSafeAction != "await_human_authority_review" ||
		!plan.ReviewRequestRecorded {
		t.Fatalf("review plan = %#v", plan)
	}
}

func pendingReviewRecord(requestID string) ReviewRecord {
	return ReviewRecord{
		At:                       time.Date(2026, 6, 20, 5, 0, 0, 0, time.UTC).Format(time.RFC3339),
		RequestID:                requestID,
		RequestState:             "ready",
		EvidenceRef:              "authority_review_request:" + requestID,
		EvidenceState:            "ready_to_attach",
		QueueItemRef:             "authority_review_queue:" + requestID,
		QueueState:               "pending_human_review",
		ApprovalState:            "not_approved",
		RequiredReviewClassCount: 5,
		PublicSafe:               true,
	}
}
