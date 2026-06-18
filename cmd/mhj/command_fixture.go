package main

import (
	"errors"
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func runHarness(root string, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: mhj harness <home|finance|commerce>")
	}
	var report commands.HarnessReport
	switch args[0] {
	case "home":
		report = commands.RunHomeHarness()
	case "finance":
		report = commands.RunFinanceHarness(root)
	case "commerce":
		report = commands.RunCommerceHarness(root)
	default:
		return errors.New("usage: mhj harness <home|finance|commerce>")
	}
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.Passed {
		return fmt.Errorf("%s harness failed", report.Name)
	}
	return nil
}
