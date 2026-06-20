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

type RecordRequest struct {
	At               string   `json:"at,omitempty"`
	ConsentKind      string   `json:"consent_kind"`
	SubjectScope     string   `json:"subject_scope"`
	Status           string   `json:"status"`
	ReviewStatus     string   `json:"review_status"`
	AuthorityProfile string   `json:"authority_profile"`
	EvidenceRefs     []string `json:"evidence_refs"`
	ExpiresAt        string   `json:"expires_at,omitempty"`
}

type RecordResult struct {
	ConsentKind                 string `json:"consent_kind"`
	SubjectScope                string `json:"subject_scope"`
	Status                      string `json:"status"`
	ReviewStatus                string `json:"review_status"`
	AuthorityProfile            string `json:"authority_profile"`
	EvidenceRefCount            int    `json:"evidence_ref_count"`
	ReadinessState              string `json:"readiness_state"`
	ActiveConsentCount          int    `json:"active_consent_count"`
	MissingRequiredConsentCount int    `json:"missing_required_consent_count"`
	ConsentDebtCount            int    `json:"consent_debt_count"`
	RecordedAt                  string `json:"recorded_at"`
}
