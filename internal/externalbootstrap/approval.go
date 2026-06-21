package externalbootstrap

import (
	"github.com/kimsemi-home/myhome-jarvis/internal/authority"
	"github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"
)

func approvalUnlocksCandidate(
	status authority.ApprovalDecisionStatus,
	split externalevidence.RepoSplitDecisionPacket,
) bool {
	if !status.CanCreateRepo || status.UnrelatedAuthorityGranted {
		return false
	}
	for _, summary := range status.ScopeSummaries {
		if summary.Scope == "repo_creation" &&
			summary.Target == split.FutureRepoCandidate &&
			summary.LeaseState == "active" &&
			summary.CanUnlockScopeOnly &&
			summary.GrantFlags.RepoCreationAllowed &&
			!summary.GrantFlags.WorkflowChangesAllowed &&
			!summary.GrantFlags.ExternalWritesAllowed &&
			!summary.GrantFlags.SelfApprovalAllowed {
			return true
		}
	}
	return false
}

func approvalLeaseExpiresAt(
	status authority.ApprovalDecisionStatus,
	split externalevidence.RepoSplitDecisionPacket,
) string {
	for _, summary := range status.ScopeSummaries {
		if summary.Scope == "repo_creation" &&
			summary.Target == split.FutureRepoCandidate &&
			summary.LeaseState == "active" {
			return summary.ExpiresAt
		}
	}
	return ""
}
