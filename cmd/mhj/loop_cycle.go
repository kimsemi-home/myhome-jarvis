package main

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/orchestrator"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func writeLoopWorkerCycle(ctx context.Context, root string) (scheduler.JobResult, error) {
	linearStatus := linear.CurrentStatus(root)
	linearSummary := linear.SummarizeStatus(linearStatus)
	linearNext := linear.NextIssue(ctx, root, http.DefaultClient)
	linearNextSummary := linear.SummarizeOperation(linearNext)
	securityStatus, err := security.StatusForRoot(root)
	if err != nil {
		return scheduler.JobResult{}, err
	}
	plannerStatus, err := planner.StatusForRoot(root)
	if err != nil {
		return scheduler.JobResult{}, err
	}
	if !linearNext.Synced {
		if err := linear.AppendOfflineEvent(root, "linear_next", linearNext.Message); err != nil {
			return scheduler.JobResult{}, err
		}
	}
	result := "scheduler heartbeat checkpoint recorded"
	if !securityStatus.OK {
		result = "scheduler heartbeat checkpoint recorded with public-safety findings"
	}
	path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
		Task:           "loop worker",
		LinearStatus:   linearSummary,
		LinearNext:     &linearNextSummary,
		PlannerStatus:  plannerStatus,
		SecurityStatus: securityStatus,
		Result:         result,
		Next:           "Continue local-first fixture and daemon surface expansion.",
	})
	if err != nil {
		return scheduler.JobResult{}, err
	}
	checkpointRef := path
	if rel, err := filepath.Rel(root, path); err == nil {
		checkpointRef = filepath.ToSlash(rel)
	}
	if _, err := appendControlPlaneManifest(root, "loop_worker_cycle", "loop_worker", checkpointRef); err != nil {
		return scheduler.JobResult{}, err
	}
	return scheduler.JobResult{Checkpoint: path}, nil
}
