package supervisor

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestStatusProbesRecordedDaemonHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/health" {
			t.Fatalf("path = %q", request.URL.Path)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	host, port := probeServerHostPort(t, server.URL)
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

func probeServerHostPort(t *testing.T, url string) (string, int) {
	t.Helper()
	host, portText, err := net.SplitHostPort(strings.TrimPrefix(url, "http://"))
	if err != nil {
		t.Fatal(err)
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		t.Fatal(err)
	}
	return host, port
}
