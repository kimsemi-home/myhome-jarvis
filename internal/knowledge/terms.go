package knowledge

import "strings"

type termMatch struct {
	Concept string
	Term    string
}

func termsForConcepts(concepts []Concept) []termMatch {
	var terms []termMatch
	for _, concept := range concepts {
		for _, term := range conceptTerms(concept) {
			terms = append(terms, termMatch{Concept: concept.CanonicalName, Term: term})
		}
	}
	return terms
}

func conceptTerms(concept Concept) []string {
	terms := []string{concept.CanonicalName, concept.BoundedContext, concept.DDDKind}
	terms = append(terms, concept.AllowedAliases...)
	return append(terms, concept.RelatedConcepts...)
}

func matchesAny(queryKey string, values ...string) bool {
	for _, value := range values {
		key := normalizedTerm(value)
		if key != "" && (strings.Contains(key, queryKey) || strings.Contains(queryKey, key)) {
			return true
		}
	}
	return false
}

func matchLine(lowerLine string, term string) bool {
	term = strings.TrimSpace(strings.ToLower(term))
	if term == "" {
		return false
	}
	return strings.Contains(lowerLine, term) ||
		strings.Contains(normalizedTerm(lowerLine), normalizedTerm(term))
}

func normalizedTerm(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	replacer := strings.NewReplacer(" ", "", "-", "", "_", "", "/", "", ".", "")
	return replacer.Replace(value)
}

func clean(value string) string {
	return strings.TrimSpace(value)
}
