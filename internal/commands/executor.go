package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Execution struct {
	Label          string   `json:"label"`
	Argv           []string `json:"argv"`
	Executed       bool     `json:"executed"`
	Skipped        bool     `json:"skipped,omitempty"`
	ExitCode       int      `json:"exit_code,omitempty"`
	Output         string   `json:"output,omitempty"`
	Error          string   `json:"error,omitempty"`
	DurationMillis int64    `json:"duration_millis,omitempty"`
}

type Runner func(context.Context, Invocation) Execution

type ExecuteOptions struct {
	Platform string
	Timeout  time.Duration
	Runner   Runner
}

func Execute(ctx context.Context, plan Plan, options ExecuteOptions) (Plan, error) {
	if !plan.ExecuteAllowed {
		plan.Warnings = append(plan.Warnings, "execution skipped because execute was not explicitly allowed")
		return plan, nil
	}
	platform := options.Platform
	if strings.TrimSpace(platform) == "" {
		platform = runtime.GOOS
	}
	if platform != "darwin" {
		plan.Warnings = append(plan.Warnings, "execution skipped because macOS is required")
		plan.Executions = skippedExecutions(plan.Invocations, "requires darwin platform")
		return plan, nil
	}
	for _, invocation := range plan.Invocations {
		if err := validateInvocation(invocation); err != nil {
			return Plan{}, err
		}
	}
	runner := options.Runner
	if runner == nil {
		runner = defaultRunner
	}
	timeout := options.Timeout
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	plan.DryRun = false
	for _, invocation := range plan.Invocations {
		runCtx, cancel := context.WithTimeout(ctx, timeout)
		execution := runner(runCtx, invocation)
		cancel()
		execution.Label = invocation.Label
		execution.Argv = append([]string(nil), invocation.Argv...)
		plan.Executions = append(plan.Executions, execution)
	}
	return plan, nil
}

func validateInvocation(invocation Invocation) error {
	if len(invocation.Argv) == 0 {
		return errors.New("execution requires an argv plan")
	}
	executable := filepath.Base(invocation.Argv[0])
	switch executable {
	case "open", "osascript", "pmset":
	default:
		return fmt.Errorf("executable %q is not allowed", executable)
	}
	for _, arg := range invocation.Argv {
		if strings.ContainsRune(arg, 0) {
			return errors.New("argv contains NUL byte")
		}
	}
	return nil
}

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
