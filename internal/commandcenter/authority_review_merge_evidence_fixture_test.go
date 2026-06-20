package commandcenter

func authorityReviewMergeEvidenceFixture() MergeEvidenceSummary {
	return MergeEvidenceSummary{
		PublicSafe:                   true,
		DefaultBehavior:              "merge_when_eligible",
		MergePreference:              "merge_after_checks_pass",
		PostMergeEvidenceRequired:    true,
		LinearCompletionRequired:     true,
		MainQualityRunRequired:       true,
		PrivateDataScanRequired:      true,
		EligibleGateCount:            5,
		RequiredEvidenceCount:        11,
		MissingGateCount:             0,
		MissingRequiredEvidenceCount: 0,
		MergeReady:                   true,
		MergeBlockedUntilEvidence:    false,
	}
}
