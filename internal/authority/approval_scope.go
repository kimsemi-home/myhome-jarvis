package authority

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func validateApprovalScope(
	record ApprovalDecisionRecord,
	packet externalevidence.RepoSplitDecisionPacket,
) error {
	if !approvalFlagsMatchScope(record) {
		return fmt.Errorf("approval grant flags do not match scope")
	}
	if record.Scope != "repo_creation" &&
		record.Scope != "workflow_change" &&
		record.Scope != "external_write" {
		return fmt.Errorf("approval scope is not supported")
	}
	if record.Scope == "repo_creation" &&
		record.Target != packet.FutureRepoCandidate {
		return fmt.Errorf("repo creation target does not match packet")
	}
	return nil
}

func validateRepoSplitPacketRef(
	record ApprovalDecisionRecord,
	packet externalevidence.RepoSplitDecisionPacket,
) error {
	if record.DecisionPacketRef != "external_evidence_repo_split_decision" ||
		record.DecisionPacketContext != packet.Context ||
		!packet.PublicSafe || packet.CanCreateRepo {
		return fmt.Errorf("approval must reference review-only repo split packet")
	}
	return nil
}
