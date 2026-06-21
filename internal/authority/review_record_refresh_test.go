package authority

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRefreshReviewRequestRecordsPendingNonApproval(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	copyGeneratedFixtures(t, root)
	writePolicy(t, root, policy)
	commitPublicFixture(t, root)
	writeReviewRequestableEvidence(t, root)

	result, err := RefreshReviewRequest(root, "kim-210")
	if err != nil {
		t.Fatal(err)
	}
	if result.LinearIssueRef != "KIM-210" ||
		result.LedgerState != "recorded_private" ||
		result.ApprovalState != "not_approved" ||
		result.QueueState != "pending_human_review" ||
		result.ApprovalGranted ||
		result.ExternalWritesAllowed ||
		result.SelfApprovalAllowed ||
		result.RequiredReviewClassCount == 0 ||
		result.RecordedAt == "" ||
		!result.PublicSafe {
		t.Fatalf("refresh result = %#v", result)
	}

	ledgerPath := filepath.Join(root, filepath.FromSlash(policy.PrivateReviewRequestLedger))
	body, err := os.ReadFile(ledgerPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), result.RequestID) ||
		!strings.Contains(string(body), `"approval_granted":false`) ||
		!strings.Contains(string(body), `"external_writes_allowed":false`) ||
		!strings.Contains(string(body), `"self_approval_allowed":false`) {
		t.Fatalf("ledger row did not preserve pending non-approval: %s", body)
	}
}

func TestRefreshReviewRequestResultRedactsPrivateRefs(t *testing.T) {
	root := t.TempDir()
	copyGeneratedFixtures(t, root)
	writePolicy(t, root, testPolicy())
	commitPublicFixture(t, root)
	writeReviewRequestableEvidence(t, root)

	result, err := RefreshReviewRequest(root, "KIM-210")
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"evidence_ref", "queue_item_ref", "data/private", "requests.jsonl",
		"/" + "Users" + "/", "token", "credential",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("refresh result leaked %q in %s", forbidden, body)
		}
	}
}
