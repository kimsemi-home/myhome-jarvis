package financeconnector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const myDataTransactionsPath = "/v2/bank/accounts/deposit/transactions"

func (client LiveClient) requestPage(
	ctx context.Context,
	connection ConnectionConfig,
	credential credentialEnvelope,
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
	data, err := client.doMyDataJSON(
		ctx, http.MethodPost, myDataTransactionsPath, url.Values{}, credential, encoded,
	)
	if err != nil {
		return myDataResponse{}, err
	}
	var decoded myDataResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		return myDataResponse{}, err
	}
	if err := validateMyDataResponse(decoded.RspCode, decoded.RspMsg); err != nil {
		return myDataResponse{}, err
	}
	return decoded, nil
}
