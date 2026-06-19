package knowledge

import (
	"errors"
	"strings"
	"time"
)

func Search(root string, query string) (SearchReport, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return SearchReport{}, errors.New("knowledge search query is required")
	}
	registry, err := ReadRegistry(root)
	if err != nil {
		return SearchReport{}, err
	}
	return searchRegistry(root, registry, query)
}

func searchRegistry(root string, registry Registry, query string) (SearchReport, error) {
	report := SearchReport{Query: query, CheckedAt: time.Now().UTC().Format(time.RFC3339)}
	matched := matchedConcepts(registry, query)
	matchedHarnesses := matchedHarnessCases(registry, query, matched)
	report.Concepts = conceptSummaries(matched)
	report.Events = eventSummaries(matchedEvents(registry, query, matched))
	report.HarnessCases = harnessSummaries(matchedHarnesses)
	report.DuplicateSuspicions = duplicateSuspicionsFor(registry, query)
	files, err := indexFiles(root, registry.KnowledgeIndexSchema.IndexRoots)
	if err != nil {
		return SearchReport{}, err
	}
	terms := termsForConcepts(matched)
	if len(terms) == 0 {
		terms = []termMatch{{Term: query}}
	}
	hits, issues, err := searchFiles(root, files, terms, 80)
	if err != nil {
		return SearchReport{}, err
	}
	report.Hits = hits
	report.LinearIssues = issues
	report.MustRead = mustReadFiles(root, matched, matchedHarnesses, hits)
	return report, nil
}
