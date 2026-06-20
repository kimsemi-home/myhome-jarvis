package codexcost

import "errors"

func classifyAttributionError(status *AttributionStatus, err error) {
	if errors.Is(err, errMissingEvidenceRef) {
		status.MissingEvidenceCount++
		return
	}
	status.InvalidRecordCount++
}

func applyAttribution(status *AttributionStatus, record AttributionRecord) {
	status.RecordCount++
	status.EntryUnits += record.Amount
	status.ByScope[record.Scope] += record.Amount
	subjects := status.subjectsByScope[record.Scope]
	if subjects == nil {
		subjects = map[string]bool{}
		status.subjectsByScope[record.Scope] = subjects
	}
	subjects[record.SubjectHash] = true
	if record.Amount > status.costRefUnits[record.CostRef] {
		status.costRefUnits[record.CostRef] = record.Amount
	}
}

func finalizeAttributionStatus(status AttributionStatus) AttributionStatus {
	for scope, subjects := range status.subjectsByScope {
		status.SubjectCountByScope[scope] = len(subjects)
	}
	for _, amount := range status.costRefUnits {
		status.CoverageUnits += amount
	}
	status.TotalUnits = status.CoverageUnits
	status.CostRefCount = len(status.costRefUnits)
	status.subjectsByScope = nil
	status.costRefUnits = nil
	return status
}
