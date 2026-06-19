package knowledge

type DomainEvent struct {
	Name           string   `json:"name"`
	BoundedContext string   `json:"bounded_context"`
	Description    string   `json:"description"`
	EmittedBy      string   `json:"emitted_by"`
	PayloadFields  []string `json:"payload_fields"`
}

type HarnessCase struct {
	Name           string `json:"name"`
	BoundedContext string `json:"bounded_context"`
	Command        string `json:"command"`
	EvidenceTarget string `json:"evidence_target"`
	Description    string `json:"description"`
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
