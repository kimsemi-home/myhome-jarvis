package localfinanceevidence

type OperatorCalls struct {
	Ledger    int `json:"ledger"`
	Portfolio int `json:"portfolio"`
	Revenue   int `json:"revenue"`
}

type OperatorSourceHashes struct {
	Ledger    string `json:"ledger"`
	Portfolio string `json:"portfolio"`
	Revenue   string `json:"revenue"`
}

type OperatorSnapshot struct {
	SchemaVersion         string               `json:"schema_version"`
	Month                 string               `json:"month"`
	Currency              string               `json:"currency"`
	RevenueStatus         string               `json:"revenue_status"`
	TransactionCount      int                  `json:"transaction_count"`
	CardSpendMinor        int64                `json:"card_spend_minor"`
	CardRefundMinor       int64                `json:"card_refund_minor"`
	NetCardSpendMinor     int64                `json:"net_card_spend_minor"`
	LedgerIncomeMinor     int64                `json:"ledger_income_minor"`
	LedgerCostMinor       int64                `json:"ledger_cost_minor"`
	LedgerNetIncomeMinor  int64                `json:"ledger_net_income_minor"`
	YouTubeGrossMinor     int64                `json:"youtube_gross_minor"`
	YouTubeCostMinor      int64                `json:"youtube_cost_minor"`
	YouTubeNetMinor       int64                `json:"youtube_net_minor"`
	TrackedSurplusMinor   int64                `json:"tracked_surplus_minor"`
	AssetAsOf             string               `json:"asset_as_of"`
	AssetCashMinor        int64                `json:"asset_cash_minor"`
	AssetSecuritiesMinor  int64                `json:"asset_securities_minor"`
	AssetProfitLossMinor  int64                `json:"asset_profit_loss_minor"`
	AssetTotalMinor       int64                `json:"asset_total_minor"`
	LiquidityRatioBPS     int64                `json:"liquidity_ratio_bps"`
	CreditSpendToAssetBPS int64                `json:"credit_spend_to_asset_bps"`
	SourceHashes          OperatorSourceHashes `json:"source_hashes"`
	Checks                []string             `json:"checks"`
	SnapshotHash          string               `json:"snapshot_hash"`
}

type OperatorReport struct {
	SchemaVersion          string           `json:"schema_version"`
	Component              string           `json:"component"`
	ExecutionMode          string           `json:"execution_mode"`
	Month                  string           `json:"month"`
	LoopbackOnly           bool             `json:"loopback_only"`
	CredentialsRead        bool             `json:"credentials_read"`
	ExternalNetworkEnabled bool             `json:"external_network_enabled"`
	ExternalWritesEnabled  bool             `json:"external_writes_enabled"`
	FinancialActions       bool             `json:"financial_actions"`
	ChannelWrites          bool             `json:"channel_writes"`
	ServiceInstalled       bool             `json:"service_installed"`
	Day2LedgerCompleted    bool             `json:"day2_ledger_completed"`
	Day3Completed          bool             `json:"day3_completed"`
	Day3FailedStage        string           `json:"day3_failed_stage"`
	Day3SkippedLedger      bool             `json:"day3_skipped_ledger"`
	Day5ResumeCompleted    bool             `json:"day5_resume_completed"`
	Day5SkippedLedger      bool             `json:"day5_skipped_ledger"`
	Day6ReplayCompleted    bool             `json:"day6_replay_completed"`
	Day6AllStagesSkipped   bool             `json:"day6_all_stages_skipped"`
	RetryExitCode          int              `json:"retry_exit_code"`
	MaxAttemptsPerRun      int              `json:"max_attempts_per_run"`
	CollectorCalls         OperatorCalls    `json:"collector_calls"`
	SummaryCalls           OperatorCalls    `json:"summary_calls"`
	StageRows              int              `json:"stage_rows"`
	SnapshotRows           int              `json:"snapshot_rows"`
	Snapshot               OperatorSnapshot `json:"snapshot"`
	Checks                 []string         `json:"checks"`
	ReportHash             string           `json:"report_hash"`
}
