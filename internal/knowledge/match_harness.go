package knowledge

import "sort"

func matchedHarnessCases(registry Registry, query string, concepts []Concept) []HarnessCase {
	queryKey := normalizedTerm(query)
	contexts := map[string]bool{}
	for _, concept := range concepts {
		contexts[concept.BoundedContext] = true
	}
	var matched []HarnessCase
	for _, harness := range registry.HarnessCaseContracts {
		if contexts[harness.BoundedContext] ||
			matchesAny(queryKey, harness.Name, harness.BoundedContext, harness.Command, harness.EvidenceTarget, harness.Description) {
			matched = append(matched, harness)
		}
	}
	return matched
}

func harnessSummaries(harnesses []HarnessCase) []HarnessCaseSummary {
	summaries := make([]HarnessCaseSummary, 0, len(harnesses))
	for _, harness := range harnesses {
		summaries = append(summaries, HarnessCaseSummary{
			Name:           harness.Name,
			BoundedContext: harness.BoundedContext,
			Command:        harness.Command,
			EvidenceTarget: harness.EvidenceTarget,
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Name < summaries[j].Name
	})
	return summaries
}
