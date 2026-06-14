package orchestrator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

type Checkpoint struct {
	At             string               `json:"at"`
	Task           string               `json:"task"`
	LinearStatus   linear.StatusSummary `json:"linear_status"`
	PlannerStatus  planner.Status       `json:"planner_status"`
	SecurityStatus security.Status      `json:"security_status"`
	Result         string               `json:"result"`
	Next           string               `json:"next"`
}

func WriteCheckpoint(root string, checkpoint Checkpoint) (string, error) {
	dir := filepath.Join(root, "data", "private", "checkpoints")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	now := time.Now().UTC()
	if checkpoint.At == "" {
		checkpoint.At = now.Format(time.RFC3339)
	}
	data, err := json.MarshalIndent(checkpoint, "", "  ")
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	name := now.Format("20060102T150405.000000000Z") + ".json"
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return "", err
	}
	return path, nil
}
