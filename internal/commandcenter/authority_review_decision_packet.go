package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func AuthorityReviewDecisionPacketForRoot(
	root string,
) (AuthorityReviewDecisionPacket, error) {
	status, err := StatusForRoot(root)
	if err != nil {
		return AuthorityReviewDecisionPacket{}, err
	}
	plan, err := authority.ReviewPlanForRoot(root)
	if err != nil {
		return AuthorityReviewDecisionPacket{}, err
	}
	brief := authorityReviewBriefFromStatus(plan, status)
	return authorityReviewDecisionPacketFromStatus(brief, status), nil
}

func authorityReviewDecisionPacketFromStatus(
	brief AuthorityReviewBrief,
	status Status,
) AuthorityReviewDecisionPacket {
	return AuthorityReviewDecisionPacket{
		Context:                      "AuthorityReviewDecisionPacket",
		Version:                      "v1",
		PublicSafe:                   authorityReviewDecisionPacketPublicSafe(brief, status),
		Redaction:                    "review-decision-public-handoff",
		PolicyPath:                   brief.PolicyPath,
		RequestID:                    brief.RequestID,
		RequestState:                 brief.RequestState,
		QueueState:                   brief.QueueState,
		EvidenceReady:                brief.EvidenceReady,
		ReviewRequestRecorded:        brief.ReviewRequestRecorded,
		ReviewRequestLedgerState:     brief.ReviewRequestLedgerState,
		ReviewRequestAgeHours:        brief.ReviewRequestAgeHours,
		ReviewRequestStaleAfterHours: brief.ReviewRequestStaleAfterHours,
		ReviewRequestStale:           brief.ReviewRequestStale,
		ReviewEscalationAction:       brief.ReviewEscalationAction,
		RequiredReviewClasses:        append([]string{}, brief.RequiredReviewClasses...),
		RequiredReviewClassCount:     brief.RequiredReviewClassCount,
		GatedCapabilityKeys:          append([]string{}, brief.GatedCapabilityKeys...),
		BlockedGateKeys:              append([]string{}, brief.BlockedGateKeys...),
		ApprovalBoundary:             brief.ApprovalBoundary,
		RepoFactoryGate:              brief.RepoFactoryGate,
		RepoFactoryPreflight:         brief.RepoFactoryPreflight,
		PublicSafetyPosture:          authorityReviewPublicSafety(status),
		StorageEvidence:              status.StorageArchive,
		LocalRuntime:                 brief.LocalRuntime,
		MergeEvidence:                brief.MergeEvidence,
		CodexSustainability:          brief.CodexSustainability,
		ContextPack:                  brief.ContextPack,
		CapabilityReadiness:          brief.CapabilityReadiness,
		DecisionPacketState:          "review_only",
		CanApplyDecision:             false,
		DecisionOptions:              authorityReviewDecisionOptions(),
		NextSafeAction:               brief.NextSafeAction,
		CheckedAt:                    brief.CheckedAt,
	}
}
