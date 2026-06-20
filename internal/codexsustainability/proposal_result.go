package codexsustainability

func resultForProposal(record Record, status Status) ProposalRecordResult {
	return ProposalRecordResult{
		ProposalID:            record.ProposalID,
		CostPerAcceptedChange: record.CostPerAcceptedChange,
		MedianCycleMinutes:    record.MedianCycleMinutes,
		CacheSavingsUnits:     record.CacheSavingsUnits,
		ReviewState:           reviewState(status),
		ReviewGateCount:       status.ReviewGateCount,
		EvidenceRefCount:      len(record.EvidenceRefs),
		RecordedAt:            record.At,
	}
}

func reviewState(status Status) string {
	if status.ReviewGateCount > 0 {
		return "review_required"
	}
	return "clear"
}
