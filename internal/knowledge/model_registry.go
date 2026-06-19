package knowledge

type Registry struct {
	BoundedContexts            []BoundedContext   `json:"bounded_contexts"`
	DDDPatterns                []string           `json:"ddd_patterns"`
	Concepts                   []Concept          `json:"concepts"`
	DomainEvents               []DomainEvent      `json:"domain_events"`
	HarnessCaseContracts       []HarnessCase      `json:"harness_case_contracts"`
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
	DDDKind          string   `json:"ddd_kind"`
	Description      string   `json:"description"`
	AllowedAliases   []string `json:"allowed_aliases"`
	Owner            string   `json:"owner"`
	GeneratedTargets []string `json:"generated_targets"`
	RelatedConcepts  []string `json:"related_concepts"`
}
