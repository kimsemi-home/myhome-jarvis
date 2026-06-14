package planner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const generatedRelativePath = "generated/planner.generated.json"

type Policy struct {
	LoopMode              string           `json:"loop_mode"`
	MaxTaskScope          string           `json:"max_task_scope"`
	CheckpointRoot        string           `json:"checkpoint_root"`
	QualityRequired       bool             `json:"quality_required"`
	LinearOfflineFallback bool             `json:"linear_offline_fallback"`
	DefaultNext           string           `json:"default_next"`
	TaskGraph             []Task           `json:"task_graph"`
	LinearTemplates       []LinearTemplate `json:"linear_templates"`
}

type Task struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Owner     string   `json:"owner"`
	Status    string   `json:"status"`
	DependsOn []string `json:"depends_on"`
}

type LinearTemplate struct {
	Name        string   `json:"name"`
	TitlePrefix string   `json:"title_prefix"`
	Labels      []string `json:"labels"`
}

type Status struct {
	LoopMode                  string `json:"loop_mode"`
	TaskCount                 int    `json:"task_count"`
	ReadyCount                int    `json:"ready_count"`
	CompletedCount            int    `json:"completed_count"`
	BlockedExternalWriteCount int    `json:"blocked_external_write_count"`
	NextTask                  *Task  `json:"next_task,omitempty"`
	BlockedExternalWriteTasks []Task `json:"blocked_external_write_tasks,omitempty"`
	LinearTemplateCount       int    `json:"linear_template_count"`
	QualityRequired           bool   `json:"quality_required"`
	LinearOfflineFallback     bool   `json:"linear_offline_fallback"`
	CheckpointRoot            string `json:"checkpoint_root"`
	CheckedAt                 string `json:"checked_at"`
}

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(filepath.Join(root, filepath.FromSlash(generatedRelativePath)))
	if err != nil {
		return Status{}, err
	}
	if err := validatePolicy(policy); err != nil {
		return Status{}, err
	}
	status := Status{
		LoopMode:              policy.LoopMode,
		TaskCount:             len(policy.TaskGraph),
		LinearTemplateCount:   len(policy.LinearTemplates),
		QualityRequired:       policy.QualityRequired,
		LinearOfflineFallback: policy.LinearOfflineFallback,
		CheckpointRoot:        policy.CheckpointRoot,
		CheckedAt:             time.Now().UTC().Format(time.RFC3339),
	}
	taskStatuses := make(map[string]string, len(policy.TaskGraph))
	for _, task := range policy.TaskGraph {
		normalized := normalizeStatus(task.Status)
		taskStatuses[task.ID] = normalized
		switch normalized {
		case "ready":
			status.ReadyCount++
		case "completed":
			status.CompletedCount++
		case "blocked_external_write":
			status.BlockedExternalWriteCount++
			status.BlockedExternalWriteTasks = append(status.BlockedExternalWriteTasks, task)
		}
	}
	for index := range policy.TaskGraph {
		task := policy.TaskGraph[index]
		if normalizeStatus(task.Status) == "ready" && dependenciesSatisfied(task, taskStatuses) {
			status.NextTask = &task
			break
		}
	}
	return status, nil
}

func ReadPolicy(path string) (Policy, error) {
	file, err := os.Open(path)
	if err != nil {
		return Policy{}, err
	}
	defer file.Close()

	var policy Policy
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&policy); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func validatePolicy(policy Policy) error {
	if strings.TrimSpace(policy.LoopMode) == "" {
		return errors.New("planner loop mode is required")
	}
	if len(policy.TaskGraph) == 0 {
		return errors.New("planner task graph is required")
	}
	if err := validateCheckpointRoot(policy.CheckpointRoot); err != nil {
		return err
	}
	ids := map[string]bool{}
	hasExternalWriteBoundary := false
	for _, task := range policy.TaskGraph {
		id := strings.TrimSpace(task.ID)
		if id == "" {
			return errors.New("planner task id is required")
		}
		if ids[id] {
			return fmt.Errorf("duplicate planner task id %q", id)
		}
		ids[id] = true
		if normalizeStatus(task.Status) == "blocked_external_write" {
			hasExternalWriteBoundary = true
		}
		if !validStatus(task.Status) {
			return fmt.Errorf("planner task %q has invalid status %q", task.ID, task.Status)
		}
	}
	if !hasExternalWriteBoundary {
		return errors.New("planner task graph must include an external-write boundary")
	}
	for _, task := range policy.TaskGraph {
		for _, dependency := range task.DependsOn {
			if !ids[dependency] {
				return fmt.Errorf("planner task %q depends on unknown task %q", task.ID, dependency)
			}
		}
	}
	return nil
}

func validateCheckpointRoot(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return errors.New("planner checkpoint root is required")
	}
	native := filepath.FromSlash(value)
	if filepath.IsAbs(native) {
		return errors.New("planner checkpoint root must be repo-relative")
	}
	clean := filepath.ToSlash(filepath.Clean(native))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return errors.New("planner checkpoint root must stay inside the repo")
	}
	if !strings.HasPrefix(clean, "data/private/") {
		return errors.New("planner checkpoint root must stay under data/private")
	}
	return nil
}

func dependenciesSatisfied(task Task, taskStatuses map[string]string) bool {
	for _, dependency := range task.DependsOn {
		switch taskStatuses[dependency] {
		case "ready", "completed":
		default:
			return false
		}
	}
	return true
}

func normalizeStatus(status string) string {
	return strings.TrimSpace(strings.ToLower(status))
}

func validStatus(status string) bool {
	switch normalizeStatus(status) {
	case "ready", "completed", "blocked", "blocked_external_write":
		return true
	default:
		return false
	}
}
