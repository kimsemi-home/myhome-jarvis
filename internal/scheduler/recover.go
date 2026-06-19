package scheduler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func Recover(root string, policy Policy) (State, error) {
	if err := policy.Validate(); err != nil {
		return State{}, err
	}
	data, err := os.ReadFile(statePath(root, policy.Name))
	if errors.Is(err, os.ErrNotExist) {
		return initialState(policy), nil
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

func initialState(policy Policy) State {
	now := time.Now().UTC()
	return State{
		Name:          policy.Name,
		StartedAt:     now,
		LastHeartbeat: now,
		NextRunAfter:  now,
	}
}
