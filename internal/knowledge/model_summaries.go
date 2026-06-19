package knowledge

type DomainEventSummary struct {
	Name           string   `json:"name"`
	BoundedContext string   `json:"bounded_context"`
	EmittedBy      string   `json:"emitted_by"`
	PayloadFields  []string `json:"payload_fields"`
}

type HarnessCaseSummary struct {
	Name           string `json:"name"`
	BoundedContext string `json:"bounded_context"`
	Command        string `json:"command"`
	EvidenceTarget string `json:"evidence_target"`
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
