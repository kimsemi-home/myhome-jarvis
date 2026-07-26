package financeconnector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiveClientFetchesConsentedCardApprovals(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			t.Fatalf("method = %s", request.Method)
		}
		assertMyDataHeaders(t, request, "token")
		if request.URL.Path == myDataCardsPath {
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","card_list":[{"card_id":"card-1","is_consent":true}]}`))
			return
		}
		if request.URL.Path == myDataCardsPath+"/card-1/approval-domestic" {
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","approved_list":[{"approved_num":"a1","approved_dtime":"20260726142000","status":"01","merchant_name":"Cafe","approved_amt":"42000"}]}`))
			return
		}
		if request.URL.Path == myDataCardsPath+"/card-1/approval-overseas" {
			_, _ = writer.Write([]byte(`{"rsp_code":"00000","approved_list":[{"approved_num":"a2","approved_dtime":"20260726150000","status":"02","merchant_name":"Overseas","approved_amt":"10000","currency_code":"USD","krw_amt":"13000"}]}`))
			return
		}
		http.NotFound(writer, request)
	}))
	defer server.Close()
	client := LiveClient{Config: RuntimeConfig{
		Mode: ModeMyData, Provider: ProviderTossMyData, ExternalCallsLive: true,
		BaseURL: server.URL, FromDate: "20260701", ToDate: "20260726",
		IncludeCards: true, Connections: []ConnectionConfig{{Owner: "spouse", OrgCode: "ORG"}},
	}, HTTP: server.Client(), Resolve: func(context.Context, string) (string, error) {
		return "token", nil
	}}
	transactions, err := client.fetchCardTransactions(context.Background(), client.Config.Connections[0], "token")
	if err != nil {
		t.Fatal(err)
	}
	if len(transactions) != 2 || transactions[0].CardID != "card-1" || transactions[1].Direction != "credit" || transactions[1].Amount != 13000 {
		t.Fatalf("transactions = %#v", transactions)
	}
	if report := AssessLiveParity(transactions); !report.Pass || report.ParityPercent != 100 {
		t.Fatalf("parity = %#v", report)
	}
}
