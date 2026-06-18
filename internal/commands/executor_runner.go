package commands

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"
)

func defaultRunner(ctx context.Context, invocation Invocation) Execution {
	started := time.Now()
	execution := Execution{Executed: true, ExitCode: 0}
	cmd := exec.CommandContext(ctx, invocation.Argv[0], invocation.Argv[1:]...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		execution.Error = err.Error()
		execution.ExitCode = -1
		if exitError, ok := err.(*exec.ExitError); ok {
			execution.ExitCode = exitError.ExitCode()
		}
	}
	execution.Output = limitOutput(output.String())
	execution.DurationMillis = time.Since(started).Milliseconds()
	return execution
}

func skippedExecutions(invocations []Invocation, reason string) []Execution {
	executions := make([]Execution, 0, len(invocations))
	for _, invocation := range invocations {
		executions = append(executions, Execution{
			Label:   invocation.Label,
			Argv:    append([]string(nil), invocation.Argv...),
			Skipped: true,
			Error:   reason,
		})
	}
	return executions
}

func limitOutput(output string) string {
	trimmed := strings.TrimSpace(output)
	if len(trimmed) <= 2048 {
		return trimmed
	}
	return trimmed[:2048]
}
