package connectors

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusForRootReadsPublicSafeGeneratedCatalog(t *testing.T) {
	root := repoRoot(t)
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}

	if !status.FixtureOnly || status.RealCredentialsAllowed || status.ExternalAPICallsAllowed {
		t.Fatalf("connector safety flags = %#v", status)
	}
	if status.ConnectorCount < 6 || status.PlannedCount != status.ConnectorCount {
		t.Fatalf("connector counts = %#v", status)
	}
	if status.FixtureModeCount != status.ConnectorCount {
		t.Fatalf("fixture count = %d connectors = %d", status.FixtureModeCount, status.ConnectorCount)
	}
	if status.GeneratedPath != generatedConnectorPath {
		t.Fatalf("generated path = %q", status.GeneratedPath)
	}
	for _, connector := range status.Connectors {
		for _, forbidden := range []string{"token", "secret", "cookie_value", "account_id", "card_number"} {
			joined := strings.Join([]string{
				connector.Key,
				connector.Label,
				connector.Category,
				connector.Status,
				strings.Join(connector.DataClasses, " "),
				strings.Join(connector.AllowedOperations, " "),
				connector.NextStep,
			}, " ")
			if strings.Contains(strings.ToLower(joined), forbidden) {
				t.Fatalf("connector %q leaked forbidden marker %q in %q", connector.Key, forbidden, joined)
			}
		}
		for _, operation := range connector.AllowedOperations {
			switch operation {
			case "credential_request", "external_api_call", "cookie_import", "scraping", "transfer", "trade", "purchase", "payment":
				t.Fatalf("connector %q allowed forbidden operation %q", connector.Key, operation)
			}
		}
	}
}

func TestStatusRejectsUnsafeAllowedOperation(t *testing.T) {
	root := t.TempDir()
	generated := filepath.Join(root, filepath.FromSlash(generatedConnectorPath))
	if err := os.MkdirAll(filepath.Dir(generated), 0o755); err != nil {
		t.Fatal(err)
	}
	body := `{"fixture_only":true,"real_credentials_allowed":false,"external_api_calls_allowed":false,"connectors":[{"key":"banking","label":"Banking","category":"finance","status":"planned","fixture_mode":true,"data_classes":["transactions"],"allowed_operations":["read_fixture","external_api_call"],"forbidden_operations":["credential_request"],"next_step":"stay local"}]}`
	if err := os.WriteFile(generated, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := StatusForRoot(root); err == nil {
		t.Fatal("expected unsafe connector operation to fail")
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}
