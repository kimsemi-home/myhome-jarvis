package financeconnector

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeConfigDefaultsToFixtureMode(t *testing.T) {
	config, err := RuntimeConfigFromEnv(func(string) string { return "" })
	if err != nil {
		t.Fatal(err)
	}
	if config.Mode != ModeFixture || config.ExternalCallsLive || config.OnePasswordRef != "" {
		t.Fatalf("config = %#v", config)
	}
}

func TestRuntimeConfigRequiresOnePasswordReferenceForMyData(t *testing.T) {
	config, err := RuntimeConfigFromEnv(func(name string) string {
		if name == "MYHOME_FINANCE_MODE" {
			return ModeMyData
		}
		return ""
	})
	if err == nil || config.Mode != "" {
		t.Fatalf("expected missing op reference error, config = %#v", config)
	}

	config, err = RuntimeConfigFromEnv(func(name string) string {
		switch name {
		case "MYHOME_FINANCE_MODE":
			return ModeMyData
		case "MYHOME_FINANCE_1PASSWORD_REF":
			return "op://MyHome-Jarvis/Finance-Connector/api"
		case "MYHOME_FINANCE_ALLOW_EXTERNAL":
			return "true"
		default:
			return ""
		}
	})
	if err != nil || !config.LiveModeRequested || !config.ExternalCallsLive || config.OnePasswordRef == "" {
		t.Fatalf("config = %#v err = %v", config, err)
	}
}

func TestRuntimeConfigParsesConnectionReferencesWithoutSerializingSecrets(t *testing.T) {
	config, err := RuntimeConfigFromEnv(func(name string) string {
		values := map[string]string{
			"MYHOME_FINANCE_MODE":             ModeMyData,
			"MYHOME_FINANCE_ALLOW_EXTERNAL":   "true",
			"MYHOME_FINANCE_CONNECTIONS_JSON": `[{"owner":"spouse","onepassword_ref":"op://vault/item/token","org_code":"ORG","account_num":"ACCOUNT"}]`,
		}
		return values[name]
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(config.Connections) != 1 || config.Connections[0].OrgCode != "ORG" ||
		config.Connections[0].AccountNumber != "ACCOUNT" {
		t.Fatalf("connections = %#v", config.Connections)
	}
	encoded, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "op://") || strings.Contains(string(encoded), "ACCOUNT") {
		t.Fatalf("runtime config serialized secret material: %s", encoded)
	}
}
