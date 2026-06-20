package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/mergeevidence"

func summarizeMergeEvidence(status mergeevidence.Status) MergeEvidenceSummary {
	return MergeEvidenceSummary{
		PublicSafe:                   status.PublicSafe,
		DefaultBehavior:              status.DefaultBehavior,
		MergePreference:              status.MergePreference,
		PostMergeEvidenceRequired:    status.PostMergeEvidenceRequired,
		LinearCompletionRequired:     status.LinearCompletionRequired,
		MainQualityRunRequired:       status.MainQualityRunRequired,
		PrivateDataScanRequired:      status.PrivateDataScanRequired,
		EligibleGateCount:            status.EligibleGateCount,
		RequiredEvidenceCount:        status.RequiredEvidenceCount,
		MissingGateCount:             status.MissingGateCount,
		MissingRequiredEvidenceCount: status.MissingRequiredEvidenceCount,
		MergeReady:                   status.MergeReady,
		MergeBlockedUntilEvidence:    status.MergeBlockedUntilEvidence,
	}
}
