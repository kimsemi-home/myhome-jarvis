package daemon

import (
	"context"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func (server *Server) buildSummary(ctx context.Context) (domain.Summary, error) {
	loaded, err := financeconnector.LoadConfigured(ctx, server.config.Root)
	if err != nil {
		return domain.Summary{}, err
	}
	return domain.BuildSummaryWithFinance(
		server.config.Root, loaded.Transactions, loaded.Parity,
	)
}
