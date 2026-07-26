package financeconnector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SignVerification completes the Toss MyData authentication-plane callback
// with the same credential bundle used by the data plane. It returns the CI
// only to the caller that owns the private server-side flow.
func (client LiveClient) SignVerification(
	ctx context.Context,
	credential credentialEnvelope,
	request SignVerificationRequest,
) (string, error) {
	if strings.TrimSpace(credential.AuthAccessToken) == "" {
		return "", fmt.Errorf("credential bundle auth access token is missing")
	}
	if strings.TrimSpace(client.Config.AuthBaseURL) == "" {
		return "", fmt.Errorf("mydata auth base URL is required")
	}
	if strings.TrimSpace(request.CertTxID) == "" || strings.TrimSpace(request.TxID) == "" {
		return "", fmt.Errorf("mydata auth request transaction ids are required")
	}
	if strings.TrimSpace(request.SignedConsent) == "" || strings.TrimSpace(request.Consent) == "" {
		return "", fmt.Errorf("mydata auth request consent proof is required")
	}
	encoded, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	httpClient, err := client.httpClientForCredential(credential)
	if err != nil {
		return "", err
	}
	endpoint := strings.TrimRight(client.Config.AuthBaseURL, "/") + myDataSignVerificationPath
	httpRequest, err := http.NewRequestWithContext(
		ctx, http.MethodPost, endpoint, bytes.NewReader(encoded),
	)
	if err != nil {
		return "", err
	}
	httpRequest.Header.Set("Authorization", "Bearer "+credential.AuthAccessToken)
	httpRequest.Header.Set("x-api-tran-id", transactionID())
	httpRequest.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("mydata auth request returned HTTP %d", response.StatusCode)
	}
	var decoded signVerificationResponse
	limited := io.LimitReader(response.Body, 1<<20)
	if err := json.NewDecoder(limited).Decode(&decoded); err != nil {
		return "", err
	}
	if err := validateMyDataResponse(decoded.RspCode, decoded.RspMsg); err != nil {
		return "", err
	}
	if !strings.EqualFold(decoded.Result, "true") {
		return "", fmt.Errorf("mydata auth failed: %s", decoded.RspMsg)
	}
	if strings.TrimSpace(decoded.UserCI) == "" {
		return "", fmt.Errorf("mydata auth response user CI is missing")
	}
	return decoded.UserCI, nil
}
