package daemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntentAuditRecordsCommandErrors(t *testing.T) {
	server, err := New(DefaultConfig(t.TempDir(), "test"))
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/intent", bytes.NewBufferString(`{"command":"unknown","payload":{}}`))
	request.RemoteAddr = "127.0.0.1:1234"
	recorder := httptest.NewRecorder()

	server.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
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
		`"count": 1`,
		`"command": "unknown"`,
		`"success": false`,
		`"error_category": "unknown_command"`,
	} {
		if !bytes.Contains([]byte(statusBody), []byte(expected)) {
			t.Fatalf("expected %s in %s", expected, statusBody)
		}
	}
}
