package localfinanceevidence

type RevenueReport struct {
	SchemaVersion          string   `json:"schema_version"`
	Component              string   `json:"component"`
	ExecutionMode          string   `json:"execution_mode"`
	Month                  string   `json:"month"`
	Currency               string   `json:"currency"`
	RevenueStatus          string   `json:"revenue_status"`
	LoopbackOnly           bool     `json:"loopback_only"`
	CredentialsRead        bool     `json:"credentials_read"`
	OAuthPerformed         bool     `json:"oauth_performed"`
	ExternalNetworkEnabled bool     `json:"external_network_enabled"`
	ExternalWritesEnabled  bool     `json:"external_writes_enabled"`
	ChannelWritesEnabled   bool     `json:"channel_writes_enabled"`
	ServiceInstalled       bool     `json:"service_installed"`
	ChannelIdentityBound   bool     `json:"channel_identity_bound"`
	DailyRows              int      `json:"daily_rows"`
	VideoRows              int      `json:"video_rows"`
	PersistedDailyRows     int      `json:"persisted_daily_rows"`
	PersistedVideoRows     int      `json:"persisted_video_rows"`
	CostRows               int      `json:"cost_rows"`
	CostReplayDuplicates   int      `json:"cost_replay_duplicates"`
	GrossMinor             int64    `json:"gross_minor"`
	CostMinor              int64    `json:"cost_minor"`
	NetMinor               int64    `json:"net_minor"`
	Views                  int64    `json:"views"`
	WatchMinutes           float64  `json:"watch_minutes"`
	MonetizedPlaybacks     int64    `json:"monetized_playbacks"`
	ChannelRequests        int      `json:"channel_requests"`
	DailyReportRequests    int      `json:"daily_report_requests"`
	VideoReportRequests    int      `json:"video_report_requests"`
	LedgerRequests         int      `json:"ledger_requests"`
	InjectedFailures       int      `json:"injected_failures"`
	ForbiddenRequests      int      `json:"forbidden_requests"`
	ChannelWriteRequests   int      `json:"channel_write_requests"`
	LedgerStoredEvents     int      `json:"ledger_stored_events"`
	FirstLedgerInserted    bool     `json:"first_ledger_inserted"`
	ReplayLedgerInserted   bool     `json:"replay_ledger_inserted"`
	LedgerFingerprint      string   `json:"ledger_fingerprint"`
	EvidenceHash           string   `json:"evidence_hash"`
	Checks                 []string `json:"checks"`
	ReportHash             string   `json:"report_hash"`
}
