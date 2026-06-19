package supervisor

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func normalizeState(state DaemonState) (DaemonState, error) {
	if strings.TrimSpace(state.Name) == "" {
		state.Name = daemonStateName
	}
	if state.PID <= 0 {
		return DaemonState{}, errors.New("daemon pid is required")
	}
	if strings.TrimSpace(state.Host) == "" {
		state.Host = "127.0.0.1"
	}
	if state.Port <= 0 || state.Port > 65535 {
		return DaemonState{}, fmt.Errorf("invalid port %d", state.Port)
	}
	if strings.TrimSpace(state.Address) == "" {
		state.Address = net.JoinHostPort(state.Host, strconv.Itoa(state.Port))
	}
	return stampState(state), nil
}

func stampState(state DaemonState) DaemonState {
	now := time.Now().UTC().Format(time.RFC3339)
	if strings.TrimSpace(state.StartedAt) == "" {
		state.StartedAt = now
	}
	state.UpdatedAt = now
	return state
}
