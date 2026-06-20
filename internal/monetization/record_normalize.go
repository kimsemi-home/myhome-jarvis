package monetization

import "time"

func normalizeRecordRequest(
	policy Policy,
	request RecordRequest,
	now time.Time,
) (Record, error) {
	recordedAt, err := normalizeRecordedAt(request.At, now)
	if err != nil {
		return Record{}, err
	}
	record := Record{
		At:                recordedAt,
		ExperimentID:      request.ExperimentID,
		HypothesisKey:     request.HypothesisKey,
		State:             request.State,
		DecisionKind:      request.DecisionKind,
		ReviewStatus:      request.ReviewStatus,
		ExpectedValueBand: request.ExpectedValueBand,
		CostEstimateUnits: request.CostEstimateUnits,
		CostUnitKind:      request.CostUnitKind,
		EvidenceRefs:      request.EvidenceRefs,
	}
	return normalizeRecord(policy, record)
}
