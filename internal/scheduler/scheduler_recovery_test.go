package scheduler

import (
	"context"
	"errors"
	"testing"
)

func TestFailureBackoffAndRecovery(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	snapshot, err := RunCycles(context.Background(), root, policy, 1, func(context.Context) (JobResult, error) {
		return JobResult{}, errors.New("planned failure")
	})
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.State.ConsecutiveFailures != 1 {
		t.Fatalf("failures = %d", snapshot.State.ConsecutiveFailures)
	}
	if snapshot.State.LastError != "planned failure" {
		t.Fatalf("last error = %q", snapshot.State.LastError)
	}
	if snapshot.State.NextRunAfter.Sub(snapshot.State.LastAttempt) != policy.MinBackoff {
		t.Fatalf("next run/backoff mismatch: %#v", snapshot.State)
	}

	recovered, err := Recover(root, policy)
	if err != nil {
		t.Fatal(err)
	}
	if !recovered.Recovered || recovered.ConsecutiveFailures != 1 {
		t.Fatalf("unexpected recovered state: %#v", recovered)
	}
}
