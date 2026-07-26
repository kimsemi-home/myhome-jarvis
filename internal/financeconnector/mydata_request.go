package financeconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const myDataTransactionsPath = "/v2/bank/accounts/deposit/transactions"

func (client LiveClient) requestPage(
	ctx context.Context,
	connection ConnectionConfig,
	token string,
	nextPage string,
) (myDataResponse, error) {
	body := map[string]string{
		"org_code":    connection.OrgCode,
		"account_num": connection.AccountNumber,
		"from_date":   client.fromDate(),
		"to_date":     client.toDate(),
		"limit":       "500",
	}
	if nextPage != "" {
		body["next_page"] = nextPage
	}
	encoded, err := json.Marshal(body)
	if err != nil {
		return myDataResponse{}, err
	}
	endpoint := strings.TrimRight(client.Config.BaseURL, "/") + myDataTransactionsPath
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(string(encoded)))
	if err != nil {
		return myDataResponse{}, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-tran-id", transactionID())
	request.Header.Set("x-api-type", "2")
	if client.Config.ClientID != "" {
		request.Header.Set("x-client-id", client.Config.ClientID)
	}
	response, err := client.HTTP.Do(request)
	if err != nil {
		return myDataResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return myDataResponse{}, fmt.Errorf("mydata request returned HTTP %d", response.StatusCode)
	}
	var decoded myDataResponse
	if err := json.NewDecoder(response.Body).Decode(&decoded); err != nil {
		return myDataResponse{}, err
	}
	if decoded.RspCode != "" && decoded.RspCode != "00000" {
		return myDataResponse{}, fmt.Errorf("mydata response %s: %s", decoded.RspCode, decoded.RspMsg)
	}
	return decoded, nil
}
