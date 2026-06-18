package connectors

import (
	"fmt"
	"strings"
)

func sanitizeConnector(connector Connector) (Connector, error) {
	connector.Key = normalizeToken(connector.Key)
	connector.Label = strings.TrimSpace(connector.Label)
	connector.Category = normalizeToken(connector.Category)
	connector.Status = normalizeToken(connector.Status)
	connector.NextStep = strings.TrimSpace(connector.NextStep)
	connector.DataClasses = normalizeList(connector.DataClasses)
	connector.AllowedOperations = normalizeList(connector.AllowedOperations)
	connector.ForbiddenOperations = normalizeList(connector.ForbiddenOperations)
	if connector.Key == "" {
		return Connector{}, fmt.Errorf("connector key is required")
	}
	if connector.Label == "" {
		return Connector{}, fmt.Errorf("connector %q label is required", connector.Key)
	}
	if connector.Category == "" {
		return Connector{}, fmt.Errorf("connector %q category is required", connector.Key)
	}
	if connector.Status == "" {
		return Connector{}, fmt.Errorf("connector %q status is required", connector.Key)
	}
	if !connector.FixtureMode {
		return Connector{}, fmt.Errorf("connector %q must stay fixture-only", connector.Key)
	}
	for _, operation := range connector.AllowedOperations {
		if isForbiddenAllowedOperation(operation) {
			return Connector{}, fmt.Errorf("connector %q exposes forbidden allowed operation %q", connector.Key, operation)
		}
	}
	return connector, nil
}

func isForbiddenAllowedOperation(operation string) bool {
	switch operation {
	case "external_api_call", "credential_request", "cookie_import", "scraping":
		return true
	case "transfer", "trade", "purchase", "payment":
		return true
	default:
		return false
	}
}
