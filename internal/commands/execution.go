package commands

import (
	"context"
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
