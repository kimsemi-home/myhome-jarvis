package incidents

import (
	"fmt"
	"time"
)

func normalizeIncident(policy Policy, incident Incident) (Incident, error) {
	normalized := Incident{
		ID:              publicText(incident.ID),
		At:              publicText(incident.At),
		Kind:            normalizeToken(incident.Kind),
		Stage:           normalizeToken(incident.Stage),
		Status:          normalizeToken(incident.Status),
		OwnerRole:       normalizeToken(incident.OwnerRole),
		QuarantineState: normalizeToken(incident.QuarantineState),
		EvidenceRefs:    normalizeRefs(incident.EvidenceRefs),
	}
	if normalized.QuarantineState == "" {
		normalized.QuarantineState = "none"
	}
	if normalized.At == "" {
		return Incident{}, fmt.Errorf("incident at timestamp is required")
	}
	if _, err := time.Parse(time.RFC3339, normalized.At); err != nil {
		return Incident{}, fmt.Errorf("incident at timestamp is invalid")
	}
	if err := validateIncidentTokens(policy, normalized); err != nil {
		return Incident{}, err
	}
	if len(normalized.EvidenceRefs) == 0 {
		return Incident{}, errMissingEvidence
	}
	for _, ref := range normalized.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return Incident{}, err
		}
		if err := rejectSensitiveText(ref); err != nil {
			return Incident{}, err
		}
	}
	return normalized, nil
}

func validateIncidentTokens(policy Policy, incident Incident) error {
	if !contains(normalizeList(policy.AllowedKinds), incident.Kind) {
		return fmt.Errorf("incident kind %q is not allowed", incident.Kind)
	}
	if !contains(normalizeList(policy.Lifecycle), incident.Stage) {
		return fmt.Errorf("incident stage %q is not allowed", incident.Stage)
	}
	if !contains(normalizeList(policy.AllowedStatuses), incident.Status) {
		return fmt.Errorf("incident status %q is not allowed", incident.Status)
	}
	if incident.OwnerRole == "" {
		return errMissingOwner
	}
	if !contains(normalizeList(policy.OwnerRoles), incident.OwnerRole) {
		return fmt.Errorf("incident owner role %q is not allowed", incident.OwnerRole)
	}
	if !contains(normalizeList(policy.QuarantineStates), incident.QuarantineState) {
		return fmt.Errorf("incident quarantine state %q is not allowed", incident.QuarantineState)
	}
	return nil
}
