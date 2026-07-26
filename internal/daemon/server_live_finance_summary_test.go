package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestLiveFinanceSummaryReachesDaemonBoundary(t *testing.T) {
	provider := httptest.NewServer(http.HandlerFunc(daemonLiveProvider))
	defer provider.Close()
	configureDaemonLiveEnvironment(t, provider.URL)
	root := t.TempDir()
	for _, file := range []string{
		"generated/finance_consent.generated.json",
		"generated/storage.generated.json",
		"fixtures/commerce_purchases.jsonl",
	} {
		copyTestFile(t, repoRoot(t), root, file)
	}
	writeDaemonTestFile(t, root, "data/private/finance/consent.jsonl", daemonActiveLedger)

	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/domain/summary", nil)
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()
	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	for _, expected := range []string{
		`"connector_fixture_only": false`, `"connector_parity_percent": 100`,
		`"credit_minor_units": 4513000`, `"debit_minor_units": 84000`,
		`"finance_net_minor_units": 4429000`, `"card_debit_minor_units": 42000`,
		`"owner": "spouse"`,
	} {
		if !bytes.Contains(recorder.Body.Bytes(), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, recorder.Body.String())
		}
	}
	for _, forbidden := range []string{"replay-token", "ACCOUNT", "CARD-1", "CARD-2"} {
		if bytes.Contains(recorder.Body.Bytes(), []byte(forbidden)) {
			t.Fatalf("live secret material leaked: %s", recorder.Body.String())
		}
	}
}

func configureDaemonLiveEnvironment(t *testing.T, baseURL string) {
	t.Helper()
	opDir := t.TempDir()
	opPath := filepath.Join(opDir, "op")
	if err := os.WriteFile(opPath, []byte("#!/bin/sh\nprintf '%s' replay-token\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", opDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("MYHOME_FINANCE_MODE", "mydata")
	t.Setenv("MYHOME_FINANCE_ALLOW_EXTERNAL", "true")
	t.Setenv("MYHOME_MYDATA_BASE_URL", baseURL)
	t.Setenv("MYHOME_MYDATA_INCLUDE_CARDS", "true")
	t.Setenv("MYHOME_FINANCE_CONNECTIONS_JSON", `[{"owner":"spouse","onepassword_ref":"op://vault/item/token","org_code":"ORG","account_num":"ACCOUNT"}]`)
}
