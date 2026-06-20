package codexcost

type AttributionRecord struct {
	At           string   `json:"at"`
	SemanticHash string   `json:"semantic_hash,omitempty"`
	Scope        string   `json:"scope"`
	SubjectKey   string   `json:"subject_key"`
	SubjectHash  string   `json:"subject_hash,omitempty"`
	CostRef      string   `json:"cost_ref,omitempty"`
	UnitKind     string   `json:"unit_kind"`
	Amount       int64    `json:"amount"`
	Basis        string   `json:"basis"`
	EvidenceRefs []string `json:"evidence_refs"`
}

type AttributionRequest struct {
	At           string   `json:"at,omitempty"`
	Scope        string   `json:"scope"`
	SubjectKey   string   `json:"subject_key"`
	CostRef      string   `json:"cost_ref,omitempty"`
	UnitKind     string   `json:"unit_kind"`
	Amount       int64    `json:"amount"`
	Basis        string   `json:"basis"`
	EvidenceRefs []string `json:"evidence_refs"`
}

type AttributionResult struct {
	Scope            string `json:"scope"`
	UnitKind         string `json:"unit_kind"`
	Amount           int64  `json:"amount"`
	Basis            string `json:"basis"`
	SubjectHash      string `json:"subject_hash"`
	CostRef          string `json:"cost_ref"`
	EvidenceRefCount int    `json:"evidence_ref_count"`
	RecordedAt       string `json:"recorded_at"`
}
