package authority

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReviewRecordResultRedactsEvidenceRefs(t *testing.T) {
	record := ReviewRecord{
		At:                       "2026-06-20T05:00:00Z",
		RequestID:                "authority-review-123456abcdef",
		EvidenceRef:              "authority_review_request:authority-review-123456abcdef",
		QueueItemRef:             "authority_review_queue:authority-review-123456abcdef",
		QueueState:               "pending_human_review",
		RequiredReviewClassCount: 5,
		ApprovalState:            "not_approved",
		PublicSafe:               true,
	}
	body, err := json.Marshal(resultForReviewRecord(record))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "evidence_ref") ||
		strings.Contains(string(body), "queue_item_ref") {
		t.Fatalf("review record result leaked refs: %s", body)
	}
}
