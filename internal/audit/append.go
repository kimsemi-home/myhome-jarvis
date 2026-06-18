package audit

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AppendCommandIntent(root string, event CommandIntentEvent) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	event = normalizeCommandIntentEvent(event)
	path := commandIntentPath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	return appendCommandIntentLine(path, event)
}

func normalizeCommandIntentEvent(event CommandIntentEvent) CommandIntentEvent {
	if strings.TrimSpace(event.At) == "" {
		event.At = time.Now().UTC().Format(time.RFC3339)
	}
	event.Source = normalizeSource(event.Source)
	event.Command = normalizeCommandName(event.Command)
	return event
}

func appendCommandIntentLine(path string, event CommandIntentEvent) error {
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
