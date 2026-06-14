package audit

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func TestCommandIntentStatusMissingJournal(t *testing.T) {
	status, err := CommandIntentStatusForRoot(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists {
		t.Fatalf("expected missing journal, got %#v", status)
	}
	if status.Path != commandIntentRelativePath {
		t.Fatalf("path = %q", status.Path)
	}
}

func TestAppendCommandIntentWritesPrivateRedactedJSONL(t *testing.T) {
	root := t.TempDir()
	event := CommandIntentFromPlan("daemon", "open-url", false, commands.Plan{
		Name:   "open_url",
		DryRun: true,
		Invocations: []commands.Invocation{
			{Label: "open_url", Argv: []string{"open", "https://example.invalid/private"}, URL: "https://example.invalid/private"},
		},
	}, nil)

	if err := AppendCommandIntent(root, event); err != nil {
		t.Fatal(err)
	}
	status, err := CommandIntentStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 {
		t.Fatalf("status = %#v", status)
	}
	if status.Last == nil || status.Last.Command != "open_url" {
		t.Fatalf("last = %#v", status.Last)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(commandIntentRelativePath)))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"payload", "argv", "https://example.invalid/private", root} {
		if strings.Contains(string(data), forbidden) {
			t.Fatalf("audit journal leaked %q in %s", forbidden, data)
		}
	}
}

func TestCommandIntentFromPlanClassifiesFailures(t *testing.T) {
	event := CommandIntentFromPlan("cli", "missing", true, commands.Plan{}, errors.New(`unknown command "missing"`))

	if event.Success {
		t.Fatalf("expected failure event")
	}
	if event.Command != "missing" {
		t.Fatalf("command = %q", event.Command)
	}
	if event.ErrorCategory != "unknown_command" {
		t.Fatalf("error category = %q", event.ErrorCategory)
	}
}
