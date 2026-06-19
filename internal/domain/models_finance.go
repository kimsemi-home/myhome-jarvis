package domain

type FinanceSummary struct {
	Records                int                   `json:"records"`
	Currency               string                `json:"currency"`
	CreditMinorUnits       int64                 `json:"credit_minor_units"`
	DebitMinorUnits        int64                 `json:"debit_minor_units"`
	NetMinorUnits          int64                 `json:"net_minor_units"`
	SubscriptionMinorUnits int64                 `json:"subscription_minor_units"`
	SubscriptionCount      int                   `json:"subscription_count"`
	CardDebitMinorUnits    int64                 `json:"card_debit_minor_units"`
	CardDebitCount         int                   `json:"card_debit_count"`
	Categories             []string              `json:"categories"`
	OwnerBreakdown         []FinanceOwnerSummary `json:"owner_breakdown"`
}

type FinanceOwnerSummary struct {
	Owner            string `json:"owner"`
	Records          int    `json:"records"`
	Currency         string `json:"currency"`
	CreditMinorUnits int64  `json:"credit_minor_units"`
	DebitMinorUnits  int64  `json:"debit_minor_units"`
	NetMinorUnits    int64  `json:"net_minor_units"`
}

type moneyAmount struct {
	MinorUnits int64  `json:"minor_units"`
	Currency   string `json:"currency"`
}

type financeTransaction struct {
	Owner     string      `json:"owner"`
	Amount    moneyAmount `json:"amount"`
	Direction string      `json:"direction"`
	Category  string      `json:"category"`
	CardID    string      `json:"card_id"`
}
