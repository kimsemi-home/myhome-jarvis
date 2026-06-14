package supervisor

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const daemonStateName = "daemon"

type DaemonState struct {
	Name           string `json:"name"`
	PID            int    `json:"pid"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Address        string `json:"address"`
	Version        string `json:"version"`
	ExecuteEnabled bool   `json:"execute_enabled"`
	LANBindAllowed bool   `json:"lan_bind_allowed"`
	StartedAt      string `json:"started_at"`
	UpdatedAt      string `json:"updated_at"`
}

type DaemonStatus struct {
	Name           string `json:"name"`
	Recorded       bool   `json:"recorded"`
	StatePath      string `json:"state_path"`
	PID            int    `json:"pid,omitempty"`
	Address        string `json:"address,omitempty"`
	Version        string `json:"version,omitempty"`
	StartedAt      string `json:"started_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	ProcessRunning bool   `json:"process_running"`
	ProbeOK        bool   `json:"probe_ok"`
	ProbeStatus    int    `json:"probe_status,omitempty"`
	ProbeURL       string `json:"probe_url,omitempty"`
	Stale          bool   `json:"stale"`
	Message        string `json:"message"`
	CheckedAt      string `json:"checked_at"`
}

func NewDaemonState(root string, host string, port int, version string, executeEnabled bool, lanBindAllowed bool) (DaemonState, error) {
	if strings.TrimSpace(root) == "" {
		return DaemonState{}, errors.New("root is required")
	}
	if strings.TrimSpace(host) == "" {
		host = "127.0.0.1"
	}
	if port <= 0 || port > 65535 {
		return DaemonState{}, fmt.Errorf("invalid port %d", port)
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return DaemonState{
		Name:           daemonStateName,
		PID:            os.Getpid(),
		Host:           host,
		Port:           port,
		Address:        net.JoinHostPort(host, strconv.Itoa(port)),
		Version:        version,
		ExecuteEnabled: executeEnabled,
		LANBindAllowed: lanBindAllowed,
		StartedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func WriteDaemonState(root string, state DaemonState) (string, error) {
	if strings.TrimSpace(root) == "" {
		return "", errors.New("root is required")
	}
	if strings.TrimSpace(state.Name) == "" {
		state.Name = daemonStateName
	}
	if state.PID <= 0 {
		return "", errors.New("daemon pid is required")
	}
	if strings.TrimSpace(state.Host) == "" {
		state.Host = "127.0.0.1"
	}
	if state.Port <= 0 || state.Port > 65535 {
		return "", fmt.Errorf("invalid port %d", state.Port)
	}
	if strings.TrimSpace(state.Address) == "" {
		state.Address = net.JoinHostPort(state.Host, strconv.Itoa(state.Port))
	}
	now := time.Now().UTC().Format(time.RFC3339)
	if strings.TrimSpace(state.StartedAt) == "" {
		state.StartedAt = now
	}
	state.UpdatedAt = now

	path := statePath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return "", err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return "", err
	}
	return filepath.ToSlash(filepath.Join("data", "private", "supervisor", "daemon-state.json")), nil
}

func ReadDaemonState(root string) (DaemonState, error) {
	data, err := os.ReadFile(statePath(root))
	if err != nil {
		return DaemonState{}, err
	}
	var state DaemonState
	if err := json.Unmarshal(data, &state); err != nil {
		return DaemonState{}, fmt.Errorf("read daemon state: %w", err)
	}
	if state.Name == "" {
		state.Name = daemonStateName
	}
	return state, nil
}

func Status(root string, client *http.Client) DaemonStatus {
	checkedAt := time.Now().UTC().Format(time.RFC3339)
	status := DaemonStatus{
		Name:      daemonStateName,
		StatePath: filepath.ToSlash(filepath.Join("data", "private", "supervisor", "daemon-state.json")),
		Stale:     true,
		Message:   "no daemon state recorded",
		CheckedAt: checkedAt,
	}
	state, err := ReadDaemonState(root)
	if errors.Is(err, os.ErrNotExist) {
		return status
	}
	if err != nil {
		status.Message = "daemon state is unreadable"
		return status
	}

	status.Recorded = true
	status.PID = state.PID
	status.Address = state.Address
	status.Version = state.Version
	status.StartedAt = state.StartedAt
	status.UpdatedAt = state.UpdatedAt
	status.ProcessRunning = processRunning(state.PID)
	status.ProbeURL = healthURL(state)
	status.ProbeOK, status.ProbeStatus = probeHealth(status.ProbeURL, client)
	status.Stale = !status.ProcessRunning || !status.ProbeOK
	switch {
	case status.ProcessRunning && status.ProbeOK:
		status.Message = "daemon is reachable"
	case status.ProcessRunning:
		status.Message = "daemon process is recorded but health probe failed"
	default:
		status.Message = "daemon state is stale"
	}
	return status
}

func processRunning(pid int) bool {
	if pid <= 0 {
		return false
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return process.Signal(syscall.Signal(0)) == nil
}

func probeHealth(url string, client *http.Client) (bool, int) {
	if strings.TrimSpace(url) == "" {
		return false, 0
	}
	if client == nil {
		client = &http.Client{Timeout: 500 * time.Millisecond}
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, 0
	}
	response, err := client.Do(request)
	if err != nil {
		return false, 0
	}
	defer response.Body.Close()
	return response.StatusCode >= 200 && response.StatusCode < 300, response.StatusCode
}

func healthURL(state DaemonState) string {
	host := strings.TrimSpace(state.Host)
	switch host {
	case "", "0.0.0.0":
		host = "127.0.0.1"
	case "::":
		host = "::1"
	}
	if state.Port <= 0 {
		return ""
	}
	return "http://" + net.JoinHostPort(host, strconv.Itoa(state.Port)) + "/health"
}

func statePath(root string) string {
	return filepath.Join(root, "data", "private", "supervisor", "daemon-state.json")
}
