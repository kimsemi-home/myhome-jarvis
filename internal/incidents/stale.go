package incidents

import "time"

func isStaleQuarantine(policy Policy, incident Incident, checkedAt time.Time) bool {
	if incident.QuarantineState != "quarantined" && incident.Status != "quarantined" {
		return false
	}
	at, err := time.Parse(time.RFC3339, incident.At)
	if err != nil {
		return false
	}
	return checkedAt.Sub(at) > time.Duration(policy.QuarantineStaleAfterHours)*time.Hour
}
