package financeconnector

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (client LiveClient) doMyDataJSON(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	token string,
	body []byte,
) ([]byte, error) {
	endpoint := strings.TrimRight(client.Config.BaseURL, "/") + path
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}
	request, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("x-api-tran-id", transactionID())
	request.Header.Set("x-api-type", client.apiType())
	if len(body) > 0 {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if client.Config.ClientID != "" {
		request.Header.Set("x-client-id", client.Config.ClientID)
	}
	response, err := client.HTTP.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("mydata request returned HTTP %d", response.StatusCode)
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, 10<<20))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func validateMyDataResponse(code string, message string) error {
	if code != "" && code != "00000" {
		return fmt.Errorf("mydata response %s: %s", code, message)
	}
	return nil
}
