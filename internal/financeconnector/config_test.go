package financeconnector

import "testing"

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
