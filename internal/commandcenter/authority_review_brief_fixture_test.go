package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewBriefPlan() authority.ReviewPlanStatus {
	return authority.ReviewPlanStatus{
		PolicyPath: "generated/authority.generated.json",
		PublicSafe: true,
		RequiredReviewClasses: []string{
			"high_risk_public_repo_review", "human_review",
			"public_repo_change_review", "public_safety_review",
			"workflow_change_review",
		},
	}
}

func authorityReviewBriefStatus(policy visionPolicy) Status {
	status := visionAuditFixtureStatus(policy)
	status.AuthorityReview = AuthorityReviewSummary{
		RequestID:                     "authority-review-f8e5c9db088a",
		RequestState:                  "ready",
		EvidenceRef:                   "authority_review_request:authority-review-f8e5c9db088a",
		EvidenceReady:                 true,
		QueueState:                    "pending_human_review",
		QueueReady:                    true,
		PublicSafe:                    true,
		ReviewRequestRecorded:         true,
		ReviewRequestLedgerState:      "recorded_pending_review",
		ReviewRequestStaleAfterHours:  24,
		ReviewRequestEscalationAction: "none",
		PendingReviewClassCount:       5,
		ReviewRequestApprovalState:    "not_approved",
		RequiredReviewClassCount:      5,
		HighRiskBlockedDecisionCount:  6,
	}
	status.RepoFactory.AuthorityReviewRequired = true
	status.RepoFactory.PublicSafetyEvidenceRequired = true
	status.RepoFactoryPreflight = RepoFactoryPreflightSummary{
		PublicSafe:                     true,
		CreationDecision:               "blocked_pending_review_evidence",
		CreationAllowed:                false,
		RepoCreationBlockedUntilReview: true,
		SelfApprovalAllowed:            false,
		TemplateReadyCount:             6,
		TemplateFileCount:              6,
		GateReadyCount:                 4,
		CreationGateCount:              5,
		BlockingGateCount:              1,
		MissingEvidenceKeys:            []string{"authority_review"},
		NextSafeAction:                 "await_human_authority_review",
	}
	status.LocalRuntime = authorityReviewHealthyRuntimeFixture()
	status.MergeEvidence = authorityReviewMergeEvidenceFixture()
	status.CodexSustainability = authorityReviewCodexSustainabilityFixture()
	status.WorkItem = summarizeWorkItem(status)
	return status
}

func authorityReviewHealthyRuntimeFixture() LocalRuntimeSummary {
	return LocalRuntimeSummary{
		PublicSafe:      true,
		EvidenceRef:     "local_runtime:supervisor",
		State:           "healthy",
		Recorded:        true,
		ProcessRunning:  true,
		ProbeOK:         true,
		NextSafeAction:  "none",
		Message:         "daemon is reachable",
		HealthDebtCount: 0,
	}
}
