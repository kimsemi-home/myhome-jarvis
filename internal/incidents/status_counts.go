package incidents

import (
	"errors"
	"time"
)

func newStatus(policy Policy, checkedAt time.Time) Status {
	return Status{
		PolicyPath:                PolicyRelativePath,
		LedgerPath:                policy.PrivateIncidentLedger,
		QuarantineStaleAfterHours: policy.QuarantineStaleAfterHours,
		ByKind:                    map[string]int{},
		ByStage:                   map[string]int{},
		ByStatus:                  map[string]int{},
		ByOwnerRole:               map[string]int{},
		ByQuarantineState:         map[string]int{},
		CheckedAt:                 checkedAt.Format(time.RFC3339),
	}
}

func addIncidentStatus(policy Policy, checkedAt time.Time, status *Status, incident Incident) {
	status.Count++
	status.ByKind[incident.Kind]++
	status.ByStage[incident.Stage]++
	status.ByStatus[incident.Status]++
	status.ByOwnerRole[incident.OwnerRole]++
	status.ByQuarantineState[incident.QuarantineState]++
	if incident.Status == "closed" {
		status.ClosedCount++
	} else {
		status.OpenCount++
	}
	if isStaleQuarantine(policy, incident, checkedAt) {
		status.StaleQuarantineCount++
	}
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, incident.At)
}

func recordIncidentError(status *Status, err error) {
	switch {
	case errors.Is(err, errMissingOwner):
		status.MissingOwnerCount++
	case errors.Is(err, errMissingEvidence):
		status.MissingEvidenceRefCount++
	default:
		status.InvalidIncidentCount++
	}
}

func incidentDebt(status Status) int {
	return status.InvalidIncidentCount +
		status.MissingOwnerCount +
		status.MissingEvidenceRefCount +
		status.StaleQuarantineCount
}
