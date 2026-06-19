package linear

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func AppendOfflineEvent(root string, kind string, message string) error {
	return AppendOfflineAction(root, kind, message, nil)
}

func AppendOfflineAction(root string, kind string, message string, payload any) error {
	queuePath := filepath.Join(root, "data", "private", "linear-offline-queue.jsonl")
	if err := os.MkdirAll(filepath.Dir(queuePath), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(queuePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	event := OfflineEvent{
		At:      time.Now().UTC().Format(time.RFC3339),
		Kind:    kind,
		Message: message,
		Synced:  false,
	}
	if payload != nil {
		return json.NewEncoder(file).Encode(struct {
			OfflineEvent
			Payload any `json:"payload"`
		}{OfflineEvent: event, Payload: payload})
	}
	return json.NewEncoder(file).Encode(event)
}
