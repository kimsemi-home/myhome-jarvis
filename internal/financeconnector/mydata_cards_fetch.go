package financeconnector

import (
	"context"
	"fmt"
)

func (client LiveClient) fetchCardTransactions(
	ctx context.Context,
	connection ConnectionConfig,
	token string,
) ([]SourceTransaction, error) {
	consentedCards, err := client.discoverCards(ctx, connection, token)
	if err != nil {
		return nil, err
	}
	cardIDs := consentedCards
	if len(connection.CardIDs) > 0 {
		cardIDs = nil
		for _, requested := range connection.CardIDs {
			for _, consented := range consentedCards {
				if requested == consented {
					cardIDs = append(cardIDs, requested)
					break
				}
			}
		}
		if len(cardIDs) != len(connection.CardIDs) {
			return nil, fmt.Errorf("configured card is not consented")
		}
	}
	var all []SourceTransaction
	for _, cardID := range cardIDs {
		for _, target := range []struct {
			path  string
			scope string
		}{{"approval-domestic", "domestic"}, {"approval-overseas", "overseas"}} {
			transactions, err := client.fetchCardApprovals(
				ctx, connection, token, cardID, target.path, target.scope,
			)
			if err != nil {
				return nil, err
			}
			all = append(all, transactions...)
		}
	}
	return all, nil
}
