package monetization

func resultForRecord(record Record, status Status) RecordResult {
	return RecordResult{
		ExperimentID:          record.ExperimentID,
		HypothesisKey:         record.HypothesisKey,
		State:                 record.State,
		DecisionKind:          record.DecisionKind,
		ReviewStatus:          record.ReviewStatus,
		ExpectedValueBand:     record.ExpectedValueBand,
		CostEstimateUnits:     record.CostEstimateUnits,
		CostUnitKind:          record.CostUnitKind,
		EvidenceRefCount:      len(record.EvidenceRefs),
		MonetizationDebtCount: status.MonetizationDebtCount,
		RecordedAt:            record.At,
	}
}
