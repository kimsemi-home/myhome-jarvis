package financeconnector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignVerificationUsesAtomicBundleAndTossContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost || request.URL.Path != myDataSignVerificationPath {
			t.Fatalf("request = %s %s", request.Method, request.URL.Path)
		}
		if request.Header.Get("Authorization") != "Bearer auth-token" {
			t.Fatalf("authorization = %q", request.Header.Get("Authorization"))
		}
		transactionID := request.Header.Get("x-api-tran-id")
		if len(transactionID) != 25 || !strings.HasPrefix(transactionID, "MHJ") {
			t.Fatalf("x-api-tran-id = %q", transactionID)
		}
		var body SignVerificationRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body.CertTxID != "cert-tx" || body.TxID != "mydata-tx" ||
			body.SignedConsent != "signed-proof" || body.Consent != "consent-proof" {
			t.Fatalf("request body = %#v", body)
		}
		_, _ = writer.Write([]byte(`{"tx_id":"mydata-tx","rsp_code":"00000","rsp_msg":"성공","result":"true","user_ci":"test-ci"}`))
	}))
	defer server.Close()

	client := LiveClient{
		Config: RuntimeConfig{AuthBaseURL: server.URL},
		HTTP:   server.Client(),
	}
	ci, err := client.SignVerification(context.Background(), credentialEnvelope{
		Provider:        ProviderTossMyData,
		AuthAccessToken: "auth-token",
		DataAccessToken: "data-token",
	}, SignVerificationRequest{
		CertTxID:         "cert-tx",
		TxID:             "mydata-tx",
		SignedConsentLen: "12",
		SignedConsent:    "signed-proof",
		ConsentType:      "1",
		ConsentLen:       "13",
		Consent:          "consent-proof",
	})
	if err != nil || ci != "test-ci" {
		t.Fatalf("ci = %q err = %v", ci, err)
	}
}
