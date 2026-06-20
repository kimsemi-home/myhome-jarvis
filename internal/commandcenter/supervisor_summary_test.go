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
	if !status.Supervisor.Stale || status.Supervisor.Message == "" {
		t.Fatalf("supervisor summary = %#v", status.Supervisor)
	}
}
