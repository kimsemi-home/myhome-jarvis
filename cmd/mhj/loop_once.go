package main

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/orchestrator"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func loopOnce(root string) error {
	linearStatus := linear.CurrentStatus(root)
	linearSummary := linear.SummarizeStatus(linearStatus)
	linearNext := linear.NextIssue(context.Background(), root, http.DefaultClient)
	linearNextSummary := linear.SummarizeOperation(linearNext)
	securityStatus, err := security.StatusForRoot(root)
	if err != nil {
		return err
	}
	plannerStatus, err := planner.StatusForRoot(root)
	if err != nil {
		return err
	}
	if linearStatus.Mode == "offline" {
		if err := linear.AppendOfflineEvent(root, "loop_once", "Local loop ran without Linear sync; synced=false."); err != nil {
			return err
		}
	}
	if !linearNext.Synced {
		if err := linear.AppendOfflineEvent(root, "linear_next", linearNext.Message); err != nil {
			return err
		}
	}
	result := "checkpoint recorded"
	if !securityStatus.OK {
		result = "checkpoint recorded with public-safety findings"
	}
	path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
		Task:           "loop once",
		LinearStatus:   linearSummary,
		LinearNext:     &linearNextSummary,
		PlannerStatus:  plannerStatus,
		SecurityStatus: securityStatus,
		Result:         result,
		Next:           "Continue local-first closed-loop hardening.",
	})
	if err != nil {
		return err
	}
	checkpointPath, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	if _, err := appendControlPlaneManifest(root, "loop_once", "loop_once", filepath.ToSlash(checkpointPath)); err != nil {
		return err
	}
	return writeJSON(map[string]any{
		"ok":              securityStatus.OK,
		"checkpoint":      filepath.ToSlash(checkpointPath),
		"linear":          linearSummary,
		"linear_next":     linearNextSummary,
		"planner_status":  plannerStatus,
		"security_status": securityStatus,
	})
}
