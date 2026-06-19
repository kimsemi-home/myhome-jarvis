package learning

import "fmt"

func validateObservation(policy Policy, observation Observation) error {
	if !contains(normalizeList(policy.AllowedKinds), observation.Kind) {
		return fmt.Errorf("learning kind %q is not allowed", observation.Kind)
	}
	if !contains(normalizeList(policy.Lifecycle), observation.Stage) {
		return fmt.Errorf("learning lifecycle stage %q is not allowed", observation.Stage)
	}
	if !contains(normalizeList(policy.AllowedStatuses), observation.Status) {
		return fmt.Errorf("learning status %q is not allowed", observation.Status)
	}
	if observation.Source == "" || observation.Summary == "" ||
		observation.Owner == "" || observation.NextAction == "" {
		return fmt.Errorf("learning record requires source, summary, owner, and next_action")
	}
	if len(observation.EvidenceRefs) == 0 {
		return fmt.Errorf("learning record requires at least one evidence ref")
	}
	for _, ref := range observation.EvidenceRefs {
		if err := validateEvidenceRef(policy, ref); err != nil {
			return err
		}
	}
	return validateObservationText(observation)
}

func validateObservationText(observation Observation) error {
	for _, value := range []string{
		observation.Source,
		observation.Summary,
		observation.Owner,
		observation.NextAction,
	} {
		if err := rejectSensitiveText(value); err != nil {
			return err
		}
	}
	return nil
}
