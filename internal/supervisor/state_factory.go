package supervisor

import (
	"errors"
	"os"
	"strings"
)

func NewDaemonState(root string, host string, port int, version string, executeEnabled bool, lanBindAllowed bool) (DaemonState, error) {
	if strings.TrimSpace(root) == "" {
		return DaemonState{}, errors.New("root is required")
	}
	state := DaemonState{
		Name:           daemonStateName,
		PID:            os.Getpid(),
		Host:           host,
		Port:           port,
		Version:        version,
		ExecuteEnabled: executeEnabled,
		LANBindAllowed: lanBindAllowed,
	}
	return normalizeState(state)
}
