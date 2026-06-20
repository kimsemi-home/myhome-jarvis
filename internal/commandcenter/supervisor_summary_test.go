package commandcenter

import "testing"

func TestStatusIncludesRedactedSupervisorSummary(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if status.Supervisor.StatePath != "data/private/supervisor/daemon-state.json" {
		t.Fatalf("supervisor path = %q", status.Supervisor.StatePath)
	}
	if status.Supervisor.Message == "" {
		t.Fatalf("supervisor summary = %#v", status.Supervisor)
	}
	if status.LocalRuntime.Stale != status.Supervisor.Stale ||
		status.LocalRuntime.ProcessRunning != status.Supervisor.ProcessRunning ||
		status.LocalRuntime.ProbeOK != status.Supervisor.ProbeOK {
		t.Fatalf("runtime/supervisor mismatch = %#v %#v",
			status.LocalRuntime, status.Supervisor)
	}
}
