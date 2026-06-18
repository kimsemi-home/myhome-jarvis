package commands

import (
	"context"
	"runtime"
	"strings"
	"time"
)

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
