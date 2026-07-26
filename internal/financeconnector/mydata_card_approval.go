package financeconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type myDataApprovalResponse struct {
	RspCode      string           `json:"rsp_code"`
	RspMsg       string           `json:"rsp_msg"`
	NextPage     string           `json:"next_page"`
	ApprovedList []myDataApproval `json:"approved_list"`
}

type myDataApproval struct {
	ApprovedNumber string          `json:"approved_num"`
	ApprovedAt     string          `json:"approved_dtime"`
	Status         string          `json:"status"`
	TransactionAt  string          `json:"trans_dtime"`
	MerchantName   string          `json:"merchant_name"`
	ApprovedAmount json.RawMessage `json:"approved_amt"`
	ModifiedAmount json.RawMessage `json:"modified_amt"`
	CurrencyCode   string          `json:"currency_code"`
	KRWAmount      json.RawMessage `json:"krw_amt"`
}

func (client LiveClient) fetchCardApprovals(
	ctx context.Context,
	connection ConnectionConfig,
	token string,
	cardID string,
	approvalPath string,
	scope string,
) ([]SourceTransaction, error) {
	var all []SourceTransaction
	nextPage := ""
	for page := 0; page < 100; page++ {
		query := url.Values{
			"org_code": {connection.OrgCode}, "from_date": {client.fromDate()},
			"to_date": {client.toDate()}, "limit": {"500"},
		}
		if nextPage != "" {
			query.Set("next_page", nextPage)
		}
		endpoint := path.Join(myDataCardsPath, cardID, approvalPath)
		data, err := client.doMyDataJSON(ctx, http.MethodGet, endpoint, query, token, nil)
		if err != nil {
			return nil, err
		}
		var response myDataApprovalResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		if err := validateMyDataResponse(response.RspCode, response.RspMsg); err != nil {
			return nil, err
		}
		for _, approval := range response.ApprovedList {
			transaction, err := approval.transaction(connection, cardID, scope, client.Config.Provider)
			if err != nil {
				return nil, err
			}
			all = append(all, transaction)
		}
		if response.NextPage == "" {
			return all, nil
		}
		nextPage = response.NextPage
	}
	return nil, fmt.Errorf("mydata card pagination exceeded 100 pages")
}
