package financeconnector

import (
	"strings"
	"testing"
)

func TestHTTPClientRejectsPartialMTLSBundle(t *testing.T) {
	client := LiveClient{}
	_, err := client.httpClientForCredential(credentialEnvelope{MTLSCertificatePEM: "certificate"})
	if err == nil || !strings.Contains(err.Error(), "mTLS certificate and private key") {
		t.Fatalf("err = %v", err)
	}
}
