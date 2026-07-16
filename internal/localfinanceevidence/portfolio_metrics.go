package localfinanceevidence

type PortfolioReconciliation struct {
	CashMinor            int64 `json:"cash_minor"`
	SecuritiesValueMinor int64 `json:"securities_value_minor"`
	HoldingValueMinor    int64 `json:"holding_value_minor"`
	ProfitLossMinor      int64 `json:"profit_loss_minor"`
	TotalValueMinor      int64 `json:"total_value_minor"`
	Reconciled           bool  `json:"reconciled"`
}

type PortfolioMetrics struct {
	TokenRequests     int `json:"token_requests"`
	BalanceRequests   int `json:"balance_requests"`
	LedgerRequests    int `json:"ledger_requests"`
	InjectedFailures  int `json:"injected_failures"`
	OrderRequests     int `json:"order_requests"`
	ForbiddenRequests int `json:"forbidden_requests"`
	RedirectRequests  int `json:"redirect_requests"`
	OversizedRequests int `json:"oversized_requests"`
}
