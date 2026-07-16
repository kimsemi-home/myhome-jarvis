package localfinanceevidence

type CreditExpectedTotals struct {
	Provided         bool  `json:"provided"`
	TransactionCount int   `json:"transaction_count"`
	DebitMinor       int64 `json:"debit_minor"`
	CreditMinor      int64 `json:"credit_minor"`
}

type CreditImportPreview struct {
	SchemaVersion      string               `json:"schema_version"`
	ExecutionMode      string               `json:"execution_mode"`
	CredentialsRead    bool                 `json:"credentials_read"`
	ExternalNetwork    bool                 `json:"external_network"`
	ExternalWrites     bool                 `json:"external_writes"`
	SourceSHA256       string               `json:"source_sha256"`
	SourceExtension    string               `json:"source_extension"`
	CandidateCount     int                  `json:"candidate_count"`
	TemplateID         string               `json:"template_id"`
	TemplateVersion    int                  `json:"template_version"`
	IssuerKey          string               `json:"issuer_key"`
	ProfileSHA256      string               `json:"profile_sha256"`
	TransactionCount   int                  `json:"transaction_count"`
	DebitMinor         int64                `json:"debit_minor"`
	CreditMinor        int64                `json:"credit_minor"`
	NetCardSpendMinor  int64                `json:"net_card_spend_minor"`
	FingerprintSetHash string               `json:"fingerprint_set_hash"`
	SuggestionCount    int                  `json:"suggestion_count"`
	Expected           CreditExpectedTotals `json:"expected_totals"`
	Reconciled         bool                 `json:"reconciled"`
	ReadyToImport      bool                 `json:"ready_to_import"`
	RawRowsReported    bool                 `json:"raw_rows_reported"`
	Checks             []string             `json:"checks"`
	PreviewHash        string               `json:"preview_hash"`
}

type CreditTemplateOnboarding struct {
	Version1Preview                 CreditImportPreview `json:"template_v1_preview"`
	Version2Preview                 CreditImportPreview `json:"template_v2_preview"`
	AmbiguousProfileRejected        bool                `json:"ambiguous_profile_rejected"`
	UnsupportedStatementRejected    bool                `json:"unsupported_statement_rejected"`
	MismatchedExpectedTotalsBlocked bool                `json:"mismatched_expected_totals_blocked"`
	SourceMutationRebound           bool                `json:"source_mutation_rebound"`
}
