package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Policy struct {
	Name              string
	Interval          time.Duration
	HeartbeatInterval time.Duration
	MinBackoff        time.Duration
	MaxBackoff        time.Duration
	CheckpointEvery   int
}

type State struct {
	Name                string    `json:"name"`
	StartedAt           time.Time `json:"started_at"`
	LastHeartbeat       time.Time `json:"last_heartbeat"`
	LastAttempt         time.Time `json:"last_attempt"`
	LastSuccess         time.Time `json:"last_success"`
	NextRunAfter        time.Time `json:"next_run_after"`
	LastCheckpoint      string    `json:"last_checkpoint"`
	LastError           string    `json:"last_error"`
	Cycles              int       `json:"cycles"`
	ConsecutiveFailures int       `json:"consecutive_failures"`
	Recovered           bool      `json:"recovered"`
}

type Snapshot struct {
	Name                     string `json:"name"`
	IntervalSeconds          int64  `json:"interval_seconds"`
	HeartbeatIntervalSeconds int64  `json:"heartbeat_interval_seconds"`
	MinBackoffSeconds        int64  `json:"min_backoff_seconds"`
	MaxBackoffSeconds        int64  `json:"max_backoff_seconds"`
	CheckpointEvery          int    `json:"checkpoint_every"`
	State                    State  `json:"state"`
}

type JobResult struct {
	Checkpoint string
}

type Job func(context.Context) (JobResult, error)

func ClosedLoopPolicy() Policy {
	return Policy{
		Name:              "closed_loop",
		Interval:          time.Minute,
		HeartbeatInterval: 15 * time.Second,
		MinBackoff:        5 * time.Second,
		MaxBackoff:        5 * time.Minute,
		CheckpointEvery:   1,
	}
}

func (policy Policy) Validate() error {
	if strings.TrimSpace(policy.Name) == "" {
		return errors.New("scheduler policy name is required")
	}
	if policy.Interval <= 0 {
		return errors.New("scheduler interval must be positive")
	}
	if policy.HeartbeatInterval <= 0 {
		return errors.New("scheduler heartbeat interval must be positive")
	}
	if policy.MinBackoff <= 0 || policy.MaxBackoff <= 0 {
		return errors.New("scheduler backoff values must be positive")
	}
	if policy.MaxBackoff < policy.MinBackoff {
		return errors.New("scheduler max backoff must be greater than or equal to min backoff")
	}
	if policy.CheckpointEvery <= 0 {
		return errors.New("scheduler checkpoint interval must be positive")
	}
	return nil
}

func Status(root string, policy Policy) (Snapshot, error) {
	state, err := Recover(root, policy)
	if err != nil {
		return Snapshot{}, err
	}
	return snapshot(policy, state), nil
}

func Recover(root string, policy Policy) (State, error) {
	if err := policy.Validate(); err != nil {
		return State{}, err
	}
	path := statePath(root, policy.Name)
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		now := time.Now().UTC()
		return State{
			Name:          policy.Name,
			StartedAt:     now,
			LastHeartbeat: now,
			NextRunAfter:  now,
		}, nil
	}
	if err != nil {
		return State{}, err
	}
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, fmt.Errorf("recover scheduler state: %w", err)
	}
	if state.Name == "" {
		state.Name = policy.Name
	}
	state.Recovered = true
	return state, nil
}

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
		if err := ctx.Err(); err != nil {
			return Snapshot{}, err
		}
		now := time.Now().UTC()
		state.LastHeartbeat = now
		state.LastAttempt = now
		state.Cycles++
		result, err := job(ctx)
		if err != nil {
			state.ConsecutiveFailures++
			state.LastError = err.Error()
			state.NextRunAfter = now.Add(backoff(policy, state.ConsecutiveFailures))
		} else {
			state.ConsecutiveFailures = 0
			state.LastError = ""
			state.LastSuccess = now
			state.NextRunAfter = now.Add(policy.Interval)
			if result.Checkpoint != "" {
				state.LastCheckpoint = checkpointPath(root, result.Checkpoint)
			}
		}
		if err := WriteState(root, state); err != nil {
			return Snapshot{}, err
		}
	}
	return snapshot(policy, state), nil
}

func WriteState(root string, state State) error {
	if strings.TrimSpace(state.Name) == "" {
		return errors.New("scheduler state name is required")
	}
	dir := filepath.Join(root, "data", "private", "scheduler")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(statePath(root, state.Name), data, 0o600)
}

func statePath(root string, name string) string {
	return filepath.Join(root, "data", "private", "scheduler", name+"-state.json")
}

func checkpointPath(root string, path string) string {
	cleaned := filepath.Clean(path)
	if root != "" {
		relative, err := filepath.Rel(root, cleaned)
		if err == nil && relative != "." && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			return filepath.ToSlash(relative)
		}
	}
	return filepath.ToSlash(cleaned)
}

func snapshot(policy Policy, state State) Snapshot {
	return Snapshot{
		Name:                     policy.Name,
		IntervalSeconds:          int64(policy.Interval.Seconds()),
		HeartbeatIntervalSeconds: int64(policy.HeartbeatInterval.Seconds()),
		MinBackoffSeconds:        int64(policy.MinBackoff.Seconds()),
		MaxBackoffSeconds:        int64(policy.MaxBackoff.Seconds()),
		CheckpointEvery:          policy.CheckpointEvery,
		State:                    state,
	}
}

func backoff(policy Policy, failures int) time.Duration {
	if failures <= 0 {
		return policy.Interval
	}
	delay := policy.MinBackoff
	for index := 1; index < failures; index++ {
		delay *= 2
		if delay >= policy.MaxBackoff {
			return policy.MaxBackoff
		}
	}
	return delay
}
