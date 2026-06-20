package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func AuthorityReviewBriefForRoot(root string) (AuthorityReviewBrief, error) {
	status, err := StatusForRoot(root)
	if err != nil {
		return AuthorityReviewBrief{}, err
	}
	plan, err := authority.ReviewPlanForRoot(root)
	if err != nil {
		return AuthorityReviewBrief{}, err
	}
	return authorityReviewBriefFromStatus(plan, status), nil
}

func authorityReviewBriefFromStatus(
	plan authority.ReviewPlanStatus,
	status Status,
) AuthorityReviewBrief {
	return AuthorityReviewBrief{
		Context:                      "AuthorityReviewBrief",
		Version:                      "v1",
		PublicSafe:                   authorityReviewBriefPublicSafe(plan, status),
		Redaction:                    "review-brief-public-handoff",
		PolicyPath:                   plan.PolicyPath,
		RequestID:                    status.AuthorityReview.RequestID,
		RequestState:                 status.AuthorityReview.RequestState,
		EvidenceRef:                  status.AuthorityReview.EvidenceRef,
		EvidenceReady:                status.AuthorityReview.EvidenceReady,
		QueueState:                   status.AuthorityReview.QueueState,
		QueueReady:                   status.AuthorityReview.QueueReady,
		ReviewRequestRecorded:        status.AuthorityReview.ReviewRequestRecorded,
		ReviewRequestLedgerState:     status.AuthorityReview.ReviewRequestLedgerState,
		ReviewRequestAgeHours:        status.AuthorityReview.ReviewRequestAgeHours,
		ReviewRequestStaleAfterHours: status.AuthorityReview.ReviewRequestStaleAfterHours,
		ReviewRequestStale:           status.AuthorityReview.ReviewRequestStale,
		ReviewEscalationAction:       status.AuthorityReview.ReviewRequestEscalationAction,
		RequiredReviewClasses:        append([]string{}, plan.RequiredReviewClasses...),
		RequiredReviewClassCount:     len(plan.RequiredReviewClasses),
		GatedCapabilityKeys:          append([]string{}, status.WorkItem.CapabilityKeys...),
		BlockedGateKeys:              append([]string{}, status.WorkItem.BlockedGateKeys...),
		WorkItemRef:                  status.WorkItem.WorkItemRef,
		WorkItemState:                status.WorkItem.WorkItemState,
		DecisionKey:                  status.WorkItem.DecisionKey,
		AuthorityRef:                 status.WorkItem.AuthorityRef,
		ApprovalBoundary:             authorityReviewBoundary(status.WorkItem),
		RepoFactoryGate:              authorityReviewRepoFactory(status.RepoFactory),
		RepoFactoryPreflight:         status.RepoFactoryPreflight,
		LocalRuntime:                 status.LocalRuntime,
		MergeEvidence:                status.MergeEvidence,
		CodexSustainability:          status.CodexSustainability,
		VisionGoalComplete:           visionGoalComplete(status),
		VisionNextSafeAction:         status.NextSafeAction,
		NextSafeAction:               authorityReviewBriefNextAction(status),
		CheckedAt:                    status.CheckedAt,
	}
}
