package audit

import "strings"

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
