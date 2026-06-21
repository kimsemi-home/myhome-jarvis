package commandcenter

func authorityReviewDecisionContractPublicSafe(
	brief AuthorityReviewBrief,
	items []AuthorityReviewDecisionContractItem,
	checks []AuthorityReviewContractEvidenceCheck,
) bool {
	if !brief.PublicSafe || brief.ApprovalBoundary.ApprovalGranted ||
		brief.ApprovalBoundary.ExternalWritesAllowed ||
		brief.ApprovalBoundary.SelfApprovalAllowed {
		return false
	}
	return authorityReviewContractItemsPublicSafe(items) &&
		authorityReviewContractChecksPublicSafe(checks)
}

func authorityReviewContractItemsPublicSafe(
	items []AuthorityReviewDecisionContractItem,
) bool {
	for _, item := range items {
		if item.ThisPacketGrantsApproval || item.AllowsExternalWrites ||
			item.AllowsRepoCreation || item.AllowsWorkflowChanges ||
			item.AllowsSelfApproval {
			return false
		}
	}
	return true
}

func authorityReviewContractChecksPublicSafe(
	checks []AuthorityReviewContractEvidenceCheck,
) bool {
	for _, check := range checks {
		if check.Required && !check.PublicSafe {
			return false
		}
	}
	return true
}
