package financeconnector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignVerificationFromReferenceRefreshesRejectedCachedToken(t *testing.T) {
	tokenCalls := 0
	signCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case tossAuthTokenPath:
			tokenCalls++
			_, _ = writer.Write([]byte(`{"access_token":"fresh-token"}`))
		case myDataSignVerificationPath:
			signCalls++
			if signCalls == 1 {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			if request.Header.Get("Authorization") != "Bearer fresh-token" {
				t.Fatalf("authorization = %q", request.Header.Get("Authorization"))
			}
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","rsp_msg":"ok","result":"true","user_ci":"refreshed-ci"}`))
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	client := LiveClient{
		Config: RuntimeConfig{Provider: ProviderTossMyData, AuthBaseURL: server.URL, AuthTokenBaseURL: server.URL},
		HTTP:   server.Client(), Resolve: func(context.Context, string, string) (credentialEnvelope, error) {
			return credentialEnvelope{
				Provider: ProviderTossMyData, ClientID: "id", ClientSecret: "secret",
				AuthAccessToken: "stale-token",
			}, nil
		},
	}
	ci, err := client.SignVerificationFromReference(context.Background(), "op://bundle", SignVerificationRequest{
		CertTxID: "cert", TxID: "tx", SignedConsent: "signed", Consent: "consent",
	})
	if err != nil || ci != "refreshed-ci" || tokenCalls != 1 || signCalls != 2 {
		t.Fatalf("ci = %q err = %v token_calls = %d sign_calls = %d", ci, err, tokenCalls, signCalls)
	}
}
