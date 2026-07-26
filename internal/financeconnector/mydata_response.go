package financeconnector

import (
	"encoding/json"
	"fmt"
	"strings"
)

type myDataResponse struct {
	RspCode   string              `json:"rsp_code"`
	RspMsg    string              `json:"rsp_msg"`
	NextPage  string              `json:"next_page"`
	TransList []myDataTransaction `json:"trans_list"`
}

type myDataTransaction struct {
	TransDTime   string          `json:"trans_dtime"`
	TransNo      string          `json:"trans_no"`
	TransType    string          `json:"trans_type"`
	CurrencyCode string          `json:"currency_code"`
	TransAmount  json.RawMessage `json:"trans_amt"`
	TransMemo    string          `json:"trans_memo"`
}

func (response myDataResponse) transactions(connection ConnectionConfig, provider string) ([]SourceTransaction, error) {
	transactions := make([]SourceTransaction, 0, len(response.TransList))
	for index, item := range response.TransList {
		amount, err := parseMyDataAmount(item.TransAmount)
		if err != nil {
			return nil, fmt.Errorf("transaction %d amount: %w", index, err)
		}
		direction, err := myDataDirection(item.TransType)
		if err != nil {
			return nil, err
		}
		id := item.TransNo
		if id == "" {
			id = fmt.Sprintf("%s-%d", item.TransDTime, index)
		}
		merchant := strings.TrimSpace(item.TransMemo)
		if merchant == "" {
			merchant = unknownMerchant
		}
		transactions = append(transactions, SourceTransaction{
			TransactionID: id, Source: provider, Owner: connection.Owner,
			OccurredAt: item.TransDTime, PostedAt: item.TransDTime,
			Amount: amount, Currency: currency(item.CurrencyCode), Direction: direction,
			MerchantName: merchant, Category: "uncategorized",
			AccountID: "mydata-account", RawRef: "mydata:" + id, Tags: []string{"mydata"},
		})
	}
	return transactions, nil
}
