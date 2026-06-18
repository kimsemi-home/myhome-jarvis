package connectors

import (
	"sort"
	"time"
)

const connectorStatusMessage = "Connector readiness is fixture-only; no credentials, cookies, external API calls, or external actions are enabled."

func StatusForRoot(root string) (Status, error) {
	policy, err := readGeneratedPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	for _, connector := range policy.Connectors {
		clean, err := sanitizeConnector(connector)
		if err != nil {
			return Status{}, err
		}
		status.addConnector(clean)
	}
	sort.Slice(status.Connectors, func(i, j int) bool {
		return status.Connectors[i].Key < status.Connectors[j].Key
	})
	status.Message = connectorStatusMessage
	return status, nil
}

func newStatus(policy generatedPolicy) Status {
	return Status{
		FixtureOnly:             policy.FixtureOnly,
		RealCredentialsAllowed:  policy.RealCredentialsAllowed,
		ExternalAPICallsAllowed: policy.ExternalAPICallsAllowed,
		GeneratedPath:           generatedConnectorPath,
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}

func (status *Status) addConnector(connector Connector) {
	status.Connectors = append(status.Connectors, connector)
	status.ConnectorCount++
	if connector.Status == "planned" {
		status.PlannedCount++
	}
	if connector.FixtureMode {
		status.FixtureModeCount++
	}
	status.ReadOnlyOperationCount += countReadOnlyOperations(connector.AllowedOperations)
	status.ForbiddenOperationCount += len(connector.ForbiddenOperations)
}
