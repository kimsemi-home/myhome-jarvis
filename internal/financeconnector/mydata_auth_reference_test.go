package financeconnector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignVerificationFromReferenceKeepsProviderAtomic(t *testing.T) {
	client := LiveClient{
		Config: RuntimeConfig{Provider: ProviderTossMyData, AuthBaseURL: "https://example.invalid"},
		Resolve: func(_ context.Context, reference, provider string) (credentialEnvelope, error) {
			if reference != "op://vault/item/bundle" || provider != ProviderTossMyData {
				t.Fatalf("reference = %q provider = %q", reference, provider)
			}
			return credentialEnvelope{Provider: "other", AuthAccessToken: "token"}, nil
		},
	}
	_, err := client.SignVerificationFromReference(context.Background(), "op://vault/item/bundle", SignVerificationRequest{})
	if err == nil || !strings.Contains(err.Error(), "provider mismatch") {
		t.Fatalf("err = %v", err)
	}
}

func TestSignVerificationFromReferenceRefreshesTokenInMemory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case tossAuthTokenPath:
			_, _ = writer.Write([]byte(`{"access_token":"issued-token"}`))
		case myDataSignVerificationPath:
			if request.Header.Get("Authorization") != "Bearer issued-token" {
				t.Fatalf("authorization = %q", request.Header.Get("Authorization"))
			}
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","rsp_msg":"ok","result":"true","user_ci":"test-ci"}`))
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	client := LiveClient{
		Config: RuntimeConfig{Provider: ProviderTossMyData, AuthBaseURL: server.URL, AuthTokenBaseURL: server.URL},
		HTTP:   server.Client(), Resolve: func(context.Context, string, string) (credentialEnvelope, error) {
			return credentialEnvelope{Provider: ProviderTossMyData, ClientID: "id", ClientSecret: "secret"}, nil
		},
	}
	ci, err := client.SignVerificationFromReference(context.Background(), "op://bundle", SignVerificationRequest{
		CertTxID: "cert", TxID: "tx", SignedConsent: "signed", Consent: "consent",
	})
	if err != nil || ci != "test-ci" {
		t.Fatalf("ci = %q err = %v", ci, err)
	}
}
