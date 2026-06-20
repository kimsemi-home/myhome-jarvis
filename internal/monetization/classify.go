package monetization

import "errors"

func classifyRecordError(status *Status, err error) {
	switch {
	case errors.Is(err, errMissingEvidenceRef):
		status.MissingEvidenceCount++
	case errors.Is(err, errMissingCostEstimate):
		status.MissingCostEstimateCount++
	default:
		status.InvalidRecordCount++
	}
}

func debtCount(status Status) int {
	return status.InvalidRecordCount +
		status.MissingEvidenceCount +
		status.MissingCostEstimateCount +
		status.ReviewRequiredCount +
		status.ExpectedValueUnknownCount
}
