package financeconnector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfiguredLiveReplaysConsentedBankData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(configuredReplayHandler))
	defer server.Close()
	configureReplayEnvironment(t, server.URL)
	root := consentRootWithActiveLedger(t)

	loaded, err := LoadConfigured(context.Background(), root)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Transactions) != 2 || !loaded.Parity.Pass || loaded.Parity.ParityPercent < 95 {
		t.Fatalf("loaded = %#v", loaded)
	}
	if loaded.Transactions[0].Owner != "spouse" || loaded.Transactions[1].MerchantName == "" {
		t.Fatalf("transactions = %#v", loaded.Transactions)
	}
}

func configureReplayEnvironment(t *testing.T, baseURL string) {
	t.Helper()
	opDir := t.TempDir()
	opPath := filepath.Join(opDir, "op")
	if err := os.WriteFile(opPath, []byte("#!/bin/sh\nprintf '%s' replay-token\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", opDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("MYHOME_FINANCE_MODE", ModeMyData)
	t.Setenv("MYHOME_FINANCE_ALLOW_EXTERNAL", "true")
	t.Setenv("MYHOME_MYDATA_PROVIDER", ProviderTossMyData)
	t.Setenv("MYHOME_MYDATA_BASE_URL", baseURL)
	t.Setenv("MYHOME_MYDATA_INCLUDE_CARDS", "false")
	t.Setenv("MYHOME_FINANCE_CONNECTIONS_JSON", `[{"owner":"spouse","onepassword_ref":"op://vault/item/token","org_code":"ORG","account_num":"ACCOUNT"}]`)
}

func configuredReplayHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == myDataAccountsPath {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","account_list":[{"account_num":"ACCOUNT","is_consent":true}]}`))
		return
	}
	if request.URL.Path == myDataTransactionsPath {
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","trans_list":[{"trans_dtime":"20260726090000","trans_no":"income-1","trans_type":"02","currency_code":"KRW","trans_amt":"4500000","trans_memo":"Salary"},{"trans_dtime":"20260726100000","trans_no":"spend-1","trans_type":"03","currency_code":"KRW","trans_amt":42000,"trans_memo":"Cafe"}]}`))
		return
	}
	http.NotFound(writer, request)
}
