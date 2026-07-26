package financeconnector

import (
	"fmt"
	"strings"
)

func (approval myDataApproval) transaction(
	connection ConnectionConfig,
	cardID string,
	scope string,
	provider string,
) (SourceTransaction, error) {
	amount, currencyCode, err := approval.amount()
	if err != nil {
		return SourceTransaction{}, err
	}
	id := approval.ApprovedNumber
	if id == "" {
		id = approval.ApprovedAt + "-" + cardID
	}
	postedAt := approval.TransactionAt
	if postedAt == "" {
		postedAt = approval.ApprovedAt
	}
	direction := "debit"
	if approval.Status == "02" {
		direction = "credit"
	}
	merchant := strings.TrimSpace(approval.MerchantName)
	if merchant == "" {
		merchant = unknownMerchant
	}
	return SourceTransaction{
		TransactionID: id, Source: provider, Owner: connection.Owner,
		OccurredAt: approval.ApprovedAt, PostedAt: postedAt, Amount: amount,
		Currency: currencyCode, Direction: direction, MerchantName: merchant,
		Category: "uncategorized", AccountID: "mydata-card", CardID: cardID,
		RawRef: "mydata:card:" + scope + ":" + id,
		Tags:   []string{"mydata", "card", scope},
	}, nil
}

func (approval myDataApproval) amount() (int64, string, error) {
	if approval.Status == "03" && len(approval.ModifiedAmount) > 0 {
		amount, err := parseMyDataAmount(approval.ModifiedAmount)
		if err == nil {
			return amount, normalizedCurrency(approval.CurrencyCode), nil
		}
	}
	if len(approval.KRWAmount) > 0 {
		amount, err := parseMyDataAmount(approval.KRWAmount)
		if err == nil {
			return amount, "KRW", nil
		}
	}
	amount, err := parseMyDataAmount(approval.ApprovedAmount)
	if err != nil && len(approval.ModifiedAmount) > 0 {
		amount, err = parseMyDataAmount(approval.ModifiedAmount)
	}
	if err != nil {
		return 0, "", fmt.Errorf("card approval amount: %w", err)
	}
	currency := strings.TrimSpace(approval.CurrencyCode)
	if currency == "" {
		currency = "KRW"
	}
	return amount, currency, nil
}
