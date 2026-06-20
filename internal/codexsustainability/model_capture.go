package codexsustainability

type CaptureStatus struct {
	PolicyPath                 string `json:"policy_path"`
	QualityLedgerPath          string `json:"quality_ledger_path"`
	SustainabilityLedgerPath   string `json:"sustainability_ledger_path"`
	CaptureState               string `json:"capture_state"`
	QualityRunCount            int    `json:"quality_run_count"`
	LastQualityOK              bool   `json:"last_quality_ok"`
	TrendBaselineVersion       string `json:"trend_baseline_version,omitempty"`
	CycleMinutes               int64  `json:"cycle_minutes,omitempty"`
	RecordedRecordCount        int    `json:"recorded_record_count"`
	EvidenceRef                string `json:"evidence_ref,omitempty"`
	PublicSafe                 bool   `json:"public_safe"`
	Redaction                  string `json:"redaction"`
	ApprovalState              string `json:"approval_state"`
	ApprovalGranted            bool   `json:"approval_granted"`
	ExternalWritesAllowed      bool   `json:"external_writes_allowed"`
	SelfApprovalAllowed        bool   `json:"self_approval_allowed"`
	RawEvidencePublicAllowed   bool   `json:"raw_evidence_public_allowed"`
	PrivateLedgerWriteRequired bool   `json:"private_ledger_write_required"`
	CheckedAt                  string `json:"checked_at"`
}
