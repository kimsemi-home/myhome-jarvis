package externalbootstrap

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

func PacketForRoot(root string) (Packet, error) {
	split, err := externalevidence.RepoSplitDecisionPacketForRoot(root)
	if err != nil {
		return Packet{}, err
	}
	approval, err := authority.ApprovalDecisionStatusForRoot(root)
	if err != nil {
		return Packet{}, err
	}
	factory, err := repofactory.DecisionPacketForRoot(root)
	if err != nil {
		return Packet{}, err
	}
	return packetFromEvidence(root, split, approval, factory, time.Now().UTC())
}

func packetFromEvidence(
	root string,
	split externalevidence.RepoSplitDecisionPacket,
	approval authority.ApprovalDecisionStatus,
	factory repofactory.DecisionPacket,
	now time.Time,
) (Packet, error) {
	decision, reason := creationDecision(split, approval, factory)
	allowed := decision == "ready_to_bootstrap_public_skeleton"
	packet := Packet{
		Context:                  "ExternalEvidenceRepoBootstrapPacket",
		Version:                  "v1",
		PublicSafe:               packetPublicSafe(split, approval, factory),
		CandidateRepo:            split.FutureRepoCandidate,
		CreationDecision:         decision,
		CreationAllowed:          allowed,
		CreationBlockedReason:    reason,
		RequiredApprovalScope:    "repo_creation",
		ApprovalLedgerState:      approval.LedgerState,
		ApprovalLeaseState:       approval.LatestLeaseState,
		ApprovalLeaseExpiresAt:   approvalLeaseExpiresAt(approval, split),
		ApprovalUnlocksScopeOnly: approvalUnlocksCandidate(approval, split),
		RepoSplitDecisionState:   split.DecisionState,
		RepoSplitCheckedAt:       split.CheckedAt,
		RepoFactoryDecision:      factory.CreationDecision,
		ContextHandoff:           contextHandoff(factory),
		SkeletonFiles:            skeletonFiles(factory),
		PrivateLakeStaysPrivate:  split.PrivateLakeStaysPrivate,
		RawPayloadPublicAllowed:  split.RawPayloadPublicAllowed,
		ExternalWritesAllowed:    split.ExternalWritesAllowed,
		NextSafeAction:           nextSafeAction(decision, reason),
		CheckedAt:                now.Format(time.RFC3339),
	}
	var err error
	packet.HashCacheInputs, err = hashCacheInputs(root, factory)
	return packet, err
}
