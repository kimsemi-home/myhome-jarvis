package financeconnector

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreflightValidatesAtomicBundlesWithoutReturningSecrets(t *testing.T) {
	root := consentRootWithActiveLedger(t)
	opDir := t.TempDir()
	bundle := `{"schema":"myhome.finance.credential/v1","provider":"toss_mydata","client_id":"client","client_secret":"secret","data_access_token":"data-token"}`
	opPath := filepath.Join(opDir, "op")
	if err := os.WriteFile(opPath, []byte("#!/bin/sh\nprintf '%s' '"+bundle+"'\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", opDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	config := RuntimeConfig{
		Mode: ModeMyData, Provider: ProviderTossMyData, BaseURL: "https://test.invalid",
		AuthBaseURL: "https://auth.invalid", AuthTokenBaseURL: "https://oauth.invalid",
		ExternalCallsLive: true, Connections: []ConnectionConfig{
			{Owner: "user", Credential: CredentialBinding{OnePasswordRef: "op://vault/user"}},
			{Owner: "spouse", Credential: CredentialBinding{OnePasswordRef: "op://vault/spouse"}},
		},
	}
	report, err := Preflight(context.Background(), root, config)
	if err != nil || !report.Ready || !report.OnePasswordSession || len(report.Connections) != 2 {
		t.Fatalf("report = %#v err = %v", report, err)
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"op://", "client", "secret", "data-token"} {
		if strings.Contains(string(encoded), forbidden) {
			t.Fatalf("preflight leaked %q: %s", forbidden, encoded)
		}
	}
}
