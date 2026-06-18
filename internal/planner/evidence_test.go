package planner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

func TestStatusForRootSeparatesExternalWriteGateFromSyncedEvidence(t *testing.T) {
	root := t.TempDir()
	generatedDir := filepath.Join(root, "generated")
	if err := os.MkdirAll(generatedDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writePlannerPolicyFixture(
		t,
		filepath.Join(generatedDir, "planner.generated.json"),
		"data/private/checkpoints",
		"completed",
	)
	if err := linear.AppendWriteEvidence(root, "linear_transition", "KIM-12"); err != nil {
		t.Fatal(err)
	}

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}

	if !status.ExternalWriteGate.BoundaryTaskBlocked || status.BlockedExternalWriteCount != 1 {
		t.Fatalf("external write gate not preserved: %#v", status)
	}
	if status.LinearWriteEvidence.SyncedMutationCount != 1 || !status.LinearWriteEvidence.HasSyncedMutation {
		t.Fatalf("linear write evidence status = %#v", status.LinearWriteEvidence)
	}
	if status.LinearWriteEvidence.LatestSyncedMutation == nil {
		t.Fatalf("latest write evidence missing: %#v", status.LinearWriteEvidence)
	}
	if status.LinearWriteEvidence.LatestSyncedMutation.Action != "linear_transition" || status.LinearWriteEvidence.LatestSyncedMutation.IssueKey != "KIM-12" {
		t.Fatalf("latest write evidence = %#v", status.LinearWriteEvidence.LatestSyncedMutation)
	}
}
