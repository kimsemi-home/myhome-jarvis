package financeconnector

import (
	"encoding/json"
	"testing"
)

func TestCredentialEnvelopeKeepsAuthAndDataTogether(t *testing.T) {
	raw := `{"schema":"myhome.finance.credential/v1","provider":"toss_mydata","client_id":"test-client","client_secret":"test-secret","auth_access_token":"auth-token","data_access_token":"data-token","mtls_certificate_pem":"certificate","mtls_private_key_pem":"private-key"}`
	envelope, err := parseCredentialEnvelope(raw, ProviderTossMyData)
	if err != nil || envelope.ClientID != "test-client" || envelope.DataAccessToken != "data-token" {
		t.Fatalf("envelope = %#v err = %v", envelope, err)
	}
	encoded, err := json.Marshal(envelope)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"schema":"myhome.finance.credential/v1","provider":"toss_mydata"}` {
		t.Fatalf("credential bundle leaked secret fields: %s", encoded)
	}
}

func TestCredentialEnvelopeRejectsUnbundledToken(t *testing.T) {
	if _, err := parseCredentialEnvelope("data-token", ProviderTossMyData); err == nil {
		t.Fatal("plain token must not bypass the credential bundle")
	}
}
