package audit

const commandIntentRelativePath = "data/private/audit/command-intents.jsonl"

type CommandIntentEvent struct {
	At               string `json:"at"`
	Source           string `json:"source"`
	Command          string `json:"command"`
	DryRun           bool   `json:"dry_run"`
	ExecuteRequested bool   `json:"execute_requested"`
	ExecuteAllowed   bool   `json:"execute_allowed"`
	InvocationCount  int    `json:"invocation_count"`
	WarningCount     int    `json:"warning_count"`
	Success          bool   `json:"success"`
	ErrorCategory    string `json:"error_category,omitempty"`
}

type CommandIntentStatus struct {
	Path      string              `json:"path"`
	Exists    bool                `json:"exists"`
	Count     int                 `json:"count"`
	Last      *CommandIntentEvent `json:"last,omitempty"`
	CheckedAt string              `json:"checked_at"`
}
