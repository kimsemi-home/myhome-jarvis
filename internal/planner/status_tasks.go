package planner

import "strings"

func summarizeTasks(tasks []Task, status *Status) {
	taskStatuses := make(map[string]string, len(tasks))
	for _, task := range tasks {
		normalized := normalizeStatus(task.Status)
		taskStatuses[task.ID] = normalized
		addTaskStatus(task, normalized, status)
	}
	for index := range tasks {
		task := tasks[index]
		if normalizeStatus(task.Status) == "ready" && dependenciesSatisfied(task, taskStatuses) {
			status.NextTask = &task
			return
		}
	}
}

func addTaskStatus(task Task, normalized string, status *Status) {
	switch normalized {
	case "ready":
		status.ReadyCount++
	case "completed":
		status.CompletedCount++
	case "blocked_external_write":
		status.BlockedExternalWriteCount++
		status.BlockedExternalWriteTasks = append(status.BlockedExternalWriteTasks, task)
		if task.ID == status.ExternalWriteGate.BoundaryTaskID {
			status.ExternalWriteGate.BoundaryTaskBlocked = true
		}
	}
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
