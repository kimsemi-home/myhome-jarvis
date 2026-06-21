package externalbootstrap

import (
	"testing"
	"time"
)

func TestPacketRejectsApprovalForDifferentRepo(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	approval := approvalFixture(now)
	approval.ScopeSummaries[0].Target = "kimsemi-home/other-repo"
	approval.LatestTarget = "kimsemi-home/other-repo"
	packet, err := packetFromEvidence(
		repoRoot(t),
		splitFixture(now),
		approval,
		factoryFixture(),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	if packet.CreationAllowed ||
		packet.CreationBlockedReason != "authority_approval" {
		t.Fatalf("bootstrap packet = %#v", packet)
	}
	if packet.ApprovalLeaseExpiresAt != "" {
		t.Fatalf("wrong-target lease leaked into packet = %#v", packet)
	}
}

func TestPacketRejectsUnrelatedAuthority(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	approval := approvalFixture(now)
	approval.UnrelatedAuthorityGranted = true
	approval.ScopeSummaries[0].GrantFlags.WorkflowChangesAllowed = true
	approval.ScopeSummaries[0].CanUnlockScopeOnly = false
	packet, err := packetFromEvidence(
		repoRoot(t),
		splitFixture(now),
		approval,
		factoryFixture(),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	if packet.CreationAllowed ||
		packet.CreationBlockedReason != "unrelated_authority" {
		t.Fatalf("bootstrap packet = %#v", packet)
	}
}
