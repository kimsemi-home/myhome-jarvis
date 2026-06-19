package knowledge

import (
	"fmt"
)

func validateDDDKindUse(registry Registry, usedKinds map[string]bool) []string {
	var failures []string
	for _, pattern := range registry.DDDPatterns {
		if !usedKinds[clean(pattern)] {
			failures = append(failures, fmt.Sprintf("ddd kind %q is not used by any concept", pattern))
		}
	}
	return failures
}

func validateRelatedConcepts(registry Registry, concepts map[string]bool) []string {
	var failures []string
	for _, concept := range registry.Concepts {
		for _, related := range concept.RelatedConcepts {
			if !concepts[related] {
				failures = append(failures, fmt.Sprintf(
					"concept %q references unknown related concept %q",
					concept.CanonicalName,
					related,
				))
			}
		}
	}
	return failures
}
