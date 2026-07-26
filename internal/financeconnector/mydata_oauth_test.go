package financeconnector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthAccessTokenUsesTossClientCredentialsGrant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost || request.URL.Path != tossAuthTokenPath {
			t.Fatalf("request = %s %s", request.Method, request.URL.Path)
		}
		if err := request.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if request.Form.Get("grant_type") != "client_credentials" || request.Form.Get("scope") != "ca" ||
			request.Form.Get("client_id") != "test-client" || request.Form.Get("client_secret") != "test-secret" {
			t.Fatalf("form = %#v", request.Form)
		}
		_, _ = writer.Write([]byte(`{"access_token":"issued-token","expires_in":31536000}`))
	}))
	defer server.Close()

	client := LiveClient{Config: RuntimeConfig{AuthTokenBaseURL: server.URL}, HTTP: server.Client()}
	token, err := client.authAccessToken(context.Background(), credentialEnvelope{
		ClientID: "test-client", ClientSecret: "test-secret", Provider: ProviderTossMyData,
	})
	if err != nil || token != "issued-token" {
		t.Fatalf("token = %q err = %v", token, err)
	}
}

func TestAuthAccessTokenRejectsMissingBundleAuthMaterial(t *testing.T) {
	client := LiveClient{Config: RuntimeConfig{AuthTokenBaseURL: "https://example.invalid"}}
	_, err := client.authAccessToken(context.Background(), credentialEnvelope{Provider: ProviderTossMyData})
	if err == nil || !strings.Contains(err.Error(), "auth token or client credentials") {
		t.Fatalf("err = %v", err)
	}
}
