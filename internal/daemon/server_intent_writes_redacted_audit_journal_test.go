package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestIntentWritesRedactedAuditJournal(t *testing.T) {
	root := t.TempDir()
	server, err := New(DefaultConfig(root, "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"open-url","payload":{"url":"https://example.invalid/private"},"execute":false}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", recorder.Code, recorder.Body.String())
	}
	statusRequest := httptest.NewRequest(http.MethodGet, "/audit/status", nil)
	statusRequest.RemoteAddr = "127.0.0.1:1234"
	statusRecorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(statusRecorder, statusRequest)

	if statusRecorder.Code != http.StatusOK {
		t.Fatalf("status = %d body = %s", statusRecorder.Code, statusRecorder.Body.String())
	}
	statusBody := statusRecorder.Body.String()
	for _, expected := range []string{
		`"path": "data/private/audit/command-intents.jsonl"`,
		`"count": 1`,
		`"command": "open_url"`,
		`"source": "daemon"`,
		`"success": true`,
	} {
		if !bytes.Contains([]byte(statusBody), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, statusBody)
		}
	}
	data, err := os.ReadFile(filepath.Join(root, "data", "private", "audit", "command-intents.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"payload", "argv", "https://example.invalid/private"} {
		if bytes.Contains(data, []byte(forbidden)) {
			t.Fatalf("audit journal leaked %q in %s", forbidden, data)
		}
	}
}
