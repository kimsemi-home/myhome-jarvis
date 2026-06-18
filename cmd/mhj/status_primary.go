package main

import (
	"errors"

	"github.com/kimsemi-home/myhome-jarvis/internal/agentcluster"
	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/codeshape"
	"github.com/kimsemi-home/myhome-jarvis/internal/connectors"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/repo"
)

func repoStatus(root string) error {
	status, err := repo.Inspect(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func plannerStatus(root string) error {
	status, err := planner.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func connectorsStatus(root string) error {
	status, err := connectors.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func agentClusterStatus(root string) error {
	status, err := agentcluster.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func auditStatus(root string) error {
	status, err := audit.CommandIntentStatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func qualityStatus(root string) error {
	status, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func codeShapeStatus(root string) error {
	status, err := codeshape.StatusForRoot(root)
	if err != nil {
		return err
	}
	if err := writeJSON(status); err != nil {
		return err
	}
	if !status.OK {
		return errors.New("code shape budget regression")
	}
	return nil
}
