package externalbootstrap

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/repofactory"
)

func creationDecision(
	split externalevidence.RepoSplitDecisionPacket,
	approval authority.ApprovalDecisionStatus,
	factory repofactory.DecisionPacket,
) (string, string) {
	switch {
	case !split.PublicSafe:
		return "blocked_repair_external_evidence_policy", "external_evidence_policy"
	case !factory.PublicSafe:
		return "blocked_repair_repo_factory_policy", "repo_factory_policy"
	case !factory.PublicSafetyEvidence.OK:
		return "blocked_collect_public_safety_evidence", "public_safety_evidence"
	case !factory.ContextPackEvidence.Valid:
		return "blocked_repair_context_handoff", "context_pack_handoff"
	case approval.UnrelatedAuthorityGranted:
		return "blocked_repair_authority_scope", "unrelated_authority"
	case !approvalUnlocksCandidate(approval, split):
		return "blocked_missing_repo_creation_approval", "authority_approval"
	default:
		return "ready_to_bootstrap_public_skeleton", ""
	}
}

func packetPublicSafe(
	split externalevidence.RepoSplitDecisionPacket,
	approval authority.ApprovalDecisionStatus,
	factory repofactory.DecisionPacket,
) bool {
	return split.PublicSafe && factory.PublicSafe && approval.PublicSafe &&
		!split.RawPayloadPublicAllowed && split.PrivateLakeStaysPrivate &&
		!split.ExternalWritesAllowed && !approval.UnrelatedAuthorityGranted
}

func nextSafeAction(decision string, reason string) string {
	if decision == "ready_to_bootstrap_public_skeleton" {
		return "bootstrap_minimal_public_repo_skeleton"
	}
	if reason == "authority_approval" {
		return "record_human_repo_creation_approval"
	}
	return "repair_" + reason
}
