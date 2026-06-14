package knowledge

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const RegistryRelativePath = "generated/concepts.generated.json"

type Registry struct {
	BoundedContexts            []BoundedContext   `json:"bounded_contexts"`
	DDDPatterns                []string           `json:"ddd_patterns"`
	Concepts                   []Concept          `json:"concepts"`
	GeneratedArtifactContracts []ArtifactContract `json:"generated_artifact_contracts"`
	PlanningRules              PlanningRules      `json:"planning_rules"`
	KnowledgeIndexSchema       IndexSchema        `json:"knowledge_index_schema"`
}

type BoundedContext struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}

type Concept struct {
	CanonicalName    string   `json:"canonical_name"`
	BoundedContext   string   `json:"bounded_context"`
	Description      string   `json:"description"`
	AllowedAliases   []string `json:"allowed_aliases"`
	Owner            string   `json:"owner"`
	GeneratedTargets []string `json:"generated_targets"`
	RelatedConcepts  []string `json:"related_concepts"`
}

type ArtifactContract struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Owner string `json:"owner"`
}

type PlanningRules struct {
	KnowledgeIndexRequiredBeforePlanning bool     `json:"knowledge_index_required_before_planning"`
	DefaultKnowledgeQuery                string   `json:"default_knowledge_query"`
	SemanticChangesRequireSSOTFirst      bool     `json:"semantic_changes_require_ssot_first"`
	SSOTChangeRequiresCodegen            bool     `json:"ssot_change_requires_codegen"`
	SmallCohesiveChangeRequired          bool     `json:"small_cohesive_change_required"`
	ValidationSteps                      []string `json:"validation_steps"`
}

type IndexSchema struct {
	Kind                    string   `json:"kind"`
	ExternalVectorDBAllowed bool     `json:"external_vector_db_allowed"`
	CloudRAGAllowed         bool     `json:"cloud_rag_allowed"`
	IndexRoots              []string `json:"index_roots"`
	QueryTypes              []string `json:"query_types"`
	EvidenceFields          []string `json:"evidence_fields"`
}

type Check struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type VerifyReport struct {
	OK           bool    `json:"ok"`
	CheckedAt    string  `json:"checked_at"`
	ContextCount int     `json:"context_count"`
	ConceptCount int     `json:"concept_count"`
	Checks       []Check `json:"checks"`
}

type SearchReport struct {
	Query               string               `json:"query"`
	CheckedAt           string               `json:"checked_at"`
	Concepts            []ConceptSummary     `json:"concepts"`
	Hits                []Hit                `json:"hits"`
	LinearIssues        []string             `json:"linear_issues,omitempty"`
	DuplicateSuspicions []DuplicateSuspicion `json:"duplicate_suspicions"`
	MustRead            []string             `json:"must_read"`
}

type ConceptSummary struct {
	CanonicalName    string   `json:"canonical_name"`
	BoundedContext   string   `json:"bounded_context"`
	Owner            string   `json:"owner"`
	Definition       string   `json:"definition"`
	AllowedAliases   []string `json:"allowed_aliases"`
	GeneratedTargets []string `json:"generated_targets"`
	RelatedConcepts  []string `json:"related_concepts"`
}

type Hit struct {
	Path    string `json:"path"`
	Line    int    `json:"line"`
	Concept string `json:"concept,omitempty"`
	Term    string `json:"term"`
}

type DuplicateSuspicion struct {
	Term     string   `json:"term"`
	Concepts []string `json:"concepts"`
}

type Evidence struct {
	Query        string   `json:"query"`
	ConceptCount int      `json:"concept_count"`
	HitCount     int      `json:"hit_count"`
	LinearIssues []string `json:"linear_issues,omitempty"`
	MustRead     []string `json:"must_read,omitempty"`
	CheckedAt    string   `json:"checked_at"`
}

func SummarizeSearch(report SearchReport) Evidence {
	return Evidence{
		Query:        report.Query,
		ConceptCount: len(report.Concepts),
		HitCount:     len(report.Hits),
		LinearIssues: append([]string(nil), report.LinearIssues...),
		MustRead:     append([]string(nil), report.MustRead...),
		CheckedAt:    report.CheckedAt,
	}
}

func ReadRegistry(root string) (Registry, error) {
	path := filepath.Join(root, filepath.FromSlash(RegistryRelativePath))
	file, err := os.Open(path)
	if err != nil {
		return Registry{}, err
	}
	defer file.Close()

	var registry Registry
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&registry); err != nil {
		return Registry{}, err
	}
	if failures := registryFailures(root, registry); len(failures) > 0 {
		return Registry{}, errors.New(strings.Join(failures, "; "))
	}
	return registry, nil
}

func Verify(root string) (VerifyReport, error) {
	report := VerifyReport{OK: true, CheckedAt: time.Now().UTC().Format(time.RFC3339)}
	registry, err := readRegistryUnchecked(root)
	if err != nil {
		report.OK = false
		report.Checks = append(report.Checks, Check{Name: "concept artifact", Status: "fail", Message: err.Error()})
		return report, nil
	}
	report.ContextCount = len(registry.BoundedContexts)
	report.ConceptCount = len(registry.Concepts)
	failures := registryFailures(root, registry)
	if len(failures) == 0 {
		report.Checks = append(report.Checks,
			Check{Name: "bounded contexts", Status: "pass"},
			Check{Name: "duplicate concepts", Status: "pass"},
			Check{Name: "registered domain terms", Status: "pass"},
			Check{Name: "alias drift", Status: "pass"},
			Check{Name: "generated artifact contracts", Status: "pass"},
			Check{Name: "knowledge index schema", Status: "pass"},
		)
		return report, nil
	}
	report.OK = false
	for _, failure := range failures {
		report.Checks = append(report.Checks, Check{Name: "ddd verify", Status: "fail", Message: failure})
	}
	return report, nil
}

func Search(root string, query string) (SearchReport, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return SearchReport{}, errors.New("knowledge search query is required")
	}
	registry, err := ReadRegistry(root)
	if err != nil {
		return SearchReport{}, err
	}
	report := SearchReport{
		Query:     query,
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}
	matched := matchedConcepts(registry, query)
	report.Concepts = conceptSummaries(matched)
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
	report.MustRead = mustReadFiles(root, matched, hits)
	return report, nil
}

func readRegistryUnchecked(root string) (Registry, error) {
	path := filepath.Join(root, filepath.FromSlash(RegistryRelativePath))
	file, err := os.Open(path)
	if err != nil {
		return Registry{}, err
	}
	defer file.Close()

	var registry Registry
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&registry); err != nil {
		return Registry{}, err
	}
	return registry, nil
}

func registryFailures(root string, registry Registry) []string {
	var failures []string
	contexts := map[string]bool{}
	for _, context := range registry.BoundedContexts {
		name := strings.TrimSpace(context.Name)
		if name == "" {
			failures = append(failures, "bounded context name is required")
			continue
		}
		if contexts[name] {
			failures = append(failures, fmt.Sprintf("duplicate bounded context %q", name))
		}
		contexts[name] = true
	}
	concepts := map[string]bool{}
	aliases := map[string]string{}
	for _, concept := range registry.Concepts {
		name := strings.TrimSpace(concept.CanonicalName)
		if name == "" {
			failures = append(failures, "concept canonical_name is required")
			continue
		}
		if concepts[name] {
			failures = append(failures, fmt.Sprintf("duplicate concept %q", name))
		}
		concepts[name] = true
		if !contexts[concept.BoundedContext] {
			failures = append(failures, fmt.Sprintf("concept %q references unknown bounded context %q", name, concept.BoundedContext))
		}
		if len(concept.GeneratedTargets) == 0 {
			failures = append(failures, fmt.Sprintf("concept %q must declare generated targets", name))
		}
		for _, target := range concept.GeneratedTargets {
			if err := requirePublicTarget(root, target); err != nil {
				failures = append(failures, fmt.Sprintf("concept %q target %q: %v", name, target, err))
			}
		}
		for _, alias := range concept.AllowedAliases {
			key := normalizedTerm(alias)
			if key == "" {
				failures = append(failures, fmt.Sprintf("concept %q has empty alias", name))
				continue
			}
			if owner, exists := aliases[key]; exists && owner != name {
				failures = append(failures, fmt.Sprintf("alias %q is shared by %q and %q", alias, owner, name))
			}
			aliases[key] = name
		}
	}
	for _, concept := range registry.Concepts {
		for _, related := range concept.RelatedConcepts {
			if !concepts[related] {
				failures = append(failures, fmt.Sprintf("concept %q references unknown related concept %q", concept.CanonicalName, related))
			}
		}
	}
	for _, contract := range registry.GeneratedArtifactContracts {
		if err := requirePublicTarget(root, contract.Path); err != nil {
			failures = append(failures, fmt.Sprintf("artifact contract %q target %q: %v", contract.Name, contract.Path, err))
		}
	}
	if registry.KnowledgeIndexSchema.Kind != "local-lexical" {
		failures = append(failures, "knowledge index must be local-lexical")
	}
	if registry.KnowledgeIndexSchema.CloudRAGAllowed || registry.KnowledgeIndexSchema.ExternalVectorDBAllowed {
		failures = append(failures, "knowledge index must not allow cloud RAG or external vector DB")
	}
	if len(registry.KnowledgeIndexSchema.IndexRoots) == 0 {
		failures = append(failures, "knowledge index roots are required")
	}
	for _, rootPath := range registry.KnowledgeIndexSchema.IndexRoots {
		if strings.HasPrefix(rootPath, "/") || strings.Contains(rootPath, "..") {
			failures = append(failures, fmt.Sprintf("knowledge index root %q must be repo-relative", rootPath))
		}
	}
	if !registry.PlanningRules.KnowledgeIndexRequiredBeforePlanning {
		failures = append(failures, "planning rules must require KnowledgeIndex before planning")
	}
	return failures
}

func requirePublicTarget(root string, rel string) error {
	rel = strings.TrimSpace(rel)
	if rel == "" {
		return errors.New("path is required")
	}
	if filepath.IsAbs(filepath.FromSlash(rel)) || strings.Contains(rel, "..") {
		return errors.New("path must be repo-relative")
	}
	if strings.HasPrefix(rel, "data/private/") {
		return errors.New("generated target must not be private")
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(rel))); err != nil {
		return err
	}
	return nil
}

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
	terms := []string{concept.CanonicalName, concept.BoundedContext}
	terms = append(terms, concept.AllowedAliases...)
	terms = append(terms, concept.RelatedConcepts...)
	return terms
}

func indexFiles(root string, roots []string) ([]string, error) {
	seen := map[string]bool{}
	var files []string
	for _, relRoot := range roots {
		relRoot = strings.TrimSpace(relRoot)
		if relRoot == "" || strings.Contains(relRoot, "..") || filepath.IsAbs(filepath.FromSlash(relRoot)) {
			return nil, fmt.Errorf("invalid knowledge index root %q", relRoot)
		}
		path := filepath.Join(root, filepath.FromSlash(relRoot))
		info, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) && strings.HasPrefix(relRoot, "data/private/") {
				continue
			}
			return nil, err
		}
		if !info.IsDir() {
			if indexableFile(path) {
				files = appendUnique(files, seen, path)
			}
			continue
		}
		err = filepath.WalkDir(path, func(path string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				switch entry.Name() {
				case ".git", ".dart_tool", "build", "dist", "target", "bin":
					return filepath.SkipDir
				}
				return nil
			}
			if indexableFile(path) {
				files = appendUnique(files, seen, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(files)
	return files, nil
}

func indexableFile(path string) bool {
	switch filepath.Ext(path) {
	case ".go", ".lisp", ".dart", ".md", ".json", ".jsonl", ".toml", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func searchFiles(root string, files []string, terms []termMatch, limit int) ([]Hit, []string, error) {
	var hits []Hit
	issueIDs := map[string]bool{}
	issuePattern := regexp.MustCompile(`\bKIM-[0-9]+\b`)
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, nil, err
		}
		scanner := bufio.NewScanner(file)
		lineNumber := 0
		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			lower := strings.ToLower(line)
			for _, match := range terms {
				if matchLine(lower, match.Term) {
					rel, err := filepath.Rel(root, path)
					if err != nil {
						file.Close()
						return nil, nil, err
					}
					if len(hits) < limit {
						hits = append(hits, Hit{
							Path:    filepath.ToSlash(rel),
							Line:    lineNumber,
							Concept: match.Concept,
							Term:    match.Term,
						})
					}
					for _, issue := range issuePattern.FindAllString(line, -1) {
						issueIDs[issue] = true
					}
					break
				}
			}
		}
		if err := scanner.Err(); err != nil {
			file.Close()
			return nil, nil, err
		}
		file.Close()
	}
	issues := make([]string, 0, len(issueIDs))
	for issue := range issueIDs {
		issues = append(issues, issue)
	}
	sort.Strings(issues)
	return hits, issues, nil
}

func matchLine(lowerLine string, term string) bool {
	term = strings.TrimSpace(strings.ToLower(term))
	if term == "" {
		return false
	}
	return strings.Contains(lowerLine, term) || strings.Contains(normalizedTerm(lowerLine), normalizedTerm(term))
}

func duplicateSuspicionsFor(registry Registry, query string) []DuplicateSuspicion {
	queryKey := normalizedTerm(query)
	concepts := map[string][]string{}
	seen := map[string]map[string]bool{}
	for _, concept := range registry.Concepts {
		for _, term := range conceptTerms(concept) {
			key := normalizedTerm(term)
			if key == queryKey || strings.Contains(key, queryKey) || strings.Contains(queryKey, key) {
				if seen[key] == nil {
					seen[key] = map[string]bool{}
				}
				if seen[key][concept.CanonicalName] {
					continue
				}
				seen[key][concept.CanonicalName] = true
				concepts[key] = append(concepts[key], concept.CanonicalName)
			}
		}
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

func mustReadFiles(root string, concepts []Concept, hits []Hit) []string {
	seen := map[string]bool{}
	var files []string
	files = appendUniqueRel(files, seen, RegistryRelativePath)
	for _, concept := range concepts {
		for _, target := range concept.GeneratedTargets {
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(target))); err == nil {
				files = appendUniqueRel(files, seen, target)
			}
		}
	}
	for _, hit := range hits {
		files = appendUniqueRel(files, seen, hit.Path)
		if len(files) >= 12 {
			break
		}
	}
	return files
}

func appendUnique(values []string, seen map[string]bool, value string) []string {
	if seen[value] {
		return values
	}
	seen[value] = true
	return append(values, value)
}

func appendUniqueRel(values []string, seen map[string]bool, value string) []string {
	value = filepath.ToSlash(strings.TrimSpace(value))
	if value == "" || seen[value] {
		return values
	}
	seen[value] = true
	return append(values, value)
}

func normalizedTerm(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	replacer := strings.NewReplacer(" ", "", "-", "", "_", "", "/", "", ".", "")
	return replacer.Replace(value)
}
