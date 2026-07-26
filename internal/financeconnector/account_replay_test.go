package financeconnector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCommittedAccountListReplayKeepsOnlyConsentedAccounts(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "finance_toss_mydata_accounts_response.json"))
	if err != nil {
		t.Fatal(err)
	}
	var response myDataAccountListResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	if len(response.AccountList) != 2 || !response.AccountList[0].IsConsent || response.AccountList[1].IsConsent {
		t.Fatalf("account list = %#v", response.AccountList)
	}
}
