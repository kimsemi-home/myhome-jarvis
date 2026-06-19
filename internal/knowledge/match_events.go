package knowledge

import "sort"

func matchedEvents(registry Registry, query string, concepts []Concept) []DomainEvent {
	queryKey := normalizedTerm(query)
	conceptNames := map[string]bool{}
	for _, concept := range concepts {
		conceptNames[concept.CanonicalName] = true
	}
	includeAll := queryKey == normalizedTerm("DomainEvent") || queryKey == normalizedTerm("event")
	var matched []DomainEvent
	for _, event := range registry.DomainEvents {
		if includeAll || eventMatches(queryKey, conceptNames, event) {
			matched = append(matched, event)
		}
	}
	return matched
}

func eventMatches(queryKey string, conceptNames map[string]bool, event DomainEvent) bool {
	return conceptNames[event.Name] ||
		conceptNames[event.EmittedBy] ||
		matchesAny(queryKey, event.Name, event.BoundedContext, event.EmittedBy, event.Description)
}

func eventSummaries(events []DomainEvent) []DomainEventSummary {
	summaries := make([]DomainEventSummary, 0, len(events))
	for _, event := range events {
		summaries = append(summaries, DomainEventSummary{
			Name:           event.Name,
			BoundedContext: event.BoundedContext,
			EmittedBy:      event.EmittedBy,
			PayloadFields:  append([]string(nil), event.PayloadFields...),
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Name < summaries[j].Name
	})
	return summaries
}
