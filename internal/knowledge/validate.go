package knowledge

func registryFailures(root string, registry Registry) []string {
	var failures []string
	contexts := map[string]bool{}
	failures = append(failures, validateContexts(registry, contexts)...)
	state := newConceptValidationState(registry)
	failures = append(failures, validateConcepts(root, registry, contexts, state)...)
	failures = append(failures, validateDDDKindUse(registry, state.usedKinds)...)
	failures = append(failures, validateRelatedConcepts(registry, state.concepts)...)
	failures = append(failures, validateEvents(registry, contexts, state.concepts)...)
	failures = append(failures, validateHarnesses(root, registry, contexts)...)
	failures = append(failures, validateArtifacts(root, registry)...)
	failures = append(failures, validateKnowledgeSchema(registry)...)
	return append(failures, validatePlanningRules(registry)...)
}

type conceptValidationState struct {
	concepts  map[string]bool
	aliases   map[string]string
	dddKinds  map[string]bool
	usedKinds map[string]bool
}

func newConceptValidationState(registry Registry) conceptValidationState {
	state := conceptValidationState{
		concepts:  map[string]bool{},
		aliases:   map[string]string{},
		dddKinds:  map[string]bool{},
		usedKinds: map[string]bool{},
	}
	for _, pattern := range registry.DDDPatterns {
		state.dddKinds[clean(pattern)] = true
	}
	return state
}
