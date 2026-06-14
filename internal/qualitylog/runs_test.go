package qualitylog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestStatusForRootMissingJournal(t *testing.T) {
	status, err := StatusForRoot(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists {
		t.Fatalf("expected missing journal, got %#v", status)
	}
	if status.Path != RelativePath {
		t.Fatalf("path = %q", status.Path)
	}
}

func TestAppendRunWritesRedactedPrivateJournal(t *testing.T) {
	root := t.TempDir()
	run := NewRun(time.Now().Add(-time.Second), true, []Step{
		{Name: "go test", Status: "pass"},
		{Name: "flutter analyze", Status: "skip"},
	})

	if err := AppendRun(root, run); err != nil {
		t.Fatal(err)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 {
		t.Fatalf("status = %#v", status)
	}
	if status.Last == nil || !status.Last.OK || status.Last.PassCount != 1 || status.Last.SkipCount != 1 {
		t.Fatalf("last = %#v", status.Last)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(RelativePath)))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"output", "command", root} {
		if strings.Contains(string(data), forbidden) {
			t.Fatalf("quality journal leaked %q in %s", forbidden, data)
		}
	}
}

func TestNewRunCountsStepStatuses(t *testing.T) {
	run := NewRun(time.Now(), false, []Step{
		{Name: "a", Status: "pass"},
		{Name: "b", Status: "fail"},
		{Name: "c", Status: "skip"},
		{Name: "d", Status: "weird"},
	})

	if run.StepCount != 4 || run.PassCount != 1 || run.FailCount != 1 || run.SkipCount != 1 {
		t.Fatalf("run = %#v", run)
	}
	if run.Steps[3].Status != "unknown" {
		t.Fatalf("unknown status = %q", run.Steps[3].Status)
	}
}
