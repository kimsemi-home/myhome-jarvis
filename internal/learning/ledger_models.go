package learning

type RecordRequest struct {
	Kind         string   `json:"kind"`
	Source       string   `json:"source"`
	Stage        string   `json:"stage,omitempty"`
	Status       string   `json:"status,omitempty"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs"`
	Owner        string   `json:"owner"`
	NextAction   string   `json:"next_action"`
}

type Observation struct {
	ID           string   `json:"id"`
	At           string   `json:"at"`
	Kind         string   `json:"kind"`
	Source       string   `json:"source"`
	Stage        string   `json:"stage"`
	Status       string   `json:"status"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs"`
	Owner        string   `json:"owner"`
	NextAction   string   `json:"next_action"`
}

type RecordResult struct {
	ID               string `json:"id"`
	Path             string `json:"path"`
	Kind             string `json:"kind"`
	Stage            string `json:"stage"`
	Status           string `json:"status"`
	EvidenceRefCount int    `json:"evidence_ref_count"`
	RecordedAt       string `json:"recorded_at"`
}
