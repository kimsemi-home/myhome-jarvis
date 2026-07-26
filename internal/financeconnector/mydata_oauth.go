package financeconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const tossAuthTokenPath = "/token"

type tossAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (client LiveClient) authAccessToken(
	ctx context.Context, credential credentialEnvelope,
) (string, error) {
	if strings.TrimSpace(credential.AuthAccessToken) != "" {
		return credential.AuthAccessToken, nil
	}
	if strings.TrimSpace(credential.ClientID) == "" || strings.TrimSpace(credential.ClientSecret) == "" {
		return "", fmt.Errorf("credential bundle auth token or client credentials are required")
	}
	if strings.TrimSpace(client.Config.AuthTokenBaseURL) == "" {
		return "", fmt.Errorf("toss auth token base URL is required")
	}
	form := url.Values{
		"grant_type": {"client_credentials"},
		"client_id":  {credential.ClientID}, "client_secret": {credential.ClientSecret},
		"scope": {"ca"},
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost,
		strings.TrimRight(client.Config.AuthTokenBaseURL, "/")+tossAuthTokenPath,
		strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpClient, err := client.httpClientForCredential(credential)
	if err != nil {
		return "", err
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("toss auth token request returned HTTP %d", response.StatusCode)
	}
	var decoded tossAuthTokenResponse
	if err := json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(&decoded); err != nil {
		return "", err
	}
	if strings.TrimSpace(decoded.AccessToken) == "" {
		return "", fmt.Errorf("toss auth token response is missing access token")
	}
	return decoded.AccessToken, nil
}
