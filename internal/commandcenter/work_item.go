package commandcenter

func WorkItemForRoot(root string) (WorkItemStatus, error) {
	status, err := StatusForRoot(root)
	if err != nil {
		return WorkItemStatus{}, err
	}
	return WorkItemStatus{
		Context:         "UniversalWorkItem",
		Version:         "v1",
		WorkItemSummary: status.WorkItem,
		CheckedAt:       status.CheckedAt,
	}, nil
}

func summarizeWorkItem(status Status) WorkItemSummary {
	gateKeys := blockedGateKeys(status.BlockedGates)
	return WorkItemSummary{
		WorkItemRef:            workItemRef(status),
		WorkItemState:          workItemState(status),
		IntentKey:              "closed_loop_next_safe_action",
		CapabilityKeys:         capabilityKeysForGates(gateKeys),
		DecisionKey:            decisionKey(status.NextSafeAction),
		EvidenceRef:            workItemEvidenceRef(status),
		AuthorityRef:           authorityRef(status),
		GuardrailKeys:          workItemGuardrails(),
		SourceAction:           status.NextSafeAction,
		BlockedGateKeys:        gateKeys,
		QueueState:             status.AuthorityReview.QueueState,
		ReviewClassCount:       status.AuthorityReview.PendingReviewClassCount,
		ReviewRequestAgeHours:  status.AuthorityReview.ReviewRequestAgeHours,
		ReviewStaleAfterHours:  status.AuthorityReview.ReviewRequestStaleAfterHours,
		ReviewRequestStale:     status.AuthorityReview.ReviewRequestStale,
		ReviewEscalationAction: status.AuthorityReview.ReviewRequestEscalationAction,
		MergeEligibilityHint:   mergeEligibilityHint(status),
		PublicSafe:             workItemPublicSafe(status),
		Redaction:              "universal-work-item-public-status",
		ReviewOnly:             status.BlockedGateCount > 0,
		ApprovalState:          "not_approved",
		ApprovalGranted:        false,
		ExternalWritesAllowed:  false,
		SelfApprovalAllowed:    false,
		NextSafeAction:         status.NextSafeAction,
	}
}

func workItemEvidenceRef(status Status) string {
	if status.NextSafeAction == "repair_local_runtime_health" {
		return status.LocalRuntime.EvidenceRef
	}
	return status.AuthorityReview.EvidenceRef
}

func workItemPublicSafe(status Status) bool {
	return status.PublicSafe &&
		status.AuthorityReview.PublicSafe &&
		!status.Authority.SelfAuthorityAllowed
}
