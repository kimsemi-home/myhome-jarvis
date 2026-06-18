package daemon

import (
	"bytes"
	"context"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntentExecutesWithExplicitDaemonGate(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.Execute = true
	config.CommandPlatform = "darwin"
	calls := 0
	config.CommandRunner = func(runnerContext context.Context, invocation commands.Invocation) commands.Execution {
		calls++
		if invocation.Argv[0] != "osascript" {
			t.Fatalf("argv = %#v", invocation.Argv)
		}
		return commands.Execution{Executed: true, ExitCode: 0}
	}
	server, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"volume-set","payload":{"level":30},"execute":true}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	if calls != 1 {
		t.Fatalf("runner calls = %d", calls)
	}
	body := recorder.Body.String()
	for _, expected := range []string{
		`"dry_run": false`,
		`"execute_allowed": true`,
		`"executed": true`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}
