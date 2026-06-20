package storagearchive

type manifestSummary struct {
	Present             bool
	EntryCount          int
	ArchivedCount       int
	SkippedCount        int
	BudgetBreachCount   int
	InvalidEntryCount   int
	ArchivedInputBytes  int64
	ArchivedOutputBytes int64
	CompressionRatio    int
	LastEntryAt         string
	LastArchivedAt      string
	LastBudgetBreachAt  string
}
