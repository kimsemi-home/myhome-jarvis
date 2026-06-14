package audit

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

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

func AppendCommandIntent(root string, event CommandIntentEvent) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	if strings.TrimSpace(event.At) == "" {
		event.At = time.Now().UTC().Format(time.RFC3339)
	}
	event.Source = normalizeSource(event.Source)
	event.Command = normalizeCommandName(event.Command)
	path := commandIntentPath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = file.Write(data)
	return err
}

func CommandIntentStatusForRoot(root string) (CommandIntentStatus, error) {
	status := CommandIntentStatus{
		Path:      commandIntentRelativePath,
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}
	file, err := os.Open(commandIntentPath(root))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var event CommandIntentEvent
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			return status, err
		}
		status.Exists = true
		status.Count++
		status.Last = &event
	}
	if err := scanner.Err(); err != nil {
		return status, err
	}
	return status, nil
}

func commandIntentPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(commandIntentRelativePath))
}

func normalizeSource(source string) string {
	source = strings.TrimSpace(strings.ToLower(source))
	if source == "" {
		return "unknown"
	}
	return source
}

func normalizeCommandName(command string) string {
	command = strings.ReplaceAll(strings.TrimSpace(strings.ToLower(command)), "-", "_")
	if command == "" {
		return "unknown"
	}
	return command
}

func commandErrorCategory(err error) string {
	message := strings.ToLower(err.Error())
	switch {
	case strings.Contains(message, "invalid payload"), strings.Contains(message, "invalid json"):
		return "invalid_payload"
	case strings.Contains(message, "unknown command"):
		return "unknown_command"
	default:
		return "command_error"
	}
}
