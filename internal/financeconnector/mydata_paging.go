package financeconnector

import (
	"context"
	"fmt"
)

func (client LiveClient) fetchConnection(ctx context.Context, connection ConnectionConfig, token string) ([]SourceTransaction, error) {
	var all []SourceTransaction
	nextPage := ""
	for page := 0; page < 100; page++ {
		response, err := client.requestPage(ctx, connection, token, nextPage)
		if err != nil {
			return nil, err
		}
		transactions, err := response.transactions(connection, client.Config.Provider)
		if err != nil {
			return nil, err
		}
		all = append(all, transactions...)
		if response.NextPage == "" {
			return all, nil
		}
		nextPage = response.NextPage
	}
	return nil, fmt.Errorf("mydata pagination exceeded 100 pages")
}
