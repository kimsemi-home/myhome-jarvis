package evidence

type Policy struct {
	Context                  string          `json:"context"`
	Version                  string          `json:"version"`
	GeneratedArtifact        string          `json:"generated_artifact"`
	PrivateRoot              string          `json:"private_root"`
	PrivateGraphRequired     bool            `json:"private_graph_required"`
	PublicStatusRedacted     bool            `json:"public_status_redacted"`
	RawEvidencePublicAllowed bool            `json:"raw_evidence_public_allowed"`
	NodeKinds                []string        `json:"node_kinds"`
	EdgeKinds                []string        `json:"edge_kinds"`
	PrivateSources           []PrivateSource `json:"private_sources"`
	AllowedEvidencePrefixes  []string        `json:"allowed_evidence_prefixes"`
	PublicSummaryFields      []string        `json:"public_summary_fields"`
	ForbiddenPublicFields    []string        `json:"forbidden_public_fields"`
	Commands                 []string        `json:"commands"`
}

type PrivateSource struct {
	Key      string `json:"key"`
	Path     string `json:"path"`
	NodeKind string `json:"node_kind"`
	Format   string `json:"format"`
}

type SourceStatus struct {
	Key      string `json:"key"`
	NodeKind string `json:"node_kind"`
	Format   string `json:"format"`
	Present  bool   `json:"present"`
	Count    int    `json:"count"`
}

type Status struct {
	PolicyPath               string         `json:"policy_path"`
	PrivateRoot              string         `json:"private_root"`
	SourceCount              int            `json:"source_count"`
	PresentSourceCount       int            `json:"present_source_count"`
	NodeCount                int            `json:"node_count"`
	EdgeCount                int            `json:"edge_count"`
	DanglingEvidenceRefCount int            `json:"dangling_evidence_ref_count"`
	OpenLearningCount        int            `json:"open_learning_count"`
	ByNodeKind               map[string]int `json:"by_node_kind"`
	ByEdgeKind               map[string]int `json:"by_edge_kind"`
	Sources                  []SourceStatus `json:"sources"`
	LastObservedAt           string         `json:"last_observed_at,omitempty"`
	CheckedAt                string         `json:"checked_at"`
}

type learningObservation struct {
	ID           string   `json:"id"`
	At           string   `json:"at"`
	Status       string   `json:"status"`
	EvidenceRefs []string `json:"evidence_refs"`
}
