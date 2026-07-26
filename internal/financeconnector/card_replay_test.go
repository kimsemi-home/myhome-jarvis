package financeconnector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCommittedCardReplayMeetsCanonicalContract(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "finance_toss_mydata_card_approval_response.json"))
	if err != nil {
		t.Fatal(err)
	}
	var response myDataApprovalResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	var transactions []SourceTransaction
	for _, approval := range response.ApprovedList {
		transaction, err := approval.transaction(ConnectionConfig{Owner: "spouse"}, "card-replay-semmi", "domestic", ProviderTossMyData)
		if err != nil {
			t.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}
	report := AssessLiveParity(transactions)
	if len(transactions) != 2 || !report.Pass || report.ParityPercent != 100 {
		t.Fatalf("transactions = %#v report = %#v", transactions, report)
	}
}
