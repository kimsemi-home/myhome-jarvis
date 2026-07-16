package localfinanceevidence

type CreditReport struct {
	SchemaVersion         string         `json:"schema_version"`
	ExecutionMode         string         `json:"execution_mode"`
	LoopbackOnly          bool           `json:"loopback_only"`
	CredentialsRead       bool           `json:"credentials_read"`
	ExternalNetwork       bool           `json:"external_network"`
	ExternalWrites        bool           `json:"external_writes"`
	Month                 string         `json:"month"`
	AttachmentContentHash string         `json:"attachment_content_hash"`
	OAuth                 CreditOAuth    `json:"oauth_token_boundary"`
	FirstGmail            CreditSync     `json:"first_gmail_sync"`
	FirstWatch            CreditWatch    `json:"first_inbox_watch"`
	ReplayGmail           CreditSync     `json:"replay_gmail_sync"`
	ReplayWatch           CreditWatch    `json:"replay_inbox_watch"`
	ArchiveFallbackGmail  CreditSync     `json:"archive_fallback_gmail_sync"`
	ArchiveFallbackWatch  CreditWatch    `json:"archive_fallback_inbox_watch"`
	Monthly               CreditMonthly  `json:"monthly_reconciliation"`
	Emulator              CreditEmulator `json:"emulator_metrics"`
	Checks                []string       `json:"checks"`
	ReportHash            string         `json:"report_hash"`
}
