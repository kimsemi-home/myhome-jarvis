package financeconnector

import (
	"context"
	"strings"
	"testing"
)

func TestSignVerificationRejectsMissingAuthProof(t *testing.T) {
	client := LiveClient{Config: RuntimeConfig{AuthBaseURL: "https://example.invalid"}}
	_, err := client.SignVerification(context.Background(), credentialEnvelope{
		Provider: ProviderTossMyData, DataAccessToken: "data-token",
	}, SignVerificationRequest{CertTxID: "cert-tx", TxID: "mydata-tx"})
	if err == nil || !strings.Contains(err.Error(), "auth access token") {
		t.Fatalf("err = %v", err)
	}
	_, err = client.SignVerification(context.Background(), credentialEnvelope{
		Provider: ProviderTossMyData, AuthAccessToken: "auth-token", DataAccessToken: "data-token",
	}, SignVerificationRequest{CertTxID: "cert-tx", TxID: "mydata-tx"})
	if err == nil || !strings.Contains(err.Error(), "consent proof") {
		t.Fatalf("err = %v", err)
	}
}
