package supervisor

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestStatusReportsMissingStateWithRelativePath(t *testing.T) {
	status := Status(t.TempDir(), nil)

	if status.Recorded {
		t.Fatalf("expected missing state, got %#v", status)
	}
	if status.StatePath != "data/private/supervisor/daemon-state.json" {
		t.Fatalf("state path = %q", status.StatePath)
	}
	if !status.Stale {
		t.Fatalf("expected missing state to be stale")
	}
}

func TestWriteAndReadDaemonStateUsesPrivateFile(t *testing.T) {
	root := t.TempDir()
	state, err := NewDaemonState(root, "127.0.0.1", 3888, "test", false, false)
	if err != nil {
		t.Fatal(err)
	}

	path, err := WriteDaemonState(root, state)
	if err != nil {
		t.Fatal(err)
	}

	if path != "data/private/supervisor/daemon-state.json" {
		t.Fatalf("path = %q", path)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), root) {
		t.Fatalf("state leaked absolute root in %s", data)
	}
	var decoded DaemonState
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.PID != os.Getpid() {
		t.Fatalf("pid = %d", decoded.PID)
	}
	if decoded.Address != "127.0.0.1:3888" {
		t.Fatalf("address = %q", decoded.Address)
	}
}

func TestStatusProbesRecordedDaemonHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/health" {
			t.Fatalf("path = %q", request.URL.Path)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	host, portText, err := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
	if err != nil {
		t.Fatal(err)
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	state, err := NewDaemonState(root, host, port, "test", false, false)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := WriteDaemonState(root, state); err != nil {
		t.Fatal(err)
	}

	status := Status(root, server.Client())

	if !status.Recorded {
		t.Fatalf("expected recorded status")
	}
	if !status.ProcessRunning {
		t.Fatalf("expected current process to be running")
	}
	if !status.ProbeOK || status.ProbeStatus != http.StatusOK {
		t.Fatalf("probe = %v status = %d", status.ProbeOK, status.ProbeStatus)
	}
	if status.Stale {
		t.Fatalf("expected reachable daemon state, got %#v", status)
	}
}
