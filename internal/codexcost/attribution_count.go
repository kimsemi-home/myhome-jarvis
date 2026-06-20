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
	status.TotalUnits += record.Amount
	status.ByScope[record.Scope] += record.Amount
	subjects := status.subjectsByScope[record.Scope]
	if subjects == nil {
		subjects = map[string]bool{}
		status.subjectsByScope[record.Scope] = subjects
	}
	subjects[record.SubjectHash] = true
}

func finalizeAttributionStatus(status AttributionStatus) AttributionStatus {
	for scope, subjects := range status.subjectsByScope {
		status.SubjectCountByScope[scope] = len(subjects)
	}
	status.subjectsByScope = nil
	return status
}
