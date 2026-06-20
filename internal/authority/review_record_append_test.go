package authority

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppendReviewRecordWritesPrivateJSONL(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
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
	if err := appendReviewRecord(root, policy, record); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateReviewRequestLedger))
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("ledger permissions = %v", info.Mode().Perm())
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), record.RequestID) {
		t.Fatalf("ledger body = %s", body)
	}
}
