package planner

import (
	"errors"
	"fmt"
	"strings"
)

func validateTaskGraph(policy Policy) error {
	ids := map[string]bool{}
	hasExternalWriteBoundary := false
	gateTaskBlocked := false
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
			if id == strings.TrimSpace(policy.ExternalWriteGate.BoundaryTaskID) {
				gateTaskBlocked = true
			}
		}
		if !validStatus(task.Status) {
			return fmt.Errorf("planner task %q has invalid status %q", task.ID, task.Status)
		}
	}
	if !hasExternalWriteBoundary {
		return errors.New("planner task graph must include an external-write boundary")
	}
	if !gateTaskBlocked {
		return fmt.Errorf(
			"planner external-write gate task %q must be blocked_external_write",
			policy.ExternalWriteGate.BoundaryTaskID,
		)
	}
	return nil
}

func validateDependencies(tasks []Task) error {
	ids := map[string]bool{}
	for _, task := range tasks {
		ids[strings.TrimSpace(task.ID)] = true
	}
	for _, task := range tasks {
		for _, dependency := range task.DependsOn {
			if !ids[dependency] {
				return fmt.Errorf("planner task %q depends on unknown task %q", task.ID, dependency)
			}
		}
	}
	return nil
}
