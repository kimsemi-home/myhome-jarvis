package authority

type ReviewRequestEvidenceStatus struct {
	PolicyPath              string   `json:"policy_path"`
	RequestID               string   `json:"request_id"`
	RequestState            string   `json:"request_state"`
	EvidenceRef             string   `json:"evidence_ref"`
	EvidenceState           string   `json:"evidence_state"`
	EvidenceReady           bool     `json:"evidence_ready"`
	PublicSafe              bool     `json:"public_safe"`
	Redaction               string   `json:"redaction"`
	ApprovalState           string   `json:"approval_state"`
	ApprovalGranted         bool     `json:"approval_granted"`
	ExternalWritesAllowed   bool     `json:"external_writes_allowed"`
	SelfApprovalAllowed     bool     `json:"self_approval_allowed"`
	RequiredEvidenceFields  []string `json:"required_evidence_fields"`
	MissingEvidenceFieldCnt int      `json:"missing_evidence_field_count"`
	NextSafeAction          string   `json:"next_safe_action"`
	CheckedAt               string   `json:"checked_at"`
}
