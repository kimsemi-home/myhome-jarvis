package externalbootstrap

import (
	"testing"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
)

func TestPacketBlocksWithoutApprovalLease(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	packet, err := packetFromEvidence(
		repoRoot(t),
		splitFixture(now),
		approvalMissingFixture(now),
		factoryFixture(),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	if packet.CreationAllowed ||
		packet.CreationDecision != "blocked_missing_repo_creation_approval" ||
		packet.CreationBlockedReason != "authority_approval" {
		t.Fatalf("bootstrap packet = %#v", packet)
	}
	if packet.CandidateRepo != "kimsemi-home/myhome-external-evidence-lake" ||
		packet.NextSafeAction != "record_human_repo_creation_approval" {
		t.Fatalf("bootstrap target = %#v", packet)
	}
}

func approvalMissingFixture(now time.Time) authority.ApprovalDecisionStatus {
	return authority.ApprovalDecisionStatus{
		PublicSafe:       true,
		LedgerState:      "missing",
		NextSafeAction:   "record_human_approval_decision",
		CheckedAt:        now.Format(time.RFC3339),
		ScopeSummaries:   nil,
		CanCreateRepo:    false,
		CanWriteExternal: false,
	}
}

func TestPacketAllowsOnlyExactRepoCreationApproval(t *testing.T) {
	now := time.Date(2026, 6, 21, 6, 0, 0, 0, time.UTC)
	packet, err := packetFromEvidence(
		repoRoot(t),
		splitFixture(now),
		approvalFixture(now),
		factoryFixture(),
		now,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !packet.CreationAllowed ||
		packet.CreationDecision != "ready_to_bootstrap_public_skeleton" ||
		packet.NextSafeAction != "bootstrap_minimal_public_repo_skeleton" {
		t.Fatalf("bootstrap packet = %#v", packet)
	}
	if packet.ExternalWritesAllowed || packet.RawPayloadPublicAllowed ||
		!packet.PrivateLakeStaysPrivate || !packet.ApprovalUnlocksScopeOnly {
		t.Fatalf("bootstrap safety = %#v", packet)
	}
}
