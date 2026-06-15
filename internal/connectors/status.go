package connectors

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const generatedConnectorPath = "generated/connectors.generated.json"

type Status struct {
	FixtureOnly             bool        `json:"fixture_only"`
	RealCredentialsAllowed  bool        `json:"real_credentials_allowed"`
	ExternalAPICallsAllowed bool        `json:"external_api_calls_allowed"`
	ConnectorCount          int         `json:"connector_count"`
	PlannedCount            int         `json:"planned_count"`
	FixtureModeCount        int         `json:"fixture_mode_count"`
	ReadOnlyOperationCount  int         `json:"read_only_operation_count"`
	ForbiddenOperationCount int         `json:"forbidden_operation_count"`
	GeneratedPath           string      `json:"generated_path"`
	Connectors              []Connector `json:"connectors"`
	Message                 string      `json:"message"`
	CheckedAt               string      `json:"checked_at"`
}

type Connector struct {
	Key                 string   `json:"key"`
	Label               string   `json:"label"`
	Category            string   `json:"category"`
	Status              string   `json:"status"`
	FixtureMode         bool     `json:"fixture_mode"`
	DataClasses         []string `json:"data_classes"`
	AllowedOperations   []string `json:"allowed_operations"`
	ForbiddenOperations []string `json:"forbidden_operations"`
	NextStep            string   `json:"next_step"`
}

type generatedPolicy struct {
	FixtureOnly             bool        `json:"fixture_only"`
	RealCredentialsAllowed  bool        `json:"real_credentials_allowed"`
	ExternalAPICallsAllowed bool        `json:"external_api_calls_allowed"`
	Connectors              []Connector `json:"connectors"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := readGeneratedPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		FixtureOnly:             policy.FixtureOnly,
		RealCredentialsAllowed:  policy.RealCredentialsAllowed,
		ExternalAPICallsAllowed: policy.ExternalAPICallsAllowed,
		GeneratedPath:           generatedConnectorPath,
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
	for _, connector := range policy.Connectors {
		clean, err := sanitizeConnector(connector)
		if err != nil {
			return Status{}, err
		}
		status.Connectors = append(status.Connectors, clean)
		status.ConnectorCount++
		if clean.Status == "planned" {
			status.PlannedCount++
		}
		if clean.FixtureMode {
			status.FixtureModeCount++
		}
		status.ReadOnlyOperationCount += countReadOnlyOperations(clean.AllowedOperations)
		status.ForbiddenOperationCount += len(clean.ForbiddenOperations)
	}
	sort.Slice(status.Connectors, func(i, j int) bool {
		return status.Connectors[i].Key < status.Connectors[j].Key
	})
	status.Message = "Connector readiness is fixture-only; no credentials, cookies, external API calls, or external actions are enabled."
	return status, nil
}

func readGeneratedPolicy(root string) (generatedPolicy, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(generatedConnectorPath)))
	if err != nil {
		return generatedPolicy{}, err
	}
	var policy generatedPolicy
	if err := json.Unmarshal(body, &policy); err != nil {
		return generatedPolicy{}, err
	}
	return policy, nil
}

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
		switch operation {
		case "external_api_call", "credential_request", "cookie_import", "scraping", "transfer", "trade", "purchase", "payment":
			return Connector{}, fmt.Errorf("connector %q exposes forbidden allowed operation %q", connector.Key, operation)
		}
	}
	return connector, nil
}

func normalizeToken(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func normalizeList(values []string) []string {
	seen := map[string]bool{}
	var normalized []string
	for _, value := range values {
		item := normalizeToken(value)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		normalized = append(normalized, item)
	}
	sort.Strings(normalized)
	return normalized
}

func countReadOnlyOperations(operations []string) int {
	count := 0
	for _, operation := range operations {
		if operation == "read_fixture" || operation == "summarize" || operation == "recommend_review" {
			count++
		}
	}
	return count
}
