package scheduler

import (
	"context"
	"errors"
	"time"
)

func RunCycles(ctx context.Context, root string, policy Policy, cycles int, job Job) (Snapshot, error) {
	if cycles <= 0 {
		return Snapshot{}, errors.New("cycles must be positive")
	}
	if job == nil {
		return Snapshot{}, errors.New("scheduler job is required")
	}
	state, err := Recover(root, policy)
	if err != nil {
		return Snapshot{}, err
	}
	for index := 0; index < cycles; index++ {
		state, err = runCycle(ctx, root, policy, state, job)
		if err != nil {
			return Snapshot{}, err
		}
	}
	return snapshot(policy, state), nil
}

func runCycle(ctx context.Context, root string, policy Policy, state State, job Job) (State, error) {
	if err := ctx.Err(); err != nil {
		return State{}, err
	}
	now := time.Now().UTC()
	state.LastHeartbeat = now
	state.LastAttempt = now
	state.Cycles++
	result, err := job(ctx)
	applyResult(root, policy, &state, now, result, err)
	return state, WriteState(root, state)
}

func applyResult(root string, policy Policy, state *State, now time.Time, result JobResult, err error) {
	if err != nil {
		state.ConsecutiveFailures++
		state.LastError = err.Error()
		state.NextRunAfter = now.Add(backoff(policy, state.ConsecutiveFailures))
		return
	}
	state.ConsecutiveFailures = 0
	state.LastError = ""
	state.LastSuccess = now
	state.NextRunAfter = now.Add(policy.Interval)
	if result.Checkpoint != "" {
		state.LastCheckpoint = checkpointPath(root, result.Checkpoint)
	}
}
