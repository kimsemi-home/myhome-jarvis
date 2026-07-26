package financeconnector

import (
	"encoding/json"
	"fmt"
	"strings"
)

type connectionInput struct {
	Owner         string   `json:"owner"`
	CredentialRef string   `json:"credential_ref"`
	OrgCode       string   `json:"org_code"`
	AccountNumber string   `json:"account_num"`
	CardIDs       []string `json:"card_ids"`
}

func parseConnections(raw string) ([]ConnectionConfig, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	var inputs []connectionInput
	if err := json.Unmarshal([]byte(raw), &inputs); err != nil {
		return nil, fmt.Errorf("invalid MYHOME_FINANCE_CONNECTIONS_JSON: %w", err)
	}
	connections := make([]ConnectionConfig, 0, len(inputs))
	for _, input := range inputs {
		connections = append(connections, ConnectionConfig{
			Owner: input.Owner, Credential: CredentialBinding{OnePasswordRef: input.CredentialRef},
			OrgCode: input.OrgCode, AccountNumber: input.AccountNumber,
			CardIDs: input.CardIDs,
		})
	}
	return connections, nil
}

func allConnectionsHaveRefs(connections []ConnectionConfig) bool {
	for _, connection := range connections {
		if !strings.HasPrefix(strings.TrimSpace(connection.Credential.OnePasswordRef), "op://") {
			return false
		}
	}
	return true
}
