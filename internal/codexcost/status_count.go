package codexcost

import "errors"

func classifyRecordError(status *Status, err error) {
	if errors.Is(err, errMissingEvidenceRef) {
		status.MissingEvidenceCount++
	} else {
		status.InvalidRecordCount++
	}
}

func applyRecord(status *Status, record Record) {
	status.RecordCount++
	status.TotalUnits += record.Amount
	status.ByUnitKind[record.UnitKind] += record.Amount
	status.ByScope[record.Scope] += record.Amount
	status.ByStatus[record.Status]++
	if record.Status == "review_required" {
		status.ReviewRequiredCount++
	}
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, record.At)
}

func budgetState(policy Policy, total int64) string {
	if total >= policy.ReviewUnitThreshold {
		return "review_required"
	}
	if total >= policy.WarningUnitThreshold {
		return "warning"
	}
	return "ok"
}
