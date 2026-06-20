package commandcenter

type EvidenceIntegritySummary struct {
	PublicSafe               bool   `json:"public_safe"`
	CheckedEvidenceRefCount  int    `json:"checked_evidence_ref_count"`
	PresentEvidenceRefCount  int    `json:"present_evidence_ref_count"`
	DanglingEvidenceRefCount int    `json:"dangling_evidence_ref_count"`
	NextSafeAction           string `json:"next_safe_action"`
}
