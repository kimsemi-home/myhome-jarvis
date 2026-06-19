package commands

import "errors"

type Spec struct {
	Name          string   `json:"name"`
	Summary       string   `json:"summary"`
	PayloadFields []string `json:"payload_fields"`
}

type Invocation struct {
	Label string   `json:"label"`
	Argv  []string `json:"argv"`
	URL   string   `json:"url,omitempty"`
}

type Plan struct {
	Name           string       `json:"name"`
	DryRun         bool         `json:"dry_run"`
	ExecuteAllowed bool         `json:"execute_allowed"`
	Invocations    []Invocation `json:"invocations"`
	Executions     []Execution  `json:"executions,omitempty"`
	Warnings       []string     `json:"warnings,omitempty"`
}

var errInvalidPayload = errors.New("invalid payload")

func WithExecuteAllowed(plan Plan, executeAllowed bool) Plan {
	plan.ExecuteAllowed = executeAllowed
	return plan
}
