package knowledge

import (
	"sort"
	"strings"
)

func matchedConcepts(registry Registry, query string) []Concept {
	queryKey := normalizedTerm(query)
	var matched []Concept
	for _, concept := range registry.Concepts {
		for _, term := range conceptTerms(concept) {
			key := normalizedTerm(term)
			if strings.Contains(key, queryKey) || strings.Contains(queryKey, key) {
				matched = append(matched, concept)
				break
			}
		}
	}
	if len(matched) > 0 {
		return matched
	}
	for _, concept := range registry.Concepts {
		if strings.Contains(strings.ToLower(concept.Description), strings.ToLower(query)) {
			matched = append(matched, concept)
		}
	}
	return matched
}

func conceptSummaries(concepts []Concept) []ConceptSummary {
	summaries := make([]ConceptSummary, 0, len(concepts))
	for _, concept := range concepts {
		summaries = append(summaries, ConceptSummary{
			CanonicalName:    concept.CanonicalName,
			BoundedContext:   concept.BoundedContext,
			DDDKind:          concept.DDDKind,
			Owner:            concept.Owner,
			Definition:       RegistryRelativePath,
			AllowedAliases:   append([]string(nil), concept.AllowedAliases...),
			GeneratedTargets: append([]string(nil), concept.GeneratedTargets...),
			RelatedConcepts:  append([]string(nil), concept.RelatedConcepts...),
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].CanonicalName < summaries[j].CanonicalName
	})
	return summaries
}
