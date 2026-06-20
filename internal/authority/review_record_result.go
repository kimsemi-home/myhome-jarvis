package authority

func resultForReviewRecord(record ReviewRecord) ReviewRecordResult {
	return ReviewRecordResult{
		RequestID:                record.RequestID,
		LinearIssueRef:           record.LinearIssueRef,
		RequestState:             record.RequestState,
		QueueState:               record.QueueState,
		LedgerState:              "recorded_private",
		ApprovalState:            record.ApprovalState,
		ApprovalGranted:          record.ApprovalGranted,
		ExternalWritesAllowed:    record.ExternalWritesAllowed,
		SelfApprovalAllowed:      record.SelfApprovalAllowed,
		RequiredReviewClassCount: record.RequiredReviewClassCount,
		RecordedAt:               record.At,
		PublicSafe:               record.PublicSafe,
	}
}
