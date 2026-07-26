package financeconnector

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseConnections(raw string) ([]ConnectionConfig, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	var connections []ConnectionConfig
	if err := json.Unmarshal([]byte(raw), &connections); err != nil {
		return nil, fmt.Errorf("invalid MYHOME_FINANCE_CONNECTIONS_JSON: %w", err)
	}
	return connections, nil
}

func allConnectionsHaveRefs(connections []ConnectionConfig) bool {
	for _, connection := range connections {
		if !strings.HasPrefix(strings.TrimSpace(connection.OnePasswordRef), "op://") {
			return false
		}
	}
	return true
}
