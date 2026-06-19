package knowledge

import "fmt"

func validateConcepts(root string, registry Registry, contexts map[string]bool, state conceptValidationState) []string {
	var failures []string
	for _, concept := range registry.Concepts {
		name := clean(concept.CanonicalName)
		if name == "" {
			failures = append(failures, "concept canonical_name is required")
			continue
		}
		failures = append(failures, validateConceptCore(concept, name, contexts, state)...)
		failures = append(failures, validateConceptTargets(root, concept, name)...)
		failures = append(failures, validateConceptAliases(concept, name, state)...)
	}
	return failures
}

func validateConceptCore(concept Concept, name string, contexts map[string]bool, state conceptValidationState) []string {
	var failures []string
	if state.concepts[name] {
		failures = append(failures, fmt.Sprintf("duplicate concept %q", name))
	}
	state.concepts[name] = true
	if !contexts[concept.BoundedContext] {
		failures = append(failures, fmt.Sprintf("concept %q references unknown bounded context %q", name, concept.BoundedContext))
	}
	kind := clean(concept.DDDKind)
	if kind == "" || !state.dddKinds[kind] {
		failures = append(failures, fmt.Sprintf("concept %q references invalid ddd_kind %q", name, concept.DDDKind))
	} else {
		state.usedKinds[kind] = true
	}
	return failures
}

func validateConceptTargets(root string, concept Concept, name string) []string {
	var failures []string
	if len(concept.GeneratedTargets) == 0 {
		failures = append(failures, fmt.Sprintf("concept %q must declare generated targets", name))
	}
	for _, target := range concept.GeneratedTargets {
		if err := requirePublicTarget(root, target); err != nil {
			failures = append(failures, fmt.Sprintf("concept %q target %q: %v", name, target, err))
		}
	}
	return failures
}

func validateConceptAliases(concept Concept, name string, state conceptValidationState) []string {
	var failures []string
	for _, alias := range concept.AllowedAliases {
		key := normalizedTerm(alias)
		if key == "" {
			failures = append(failures, fmt.Sprintf("concept %q has empty alias", name))
			continue
		}
		if owner, exists := state.aliases[key]; exists && owner != name {
			failures = append(failures, fmt.Sprintf("alias %q is shared by %q and %q", alias, owner, name))
		}
		state.aliases[key] = name
	}
	return failures
}
