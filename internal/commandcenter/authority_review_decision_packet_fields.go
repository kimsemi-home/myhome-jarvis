package commandcenter

func authorityReviewDecisionPacketPublicSafe(
	brief AuthorityReviewBrief,
	status Status,
) bool {
	return brief.PublicSafe &&
		status.StorageArchive.PublicSafe &&
		status.StorageArchive.ConfigIsEvidence &&
		brief.LocalRuntime.PublicSafe &&
		brief.MergeEvidence.PublicSafe &&
		brief.CodexSustainability.PublicSafe &&
		brief.ContextPack.PublicSafe &&
		brief.CapabilityReadiness.PublicSafe &&
		!brief.LocalRuntime.RawRuntimePublicAllowed &&
		!brief.ApprovalBoundary.ApprovalGranted &&
		!brief.ApprovalBoundary.ExternalWritesAllowed &&
		!brief.ApprovalBoundary.SelfApprovalAllowed
}

func authorityReviewPublicSafety(status Status) AuthorityReviewPublicSafety {
	return AuthorityReviewPublicSafety{
		PublicRepoMode:                  status.Authority.PublicRepoMode,
		PublicSafetyOK:                  status.Authority.PublicSafetyOK,
		HighRiskBlockedDecisionCount:    status.AuthorityReview.HighRiskBlockedDecisionCount,
		PublicRepoReviewProfileCount:    status.AuthorityReview.PublicRepoReviewProfileCount,
		WorkflowReviewProfileCount:      status.AuthorityReview.WorkflowReviewProfileCount,
		SelfApprovalBlockedProfileCount: status.AuthorityReview.SelfApprovalBlockedProfileCount,
	}
}

func authorityReviewDecisionOptions() []AuthorityReviewDecisionOption {
	return []AuthorityReviewDecisionOption{
		authorityReviewDecisionOption(
			"keep_pending_review", "Keep pending",
			"no_authority_change",
			false, false,
		),
		authorityReviewDecisionOption(
			"approve_after_human_review", "Approve after review",
			"requires_separate_record_command",
			true, true,
		),
		authorityReviewDecisionOption(
			"reject_or_request_changes", "Reject or request changes",
			"keeps_gates_blocked",
			true, true,
		),
	}
}

func authorityReviewDecisionOption(
	key string,
	label string,
	effect string,
	requiresHuman bool,
	requiresRecord bool,
) AuthorityReviewDecisionOption {
	return AuthorityReviewDecisionOption{
		Key:                           key,
		Label:                         label,
		Effect:                        effect,
		RequiresHuman:                 requiresHuman,
		RequiresSeparateRecordCommand: requiresRecord,
		ThisPacketGrantsApproval:      false,
		AllowsExternalWrites:          false,
		AllowsRepoCreation:            false,
		AllowsWorkflowChanges:         false,
		AllowsSelfApproval:            false,
	}
}
