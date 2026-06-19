package evidencequality

import "time"

type Status struct {
	PolicyPath            string         `json:"policy_path"`
	LedgerPath            string         `json:"ledger_path"`
	Exists                bool           `json:"exists"`
	SnapshotCount         int            `json:"snapshot_count"`
	InvalidSnapshotCount  int            `json:"invalid_snapshot_count"`
	ReassessmentDebtCount int            `json:"reassessment_debt_count"`
	MissingEvidenceCount  int            `json:"missing_evidence_count"`
	StaleSnapshotCount    int            `json:"stale_snapshot_count"`
	LowQualityCount       int            `json:"low_quality_count"`
	BlockedQualityCount   int            `json:"blocked_quality_count"`
	MappingDriftCount     int            `json:"mapping_drift_count"`
	StaleAfterHours       int            `json:"stale_after_hours"`
	ByQualityLevel        map[string]int `json:"by_quality_level"`
	ByMappingConfidence   map[string]int `json:"by_mapping_confidence"`
	ByPurpose             map[string]int `json:"by_purpose"`
	ByReassessmentReason  map[string]int `json:"by_reassessment_reason"`
	LastObservedAt        string         `json:"last_observed_at,omitempty"`
	CheckedAt             string         `json:"checked_at"`
}

func newStatus(policy Policy, checkedAt time.Time) Status {
	return Status{
		PolicyPath:           PolicyRelativePath,
		LedgerPath:           policy.PrivateSnapshotLedger,
		StaleAfterHours:      policy.StaleAfterHours,
		ByQualityLevel:       map[string]int{},
		ByMappingConfidence:  map[string]int{},
		ByPurpose:            map[string]int{},
		ByReassessmentReason: map[string]int{},
		CheckedAt:            checkedAt.Format(time.RFC3339),
	}
}
