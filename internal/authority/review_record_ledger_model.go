package authority

type ReviewRecordLedgerSummary struct {
	RequestID          string
	Recorded           bool
	RecordCount        int
	InvalidRecordCount int
	LedgerState        string
	ApprovalState      string
	LastRecordedAt     string
}
