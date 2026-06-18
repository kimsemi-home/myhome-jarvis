package daemon

import (
	"bytes"
	"context"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntentExecuteRequiresDaemonExecuteMode(t *testing.T) {
	config := DefaultConfig(t.TempDir(), "test")
	config.CommandPlatform = "darwin"
	config.CommandRunner = func(runnerContext context.Context, invocation commands.Invocation) commands.Execution {
		t.Fatal("runner must not be called when daemon execute mode is disabled")
		return commands.Execution{}
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
	body := recorder.Body.String()
	for _, expected := range []string{
		`"dry_run": true`,
		`"execute_allowed": false`,
		`"execute was requested but daemon execute mode is disabled"`,
	} {
		if !bytes.Contains([]byte(body), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
}
