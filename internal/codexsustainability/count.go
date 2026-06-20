package codexsustainability

import "errors"

func classifyRecordError(status *Status, record Record, err error) {
	if errors.Is(err, errMissingEvidenceRef) {
		status.MissingEvidenceCount++
		if normalizeToken(record.RecordKind) == "feature_proposal" {
			status.OptimizationClaimWithoutEvidenceCount++
		}
		return
	}
	status.InvalidRecordCount++
}

func applyRecord(status *Status, record Record) {
	status.RecordCount++
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, record.At)
	if record.RecordKind == "trend_baseline" {
		applyTrendBaseline(status, record)
		return
	}
	if record.RecordKind == "feature_proposal" {
		applyProposal(status, record)
		return
	}
	applyMetric(status, record)
}
