package audit

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func CommandIntentFromPlan(source string, requestedCommand string, executeRequested bool, plan commands.Plan, planErr error) CommandIntentEvent {
	command := plan.Name
	if command == "" {
		command = normalizeCommandName(requestedCommand)
	}
	event := CommandIntentEvent{
		At:               time.Now().UTC().Format(time.RFC3339),
		Source:           normalizeSource(source),
		Command:          command,
		DryRun:           plan.DryRun,
		ExecuteRequested: executeRequested,
		ExecuteAllowed:   plan.ExecuteAllowed,
		InvocationCount:  len(plan.Invocations),
		WarningCount:     len(plan.Warnings),
		Success:          planErr == nil,
	}
	if planErr != nil {
		event.ErrorCategory = commandErrorCategory(planErr)
	}
	return event
}
