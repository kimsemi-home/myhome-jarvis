package authority

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRecordApprovalDecisionWritesPrivateScopedLease(t *testing.T) {
	root := t.TempDir()
	writeApprovalRootPolicies(t, root)
	now := time.Now().UTC()
	payload, err := json.Marshal(approvalFixtureRequest(now))
	if err != nil {
		t.Fatal(err)
	}
	result, err := RecordApprovalDecision(root, payload)
	if err != nil {
		t.Fatal(err)
	}
	if !result.PublicSafe || !result.CanUnlockScopeOnly ||
		result.LedgerState != "recorded_private" {
		t.Fatalf("approval result = %#v", result)
	}
	path := filepath.Join(
		root,
		filepath.FromSlash(testPolicy().PrivateApprovalDecisionLedger),
	)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("ledger permissions = %v", info.Mode().Perm())
	}
}

func TestApprovalStatusUnlocksOnlyMatchingScope(t *testing.T) {
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
	if err := appendApprovalDecision(root, policy, record); err != nil {
		t.Fatal(err)
	}
	status, err := approvalDecisionStatusForRoot(root, policy, now)
	if err != nil {
		t.Fatal(err)
	}
	if !status.CanCreateRepo || status.CanChangeWorkflow ||
		status.CanWriteExternal || status.UnrelatedAuthorityGranted ||
		status.ActiveApprovalCount != 1 || len(status.ScopeSummaries) != 1 {
		t.Fatalf("approval status = %#v", status)
	}
}
