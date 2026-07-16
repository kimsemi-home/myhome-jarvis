package localfinanceevidence

type CreditBatchStatement struct {
	SourceNameSHA256 string              `json:"source_name_sha256"`
	Preview          CreditImportPreview `json:"preview"`
}

type CreditBatchPreview struct {
	SchemaVersion        string                 `json:"schema_version"`
	ExecutionMode        string                 `json:"execution_mode"`
	CredentialsRead      bool                   `json:"credentials_read"`
	ExternalNetwork      bool                   `json:"external_network"`
	ExternalWrites       bool                   `json:"external_writes"`
	ManifestSHA256       string                 `json:"manifest_sha256"`
	StatementCount       int                    `json:"statement_count"`
	ReadyCount           int                    `json:"ready_count"`
	AllReady             bool                   `json:"all_ready"`
	RawFileNamesReported bool                   `json:"raw_file_names_reported"`
	RawRowsReported      bool                   `json:"raw_rows_reported"`
	PreviewSetHash       string                 `json:"preview_set_hash"`
	Statements           []CreditBatchStatement `json:"statements"`
	Checks               []string               `json:"checks"`
	BatchHash            string                 `json:"batch_hash"`
}
