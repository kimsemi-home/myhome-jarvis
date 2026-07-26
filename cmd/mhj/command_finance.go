package main

import (
	"context"
	"errors"

	"github.com/kimsemi-home/myhome-jarvis/internal/financeconnector"
)

func runFinance(root string, args []string) error {
	if len(args) != 1 || args[0] != "parity" {
		return errors.New("usage: mhj finance parity")
	}
	loaded, err := financeconnector.LoadConfigured(context.Background(), root)
	if err != nil {
		return err
	}
	report := loaded.Parity
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.Pass {
		return errors.New("finance connector parity below threshold")
	}
	return nil
}
