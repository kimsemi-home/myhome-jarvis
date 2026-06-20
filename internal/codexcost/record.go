package codexcost

type Record struct {
	At            string   `json:"at"`
	Scope         string   `json:"scope"`
	UnitKind      string   `json:"unit_kind"`
	Amount        int64    `json:"amount"`
	Status        string   `json:"status"`
	EvidenceRefs  []string `json:"evidence_refs"`
	RawPrompt     string   `json:"raw_prompt,omitempty"`
	RawTranscript string   `json:"raw_transcript,omitempty"`
	PrivateNotes  string   `json:"private_notes,omitempty"`
}
