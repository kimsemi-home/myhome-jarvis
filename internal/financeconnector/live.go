package financeconnector

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type LiveClient struct {
	Config  RuntimeConfig
	Resolve CredentialResolver
	HTTP    *http.Client
}

func (client LiveClient) Fetch(ctx context.Context) ([]SourceTransaction, error) {
	if client.Config.Mode != ModeMyData {
		return nil, fmt.Errorf("live client requires mydata mode")
	}
	if !client.Config.ExternalCallsLive {
		return nil, fmt.Errorf("external finance calls are disabled")
	}
	if strings.TrimSpace(client.Config.BaseURL) == "" {
		return nil, fmt.Errorf("mydata base URL is required")
	}
	if client.Resolve == nil {
		client.Resolve = ResolveOnePasswordCLI
	}
	if client.HTTP == nil {
		client.HTTP = &http.Client{Timeout: 15 * time.Second}
	}
	var all []SourceTransaction
	for _, connection := range client.Config.Connections {
		credential, err := client.Resolve(ctx, connection.Credential.OnePasswordRef, client.Config.Provider)
		if err != nil {
			return nil, err
		}
		if credential.Provider != client.Config.Provider {
			return nil, fmt.Errorf("credential bundle provider mismatch")
		}
		transactions, err := client.fetchConnection(ctx, connection, credential)
		if err != nil {
			return nil, err
		}
		all = append(all, transactions...)
		if client.Config.IncludeCards {
			cardTransactions, err := client.fetchCardTransactions(ctx, connection, credential)
			if err != nil {
				return nil, err
			}
			all = append(all, cardTransactions...)
		}
	}
	if len(all) == 0 {
		return nil, fmt.Errorf("mydata returned no transactions")
	}
	return all, nil
}
