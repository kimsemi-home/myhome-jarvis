package knowledge

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
	EventCount   int     `json:"event_count"`
	HarnessCount int     `json:"harness_count"`
	Checks       []Check `json:"checks"`
}

type SearchReport struct {
	Query               string               `json:"query"`
	CheckedAt           string               `json:"checked_at"`
	Concepts            []ConceptSummary     `json:"concepts"`
	Events              []DomainEventSummary `json:"events,omitempty"`
	HarnessCases        []HarnessCaseSummary `json:"harness_cases,omitempty"`
	Hits                []Hit                `json:"hits"`
	LinearIssues        []string             `json:"linear_issues,omitempty"`
	DuplicateSuspicions []DuplicateSuspicion `json:"duplicate_suspicions"`
	MustRead            []string             `json:"must_read"`
}

type ConceptSummary struct {
	CanonicalName    string   `json:"canonical_name"`
	BoundedContext   string   `json:"bounded_context"`
	DDDKind          string   `json:"ddd_kind"`
	Owner            string   `json:"owner"`
	Definition       string   `json:"definition"`
	AllowedAliases   []string `json:"allowed_aliases"`
	GeneratedTargets []string `json:"generated_targets"`
	RelatedConcepts  []string `json:"related_concepts"`
}
