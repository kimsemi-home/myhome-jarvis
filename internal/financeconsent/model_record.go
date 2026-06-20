package financeconsent

type Record struct {
	At               string   `json:"at"`
	ConsentKind      string   `json:"consent_kind"`
	SubjectScope     string   `json:"subject_scope"`
	Status           string   `json:"status"`
	ReviewStatus     string   `json:"review_status"`
	AuthorityProfile string   `json:"authority_profile"`
	EvidenceRefs     []string `json:"evidence_refs"`
	ExpiresAt        string   `json:"expires_at,omitempty"`
	PrivateNotes     string   `json:"private_notes,omitempty"`
}
