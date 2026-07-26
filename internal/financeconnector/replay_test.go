package financeconnector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCommittedMyDataResponseReplayMeetsLiveContract(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "finance_toss_mydata_response.json"))
	if err != nil {
		t.Fatal(err)
	}
	var response myDataResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	transactions, err := response.transactions(ConnectionConfig{Owner: "user"}, ProviderTossMyData)
	if err != nil {
		t.Fatal(err)
	}
	report := AssessLiveParity(transactions)
	if len(transactions) != 2 || !report.Pass || report.ParityPercent != 100 {
		t.Fatalf("transactions = %#v report = %#v", transactions, report)
	}
}
