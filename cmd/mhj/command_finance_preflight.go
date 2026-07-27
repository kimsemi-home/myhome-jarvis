package main

import (
	"context"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func runFinancePreflight(root string) error {
	config, err := financeconnector.RuntimeConfigFromEnv(os.Getenv)
	if err != nil {
		return err
	}
	report, preflightErr := financeconnector.Preflight(context.Background(), root, config)
	if err := writeJSON(report); err != nil {
		return err
	}
	return preflightErr
}
