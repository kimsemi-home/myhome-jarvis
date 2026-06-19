package knowledge

import "sort"

func duplicateSuspicionsFor(registry Registry, query string) []DuplicateSuspicion {
	queryKey := normalizedTerm(query)
	concepts := map[string][]string{}
	seen := map[string]map[string]bool{}
	for _, concept := range registry.Concepts {
		addDuplicateTerms(concept, queryKey, concepts, seen)
	}
	var suspicions []DuplicateSuspicion
	for term, names := range concepts {
		if len(names) < 2 {
			continue
		}
		sort.Strings(names)
		suspicions = append(suspicions, DuplicateSuspicion{Term: term, Concepts: names})
	}
	sort.Slice(suspicions, func(i, j int) bool {
		return suspicions[i].Term < suspicions[j].Term
	})
	return suspicions
}

func addDuplicateTerms(
	concept Concept,
	queryKey string,
	concepts map[string][]string,
	seen map[string]map[string]bool,
) {
	for _, term := range conceptTerms(concept) {
		key := normalizedTerm(term)
		if key != queryKey && !stringsContainsEither(key, queryKey) {
			continue
		}
		if seen[key] == nil {
			seen[key] = map[string]bool{}
		}
		if !seen[key][concept.CanonicalName] {
			seen[key][concept.CanonicalName] = true
			concepts[key] = append(concepts[key], concept.CanonicalName)
		}
	}
}
