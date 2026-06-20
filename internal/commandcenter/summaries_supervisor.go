package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/supervisor"

func summarizeSupervisor(status supervisor.DaemonStatus) SupervisorSummary {
	return SupervisorSummary{
		Recorded:       status.Recorded,
		StatePath:      status.StatePath,
		ProcessRunning: status.ProcessRunning,
		ProbeOK:        status.ProbeOK,
		Stale:          status.Stale,
		Message:        status.Message,
	}
}
