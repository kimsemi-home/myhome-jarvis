package scheduler

import (
	"errors"
	"strings"
	"time"
)

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
