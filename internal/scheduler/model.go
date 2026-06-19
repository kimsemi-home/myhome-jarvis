package scheduler

import (
	"context"
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
