package financeconnector

import (
	"encoding/json"
	"fmt"
	"strings"
)

const credentialSchemaV1 = "myhome.finance.credential/v1"

// credentialEnvelope is resolved once per connection and passed unchanged to
// every provider request. It is deliberately private so it cannot cross the
// daemon/UI or public-status boundary by accident.
type credentialEnvelope struct {
	Schema             string `json:"schema"`
	Provider           string `json:"provider"`
	ClientID           string `json:"client_id"`
	ClientSecret       string `json:"client_secret"`
	AuthAccessToken    string `json:"auth_access_token"`
	DataAccessToken    string `json:"data_access_token"`
	MTLSCertificatePEM string `json:"mtls_certificate_pem"`
	MTLSPrivateKeyPEM  string `json:"mtls_private_key_pem"`
}

func (envelope credentialEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Schema   string `json:"schema"`
		Provider string `json:"provider"`
	}{Schema: envelope.Schema, Provider: envelope.Provider})
}

func parseCredentialEnvelope(raw string, provider string) (credentialEnvelope, error) {
	var envelope credentialEnvelope
	if err := json.Unmarshal([]byte(raw), &envelope); err != nil {
		return credentialEnvelope{}, fmt.Errorf("invalid credential bundle")
	}
	if envelope.Schema != credentialSchemaV1 {
		return credentialEnvelope{}, fmt.Errorf("unsupported credential bundle schema")
	}
	if envelope.Provider != provider {
		return credentialEnvelope{}, fmt.Errorf("credential bundle provider mismatch")
	}
	if strings.TrimSpace(envelope.DataAccessToken) == "" {
		return credentialEnvelope{}, fmt.Errorf("credential bundle data access token is missing")
	}
	return envelope, nil
}
