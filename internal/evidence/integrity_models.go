package evidence

type IntegrityStatus struct {
	PolicyPath               string            `json:"policy_path"`
	PrivateRoot              string            `json:"private_root"`
	PublicSafe               bool              `json:"public_safe"`
	Redaction                string            `json:"redaction"`
	CheckedEvidenceRefCount  int               `json:"checked_evidence_ref_count"`
	PresentEvidenceRefCount  int               `json:"present_evidence_ref_count"`
	DanglingEvidenceRefCount int               `json:"dangling_evidence_ref_count"`
	PrefixCounts             []IntegrityPrefix `json:"prefix_counts"`
	NextSafeAction           string            `json:"next_safe_action"`
	CheckedAt                string            `json:"checked_at"`
}

type IntegrityPrefix struct {
	Prefix        string `json:"prefix"`
	CheckedCount  int    `json:"checked_count"`
	PresentCount  int    `json:"present_count"`
	DanglingCount int    `json:"dangling_count"`
}
