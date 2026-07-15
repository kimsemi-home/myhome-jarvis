package localfinanceevidence

type PortfolioReport struct {
	SchemaVersion    string                  `json:"schema_version"`
	ExecutionMode    string                  `json:"execution_mode"`
	LoopbackOnly     bool                    `json:"loopback_only"`
	CredentialsRead  bool                    `json:"credentials_read"`
	ExternalNetwork  bool                    `json:"external_network"`
	ExternalWrites   bool                    `json:"external_writes"`
	FinancialActions bool                    `json:"financial_actions"`
	Month            string                  `json:"month"`
	KIS              PortfolioKIS            `json:"kis_readonly_sync"`
	Store            PortfolioStore          `json:"local_store_replay"`
	Ledger           PortfolioLedger         `json:"ledger_aggregate_replay"`
	Monthly          PortfolioReconciliation `json:"monthly_reconciliation"`
	Emulator         PortfolioMetrics        `json:"emulator_metrics"`
	Checks           []string                `json:"checks"`
	ReportHash       string                  `json:"report_hash"`
}

type PortfolioKIS struct {
	Method       string `json:"method"`
	Path         string `json:"path"`
	Transaction  string `json:"transaction_id"`
	Retries      int    `json:"retries"`
	HoldingCount int    `json:"holding_count"`
}

type PortfolioStore struct {
	FirstSnapshotCount  int `json:"first_snapshot_count"`
	ReplaySnapshotCount int `json:"replay_snapshot_count"`
}

type PortfolioLedger struct {
	FirstInserted      bool `json:"first_inserted"`
	ReplayInserted     bool `json:"replay_inserted"`
	FingerprintMatched bool `json:"fingerprint_matched"`
	AggregateOnly      bool `json:"aggregate_only"`
}
