package main

import (
	"context"
	"errors"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func runCommand(root string, args []string) error {
	if len(args) != 2 {
		return errors.New("usage: mhj command <name> '<json-payload>'")
	}
	executeRequested := os.Getenv("MYHOME_EXECUTE") == "true"
	plan, err := commands.Build(args[0], []byte(args[1]))
	if err == nil {
		plan = commands.WithExecuteAllowed(plan, executeRequested)
	}
	if plan.ExecuteAllowed {
		plan, err = commands.Execute(context.Background(), plan, commands.ExecuteOptions{})
	}
	auditIntent := audit.CommandIntentFromPlan("cli", args[0], executeRequested, plan, err)
	if auditErr := audit.AppendCommandIntent(root, auditIntent); err == nil && auditErr != nil {
		return auditErr
	}
	if err != nil {
		return err
	}
	return writeJSON(plan)
}
