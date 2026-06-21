package authority

import (
	"testing"
	"time"
)

func TestApprovalStatusCountsExpiredLeases(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	record, err := normalizeApprovalDecisionRequest(
		policy,
		approvalFixtureRequest(now.Add(-2*time.Hour)),
		approvalFixturePacket(now.Add(-2*time.Hour)),
		now.Add(-2*time.Hour),
	)
	if err != nil {
		t.Fatal(err)
	}
	record.ExpiresAt = now.Add(-time.Minute).Format(time.RFC3339)
	if err := appendApprovalDecision(root, policy, record); err != nil {
		t.Fatal(err)
	}
	status, err := approvalDecisionStatusForRoot(root, policy, now)
	if err != nil {
		t.Fatal(err)
	}
	if status.ActiveApprovalCount != 0 ||
		status.ExpiredApprovalCount != 1 ||
		status.CanCreateRepo ||
		status.NextSafeAction != "record_human_approval_decision" {
		t.Fatalf("approval status = %#v", status)
	}
}

func TestApprovalStatusRejectsUnscopedLedgerRecords(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	record, err := normalizeApprovalDecisionRequest(
		policy,
		approvalFixtureRequest(now),
		approvalFixturePacket(now),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	record.GrantFlags.WorkflowChangesAllowed = true
	if err := appendApprovalDecision(root, policy, record); err != nil {
		t.Fatal(err)
	}
	status, err := approvalDecisionStatusForRoot(root, policy, now)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidRecordCount != 1 || status.ActiveApprovalCount != 0 {
		t.Fatalf("approval status = %#v", status)
	}
}
