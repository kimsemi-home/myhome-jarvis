package supervisor

import (
	"encoding/json"
	"os"
	"path/filepath"
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
