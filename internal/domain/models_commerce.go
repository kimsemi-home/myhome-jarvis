package domain

type CommerceSummary struct {
	Records                 int                          `json:"records"`
	Currency                string                       `json:"currency"`
	TotalSpendMinorUnits    int64                        `json:"total_spend_minor_units"`
	RecurringCandidateCount int                          `json:"recurring_candidate_count"`
	RecurringCandidates     []RecurringPurchaseCandidate `json:"recurring_candidates"`
	Categories              []string                     `json:"categories"`
	OwnerBreakdown          []CommerceOwnerSummary       `json:"owner_breakdown"`
}

type CommerceOwnerSummary struct {
	Owner                   string `json:"owner"`
	Records                 int    `json:"records"`
	Currency                string `json:"currency"`
	PurchaseSpendMinorUnits int64  `json:"purchase_spend_minor_units"`
}

type RecurringPurchaseCandidate struct {
	MerchantName          string `json:"merchant_name"`
	ItemName              string `json:"item_name"`
	Currency              string `json:"currency"`
	PurchaseCount         int    `json:"purchase_count"`
	LatestTotalMinorUnits int64  `json:"latest_total_minor_units"`
	LatestPurchasedAt     string `json:"latest_purchased_at"`
}

type commercePurchase struct {
	Owner              string      `json:"owner"`
	PurchasedAt        string      `json:"purchased_at"`
	MerchantName       string      `json:"merchant_name"`
	ItemName           string      `json:"item_name"`
	TotalPrice         moneyAmount `json:"total_price"`
	Category           string      `json:"category"`
	RecurringCandidate bool        `json:"recurring_candidate"`
}
