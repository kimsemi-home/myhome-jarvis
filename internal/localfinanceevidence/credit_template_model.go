package localfinanceevidence

type CreditTemplateImport struct {
	RunID                 string `json:"run_id"`
	FileSHA256            string `json:"file_sha256"`
	TemplateID            string `json:"template_id"`
	TemplateVersion       int    `json:"template_version"`
	ProfileSHA256         string `json:"profile_sha256"`
	FingerprintSetHash    string `json:"fingerprint_set_hash"`
	Read                  int    `json:"read"`
	Inserted              int    `json:"inserted"`
	Duplicates            int    `json:"duplicates"`
	SuggestionsRecorded   int    `json:"suggestions_recorded"`
	NormalizedDebitMinor  int64  `json:"normalized_debit_minor"`
	NormalizedCreditMinor int64  `json:"normalized_credit_minor"`
}

type CreditTemplateHistory struct {
	EntryCount                int    `json:"entry_count"`
	Version1Entries           int    `json:"version_1_entries"`
	Version2Entries           int    `json:"version_2_entries"`
	HistoryHash               string `json:"history_hash"`
	ExistingCategoryPreserved bool   `json:"existing_category_preserved"`
	RawRowsReported           bool   `json:"raw_rows_reported"`
}

type CreditTemplateReconciliation struct {
	TransactionCount  int   `json:"transaction_count"`
	CardSpendMinor    int64 `json:"card_spend_minor"`
	CardRefundMinor   int64 `json:"card_refund_minor"`
	NetCardSpendMinor int64 `json:"net_card_spend_minor"`
	BothVersionsMatch bool  `json:"both_versions_match"`
	MonthlyMatch      bool  `json:"monthly_match"`
}

type CreditTemplateGuards struct {
	StableIdentityConflictRejected  bool `json:"stable_identity_conflict_rejected"`
	MissingSourceIDRejected         bool `json:"missing_source_id_rejected"`
	TemplateVersionMutationRejected bool `json:"template_version_mutation_rejected"`
	PartialWritesAfterRejection     bool `json:"partial_writes_after_rejection"`
}

type CreditTemplateReport struct {
	SchemaVersion   string                       `json:"schema_version"`
	ExecutionMode   string                       `json:"execution_mode"`
	CredentialsRead bool                         `json:"credentials_read"`
	ExternalNetwork bool                         `json:"external_network"`
	ExternalWrites  bool                         `json:"external_writes"`
	Month           string                       `json:"month"`
	FirstImport     CreditTemplateImport         `json:"template_v1_import"`
	SecondImport    CreditTemplateImport         `json:"template_v2_import"`
	History         CreditTemplateHistory        `json:"classification_history"`
	Reconciliation  CreditTemplateReconciliation `json:"reconciliation"`
	Guards          CreditTemplateGuards         `json:"guard_attacks"`
	Checks          []string                     `json:"checks"`
	ReportHash      string                       `json:"report_hash"`
}
