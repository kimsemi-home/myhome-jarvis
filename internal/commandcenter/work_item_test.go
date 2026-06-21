package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestWorkItemStatusUsesUniversalVocabulary(t *testing.T) {
	status, err := WorkItemForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if status.Context != "UniversalWorkItem" || status.WorkItemRef == "" {
		t.Fatalf("work item status = %#v", status)
	}
	if status.IntentKey != "closed_loop_next_safe_action" ||
		status.DecisionKey == "" || len(status.CapabilityKeys) == 0 {
		t.Fatalf("work item vocabulary = %#v", status)
	}
	if status.EvidenceRef == "" || status.AuthorityRef == "" ||
		len(status.GuardrailKeys) == 0 {
		t.Fatalf("work item refs = %#v", status)
	}
	if status.ReviewStaleAfterHours != 24 ||
		status.ReviewEscalationAction == "" {
		t.Fatalf("work item stale guard = %#v", status)
	}
	if !status.CapabilityReadiness.PublicSafe ||
		status.CapabilityReadiness.Media.State == "" ||
		status.CapabilityReadiness.FinanceConsent.State == "" ||
		status.CapabilityReadiness.Monetization.State == "" ||
		status.CapabilityReadiness.CodexCost.State == "" {
		t.Fatalf("work item capability readiness = %#v", status.CapabilityReadiness)
	}
	if status.ApprovalGranted || status.ExternalWritesAllowed ||
		status.SelfApprovalAllowed || status.ApprovalState != "not_approved" {
		t.Fatalf("work item granted authority = %#v", status)
	}
}

func TestWorkItemStatusRedactsPrivateFields(t *testing.T) {
	status, err := WorkItemForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_rationale", "raw_evidence", "raw_prompt", "raw_transcript",
		"reviewer_identity", "linear_url", "local_absolute_path",
		"token", "credential", "cookie", "account_id", "finance_payload",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("work item leaked %q in %s", forbidden, body)
		}
	}
}
