package scheduler

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestClosedLoopPolicyHasSafetyBounds(t *testing.T) {
	policy := ClosedLoopPolicy()
	if err := policy.Validate(); err != nil {
		t.Fatal(err)
	}
	if policy.Interval <= policy.HeartbeatInterval {
		t.Fatalf("interval should leave room for heartbeats: %#v", policy)
	}
	if policy.MinBackoff <= 0 || policy.MaxBackoff < policy.MinBackoff {
		t.Fatalf("invalid backoff bounds: %#v", policy)
	}
	if policy.CheckpointEvery != 1 {
		t.Fatalf("checkpoint every = %d", policy.CheckpointEvery)
	}
}

func TestRunCyclesWritesHeartbeatAndPrivateState(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	checkpoint := filepath.Join(root, "data", "private", "checkpoints", "cycle.json")
	snapshot, err := RunCycles(context.Background(), root, policy, 2, func(context.Context) (JobResult, error) {
		return JobResult{Checkpoint: checkpoint}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.State.Cycles != 2 {
		t.Fatalf("cycles = %d", snapshot.State.Cycles)
	}
	if snapshot.State.LastHeartbeat.IsZero() || snapshot.State.LastSuccess.IsZero() {
		t.Fatalf("missing heartbeat or success: %#v", snapshot.State)
	}
	if snapshot.State.LastCheckpoint != "data/private/checkpoints/cycle.json" {
		t.Fatalf("checkpoint = %q", snapshot.State.LastCheckpoint)
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "private", "scheduler", "test_loop-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("scheduler state file is empty")
	}
}

func testPolicy() Policy {
	return Policy{
		Name:              "test_loop",
		Interval:          time.Minute,
		HeartbeatInterval: 10 * time.Second,
		MinBackoff:        5 * time.Second,
		MaxBackoff:        time.Minute,
		CheckpointEvery:   1,
	}
}
