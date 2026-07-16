package localfinanceevidence

type CreditBatchApplyStatement struct {
	SourceNameSHA256 string `json:"source_name_sha256"`
	PreviewHash      string `json:"preview_hash"`
	RowsRead         int    `json:"rows_read"`
	RowsInserted     int    `json:"rows_inserted"`
	Duplicates       int    `json:"duplicates"`
	Suggestions      int    `json:"suggestions_recorded"`
}

type CreditBatchApplyReport struct {
	SchemaVersion        string                      `json:"schema_version"`
	ExecutionMode        string                      `json:"execution_mode"`
	CredentialsRead      bool                        `json:"credentials_read"`
	ExternalNetwork      bool                        `json:"external_network"`
	ExternalWrites       bool                        `json:"external_writes"`
	LocalDatabaseWrites  bool                        `json:"local_database_writes"`
	PlanHash             string                      `json:"plan_hash"`
	ApprovalHash         string                      `json:"approval_hash"`
	ManifestSHA256       string                      `json:"manifest_sha256"`
	PreviewSetHash       string                      `json:"preview_set_hash"`
	BatchHash            string                      `json:"batch_hash"`
	StatementCount       int                         `json:"statement_count"`
	TransactionCount     int                         `json:"transaction_count"`
	RowsInserted         int                         `json:"rows_inserted"`
	Duplicates           int                         `json:"duplicates"`
	SuggestionsRecorded  int                         `json:"suggestions_recorded"`
	AllOrNothing         bool                        `json:"all_or_nothing"`
	RawFileNamesReported bool                        `json:"raw_file_names_reported"`
	RawRowsReported      bool                        `json:"raw_rows_reported"`
	Statements           []CreditBatchApplyStatement `json:"statements"`
	ApplySetHash         string                      `json:"apply_set_hash"`
	Checks               []string                    `json:"checks"`
	ReportHash           string                      `json:"report_hash"`
}
