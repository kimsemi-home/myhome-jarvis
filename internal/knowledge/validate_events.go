package knowledge

import "fmt"

func validateEvents(registry Registry, contexts map[string]bool, concepts map[string]bool) []string {
	var failures []string
	if len(registry.DomainEvents) == 0 {
		failures = append(failures, "domain events are required")
	}
	for _, event := range registry.DomainEvents {
		failures = append(failures, validateEvent(event, contexts, concepts)...)
	}
	return failures
}

func validateEvent(event DomainEvent, contexts map[string]bool, concepts map[string]bool) []string {
	var failures []string
	if clean(event.Name) == "" {
		failures = append(failures, "domain event name is required")
	}
	if !contexts[event.BoundedContext] {
		failures = append(failures, fmt.Sprintf("domain event %q references unknown bounded context %q", event.Name, event.BoundedContext))
	}
	if !concepts[event.EmittedBy] {
		failures = append(failures, fmt.Sprintf("domain event %q references unknown emitter concept %q", event.Name, event.EmittedBy))
	}
	if len(event.PayloadFields) == 0 {
		failures = append(failures, fmt.Sprintf("domain event %q must declare payload fields", event.Name))
	}
	return failures
}
