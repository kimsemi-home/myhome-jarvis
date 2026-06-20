package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewBriefPublicSafe(
	plan authority.ReviewPlanStatus,
	status Status,
) bool {
	return plan.PublicSafe &&
		status.PublicSafe &&
		status.AuthorityReview.PublicSafe &&
		status.WorkItem.PublicSafe &&
		status.RepoFactory.PublicSafe &&
		status.RepoFactoryPreflight.PublicSafe &&
		status.LocalRuntime.PublicSafe &&
		status.MergeEvidence.PublicSafe &&
		status.CodexSustainability.PublicSafe &&
		status.ContextPack.PublicSafe &&
		!status.RepoFactoryPreflight.CreationAllowed &&
		!status.LocalRuntime.RawRuntimePublicAllowed &&
		!status.WorkItem.ApprovalGranted &&
		!status.WorkItem.ExternalWritesAllowed &&
		!status.WorkItem.SelfApprovalAllowed
}

func authorityReviewBoundary(work WorkItemSummary) AuthorityReviewBoundary {
	return AuthorityReviewBoundary{
		ApprovalState:         work.ApprovalState,
		ApprovalGranted:       work.ApprovalGranted,
		ExternalWritesAllowed: work.ExternalWritesAllowed,
		SelfApprovalAllowed:   work.SelfApprovalAllowed,
		ReviewOnly:            work.ReviewOnly,
	}
}

func authorityReviewRepoFactory(summary RepoFactorySummary) AuthorityReviewRepoFactory {
	return AuthorityReviewRepoFactory{
		PublicSafe:                     summary.PublicSafe,
		AuthorityReviewRequired:        summary.AuthorityReviewRequired,
		PublicSafetyEvidenceRequired:   summary.PublicSafetyEvidenceRequired,
		RepoCreationBlockedUntilReview: summary.RepoCreationBlockedUntilReview,
		MissingCreationGateCount:       summary.MissingCreationGateCount,
		ForbiddenTemplateValueCount:    summary.ForbiddenTemplateValueCount,
	}
}

func authorityReviewBriefNextAction(status Status) string {
	if status.WorkItem.NextSafeAction != "" {
		return status.WorkItem.NextSafeAction
	}
	return status.NextSafeAction
}
