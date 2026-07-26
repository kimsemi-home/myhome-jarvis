package financeconnector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiveClientFetchesMyDataPagesThroughResolvedToken(t *testing.T) {
	requests := 0
	transactionRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests++
		assertMyDataHeaders(t, request, "test-token")
		if request.Header.Get("x-client-id") != "test-client" {
			t.Fatalf("x-client-id = %q", request.Header.Get("x-client-id"))
		}
		if request.Method == http.MethodGet && request.URL.Path == myDataAccountsPath {
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","account_list":[{"account_num":"ACCOUNT","is_consent":true}]}`))
			return
		}
		if request.Method != http.MethodPost || request.URL.Path != myDataTransactionsPath {
			t.Fatalf("request = %s %s", request.Method, request.URL.Path)
		}
		transactionRequests++
		var body map[string]string
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["org_code"] != "ORG" || body["account_num"] != "ACCOUNT" || body["from_date"] != "20260701" || body["to_date"] != "20260726" || body["limit"] != "500" {
			t.Fatalf("request body = %#v", body)
		}
		if transactionRequests == 1 {
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","next_page":"2","trans_list":[{"trans_dtime":"20260726090000","trans_no":"1","trans_type":"03","currency_code":"KRW","trans_amt":4500000,"trans_memo":"Salary"}]}`))
			return
		}
		_, _ = writer.Write([]byte(`{"rsp_code":"00000","next_page":"","trans_list":[{"trans_dtime":"20260726100000","trans_no":"2","trans_type":"02","currency_code":"KRW","trans_amt":"42000","trans_memo":"Cafe"}]}`))
	}))
	defer server.Close()

	client := LiveClient{
		Config: RuntimeConfig{
			Mode: ModeMyData, Provider: ProviderTossMyData, ExternalCallsLive: true,
			BaseURL: server.URL, FromDate: "20260701", ToDate: "20260726",
			Connections: []ConnectionConfig{{Owner: "spouse", Credential: CredentialBinding{OnePasswordRef: "op://vault/item/credential"}, OrgCode: "ORG", AccountNumber: "ACCOUNT"}},
		},
		Resolve: func(_ context.Context, reference string, provider string) (credentialEnvelope, error) {
			if reference != "op://vault/item/credential" || provider != ProviderTossMyData {
				t.Fatalf("reference = %q", reference)
			}
			return credentialEnvelope{Provider: ProviderTossMyData, ClientID: "test-client", DataAccessToken: "test-token"}, nil
		},
	}
	transactions, err := client.Fetch(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(transactions) != 2 || transactions[0].Amount != 4500000 || transactions[1].Direction != "debit" {
		t.Fatalf("transactions = %#v", transactions)
	}
	report := AssessLiveParity(transactions)
	if requests != 3 || transactions[0].Category != "uncategorized" || !report.Pass || report.ParityPercent < 95 {
		t.Fatalf("requests = %d transactions = %#v", requests, transactions)
	}
}
