package learning

import "time"

func normalizeObservation(policy Policy, request RecordRequest) (Observation, error) {
	observation := Observation{
		At:           time.Now().UTC().Format(time.RFC3339),
		Kind:         normalizeToken(request.Kind),
		Source:       normalizeToken(request.Source),
		Stage:        normalizeToken(request.Stage),
		Status:       normalizeToken(request.Status),
		Summary:      publicText(request.Summary),
		EvidenceRefs: normalizeRefs(request.EvidenceRefs),
		Owner:        normalizeToken(request.Owner),
		NextAction:   publicText(request.NextAction),
	}
	applyObservationDefaults(&observation)
	if err := validateObservation(policy, observation); err != nil {
		return Observation{}, err
	}
	observation.ID = observationID(observation)
	return observation, nil
}

func applyObservationDefaults(observation *Observation) {
	if observation.Stage == "" {
		observation.Stage = "evidence_recorded"
	}
	if observation.Status == "" {
		observation.Status = "open"
	}
}
