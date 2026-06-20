package codexcost

type Record struct {
	At            string   `json:"at"`
	SemanticHash  string   `json:"semantic_hash,omitempty"`
	Scope         string   `json:"scope"`
	UnitKind      string   `json:"unit_kind"`
	Amount        int64    `json:"amount"`
	Status        string   `json:"status"`
	EvidenceRefs  []string `json:"evidence_refs"`
	RawPrompt     string   `json:"raw_prompt,omitempty"`
	RawTranscript string   `json:"raw_transcript,omitempty"`
	PrivateNotes  string   `json:"private_notes,omitempty"`
}

type RecordRequest struct {
	At           string   `json:"at,omitempty"`
	Scope        string   `json:"scope"`
	UnitKind     string   `json:"unit_kind"`
	Amount       int64    `json:"amount"`
	Status       string   `json:"status,omitempty"`
	EvidenceRefs []string `json:"evidence_refs"`
}

type RecordResult struct {
	Scope            string `json:"scope"`
	UnitKind         string `json:"unit_kind"`
	Amount           int64  `json:"amount"`
	Status           string `json:"status"`
	EvidenceRefCount int    `json:"evidence_ref_count"`
	BudgetState      string `json:"budget_state"`
	RecordedAt       string `json:"recorded_at"`
}
