package financeconnector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCommittedCardListReplayKeepsOnlyConsentedCards(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "fixtures", "finance_toss_mydata_cards_response.json"))
	if err != nil {
		t.Fatal(err)
	}
	var response myDataCardListResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	if len(response.CardList) != 2 || !response.CardList[0].IsConsent || response.CardList[1].IsConsent {
		t.Fatalf("card list = %#v", response.CardList)
	}
}
