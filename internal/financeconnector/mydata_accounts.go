package financeconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const myDataAccountsPath = "/v2/bank/accounts"

type myDataAccountListResponse struct {
	RspCode     string          `json:"rsp_code"`
	RspMsg      string          `json:"rsp_msg"`
	NextPage    string          `json:"next_page"`
	AccountList []myDataAccount `json:"account_list"`
}

type myDataAccount struct {
	AccountNumber string `json:"account_num"`
	IsConsent     bool   `json:"is_consent"`
}

func (client LiveClient) discoverAccounts(
	ctx context.Context,
	connection ConnectionConfig,
	credential credentialEnvelope,
) ([]string, error) {
	var accounts []string
	nextPage := ""
	for page := 0; page < 100; page++ {
		query := url.Values{
			"org_code": {connection.OrgCode}, "search_timestamp": {"0"},
			"limit": {"500"},
		}
		if nextPage != "" {
			query.Del("search_timestamp")
			query.Set("next_page", nextPage)
		}
		data, err := client.doMyDataJSON(ctx, http.MethodGet, myDataAccountsPath, query, credential, nil)
		if err != nil {
			return nil, err
		}
		var response myDataAccountListResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		if err := validateMyDataResponse(response.RspCode, response.RspMsg); err != nil {
			return nil, err
		}
		for _, account := range response.AccountList {
			if account.IsConsent && account.AccountNumber != "" {
				accounts = append(accounts, account.AccountNumber)
			}
		}
		if response.NextPage == "" {
			break
		}
		nextPage = response.NextPage
	}
	if len(accounts) == 0 {
		return nil, fmt.Errorf("mydata returned no consented bank accounts")
	}
	return accounts, nil
}
