package supervisor

import (
	"errors"
	"net/http"
	"os"
	"time"
)

func Status(root string, client *http.Client) DaemonStatus {
	status := missingStatus()
	state, err := ReadDaemonState(root)
	if errors.Is(err, os.ErrNotExist) {
		return status
	}
	if err != nil {
		status.Message = "daemon state is unreadable"
		return status
	}
	applyStateStatus(&status, state, client)
	return status
}

func missingStatus() DaemonStatus {
	return DaemonStatus{
		Name:      daemonStateName,
		StatePath: daemonStateRelativePath,
		Stale:     true,
		Message:   "no daemon state recorded",
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}
}

func applyStateStatus(status *DaemonStatus, state DaemonState, client *http.Client) {
	status.Recorded = true
	status.PID = state.PID
	status.Address = state.Address
	status.Version = state.Version
	status.StartedAt = state.StartedAt
	status.UpdatedAt = state.UpdatedAt
	status.ProcessRunning = processRunning(state.PID)
	status.ProbeURL = healthURL(state)
	status.ProbeOK, status.ProbeStatus = probeHealth(status.ProbeURL, client)
	status.Stale = !status.ProcessRunning || !status.ProbeOK
	status.Message = statusMessage(*status)
}

func statusMessage(status DaemonStatus) string {
	switch {
	case status.ProcessRunning && status.ProbeOK:
		return "daemon is reachable"
	case status.ProcessRunning:
		return "daemon process is recorded but health probe failed"
	default:
		return "daemon state is stale"
	}
}
