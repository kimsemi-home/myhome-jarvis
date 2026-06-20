package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStatusSummarizesAssistantCommandCenter(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if status.Vision.CapabilityCount != 6 || status.Vision.GuardrailCount == 0 {
		t.Fatalf("vision summary = %#v", status.Vision)
	}
	if !status.PDCA.Ready || status.Cost.BudgetState != "ok" {
		t.Fatalf("pdca/cost summary = %#v %#v", status.PDCA, status.Cost)
	}
	if !status.EvidenceIntegrity.PublicSafe || status.EvidenceIntegrity.NextSafeAction == "" {
		t.Fatalf("evidence integrity summary = %#v", status.EvidenceIntegrity)
	}
	if !status.AuthorityReview.PublicSafe ||
		status.AuthorityReview.HighRiskBlockedDecisionCount != 6 {
		t.Fatalf("authority review summary = %#v", status.AuthorityReview)
	}
	if status.AuthorityReview.RequestID == "" || status.AuthorityReview.RequestState == "" {
		t.Fatalf("authority review request summary = %#v", status.AuthorityReview)
	}
	if status.AuthorityReview.EvidenceRef == "" || status.AuthorityReview.EvidenceState == "" {
		t.Fatalf("authority review evidence summary = %#v", status.AuthorityReview)
	}
	if status.AuthorityReview.QueueState == "" || status.AuthorityReview.PendingReviewClassCount == 0 {
		t.Fatalf("authority review queue summary = %#v", status.AuthorityReview)
	}
	if !status.MergeEvidence.MergeReady || status.MergeEvidence.RequiredEvidenceCount != 8 {
		t.Fatalf("merge evidence summary = %#v", status.MergeEvidence)
	}
	if status.BlockedGateCount == 0 || status.NextSafeAction == "" {
		t.Fatalf("expected a safe next action with gates: %#v", status)
	}
	if status.WorkItem.WorkItemRef == "" || status.WorkItem.WorkItemState == "" {
		t.Fatalf("work item summary = %#v", status.WorkItem)
	}
	if status.CompactState != "blocked" && status.CompactState != "gated" {
		t.Fatalf("compact state = %q", status.CompactState)
	}
}

func TestStatusDoesNotExposePrivatePayloadFields(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_prompt", "raw_transcript", "private_notes",
		"token", "secret", "credential", "local_absolute_path",
		"reviewer_identity", "linear_url", "finance_payload",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("command center leaked %q in %s", forbidden, body)
		}
	}
}
