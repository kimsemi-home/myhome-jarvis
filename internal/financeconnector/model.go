package financeconnector

import "strings"

const (
	ProviderTossMyData = "toss_mydata"
	ParityThreshold    = 95.0
)

// SourceTransaction is the provider-neutral boundary used by fixture replay.
type SourceTransaction struct {
	TransactionID string   `json:"transaction_id"`
	Source        string   `json:"source"`
	Owner         string   `json:"owner"`
	OccurredAt    string   `json:"occurred_at"`
	PostedAt      string   `json:"posted_at"`
	Amount        int64    `json:"amount"`
	Currency      string   `json:"currency"`
	Direction     string   `json:"direction"`
	MerchantName  string   `json:"merchant_name"`
	Category      string   `json:"category"`
	AccountID     string   `json:"account_id"`
	CardID        string   `json:"card_id"`
	RawRef        string   `json:"raw_ref"`
	Tags          []string `json:"tags"`
}

type ParityField struct {
	Name          string `json:"name"`
	SourceField   string `json:"source_field"`
	CanonicalPath string `json:"canonical_path"`
	Required      bool   `json:"required"`
}

func (transaction SourceTransaction) fieldValue(name string) bool {
	switch name {
	case "transaction_id":
		return strings.TrimSpace(transaction.TransactionID) != ""
	case "source":
		return strings.TrimSpace(transaction.Source) != ""
	case "owner":
		return strings.TrimSpace(transaction.Owner) != ""
	case "occurred_at":
		return strings.TrimSpace(transaction.OccurredAt) != ""
	case "posted_at":
		return strings.TrimSpace(transaction.PostedAt) != ""
	case "amount":
		return transaction.Amount != 0
	case "currency":
		return strings.TrimSpace(transaction.Currency) != ""
	case "direction":
		return strings.TrimSpace(transaction.Direction) != ""
	case "merchant_name":
		return strings.TrimSpace(transaction.MerchantName) != ""
	case "category":
		return strings.TrimSpace(transaction.Category) != ""
	case "account_id":
		return strings.TrimSpace(transaction.AccountID) != ""
	case "card_id":
		return strings.TrimSpace(transaction.CardID) != ""
	case "raw_ref":
		return strings.TrimSpace(transaction.RawRef) != ""
	case "tags":
		return len(transaction.Tags) > 0
	default:
		return false
	}
}
